package repository

import (
	"context"
	"enigmanations/eniqlo-store/internal/product"
	"enigmanations/eniqlo-store/internal/product/request"
	"enigmanations/eniqlo-store/util"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository interface {
	SearchProducts(ctx context.Context, params *request.SearchProductQueryParams, alwaysAvailable bool) ([]*product.Product, error)
}

type database struct {
	pool *pgxpool.Pool
}

func NewProductRepository(pool *pgxpool.Pool) ProductRepository {
	return &database{
		pool: pool,
	}
}

// Always available on search sku
func (db *database) SearchProducts(ctx context.Context, params *request.SearchProductQueryParams, alwaysAvailable bool) ([]*product.Product, error) {
	var (
		args  []any
		where []string
		order []string
	)

	sql := `
		SELECT
			p.id,
			p.name,
			p.sku,
			p.category,
			p.image_url,
			p.notes,
			p.price,
			p.stock,
			p."location",
			p.is_available,
			p.created_at
		FROM products p
	`

	// Request should only show product that have isAvailable == true
	if alwaysAvailable {
		args = append(args, true)
		where = append(where, fmt.Sprintf(`"is_available" = $%d`, len(args)))
	}

	// ID
	if params.Id != "" {
		args = append(args, params.Id)
		where = append(where, fmt.Sprintf(`p.id = $%d`, len(args)))
	}

	// Name
	if params.Name != "" {
		args = append(args, params.Name)
		where = append(where, fmt.Sprintf(`
			(
				p."_search" @@ plainto_tsquery('english', $%d) OR
				p."name" ilike '$%s'
			)
		`, len(args), "%"+fmt.Sprintf("$%d", len(args))+"%"))
	}
	// Category
	if params.Category != "" {
		if product.IsHasCategory(params.Category) {
			args = append(args, params.Category)
			where = append(where, fmt.Sprintf(`"category" ilike $%d`, len(args)))
		}
	}
	// Sku
	if params.Sku != "" {
		args = append(args, params.Sku)
		where = append(where, fmt.Sprintf(`"sku" = $%d`, len(args)))
	}
	// In Stock
	if params.InStock != "" {
		if params.InStock == "true" {
			args = append(args, "0")
			where = append(where, fmt.Sprintf(`"stock" > $%d`, len(args)))
		} else if params.InStock == "false" {
			args = append(args, "0")
			where = append(where, fmt.Sprintf(`"stock" = $%d`, len(args)))
		}
	}
	if params.IsAvailable != "" && util.IsBoolFromStr(params.IsAvailable) && !alwaysAvailable {
		isAvailable, err := strconv.ParseBool(params.IsAvailable)
		if nil != err {
			return nil, err
		}

		args = append(args, isAvailable)
		where = append(where, fmt.Sprintf(`"is_available" = $%d`, len(args)))
	}

	// Merge where clauses
	if len(where) > 0 {
		w := " WHERE " + strings.Join(where, " AND ") + " AND deleted_at IS NULL" // #nosec G202
		sql += w
	} else {
		w := " WHERE deleted_at IS NULL"
		sql += w
	}

	// Order by will only execute first operation at a time,
	// Apply based on latest order by
	//
	// Order by price
	if params.Price != "" && util.IsSortType(params.Price) {
		value := fmt.Sprintf("price %s", params.Price)

		order = []string{}
		order = append(order, value)
	}
	// Order by created at
	if params.CreatedAt != "" && util.IsSortType(params.CreatedAt) {
		value := fmt.Sprintf("created_at %s", params.CreatedAt)
		order = []string{}
		order = append(order, value)
	}

	// Merge order clauses
	if len(order) > 0 {
		o := " ORDER BY " + strings.Join(order, ", ")
		sql += o
	}

	// Limit (default: 5)
	if params.Limit != "" {
		sql += fmt.Sprintf(` LIMIT %s`, params.Limit)
	} else {
		sql += fmt.Sprintf(` LIMIT %d`, 5)
	}
	// Offset (default: 0)
	if params.Offset != "" {
		sql += fmt.Sprintf(` OFFSET %s`, params.Offset)
	} else {
		sql += fmt.Sprintf(` OFFSET %d`, 0)
	}

	rows, err := db.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	// close rows if error ocur
	defer rows.Close()

	// Iterate rows
	var products []*product.Product
	if rows != nil {
		for rows.Next() {
			// create 'c' for struct 'Cat'
			p := new(product.Product)

			// scan rows and place it in 'p' (product) container
			err := rows.Scan(
				&p.Id,
				&p.Name,
				&p.Sku,
				&p.Category,
				&p.ImageUrl,
				&p.Notes,
				&p.Price,
				&p.Stock,
				&p.Location,
				&p.IsAvailable,
				&p.CreatedAt,
			)

			// return nil and error if scan operation fail
			if err != nil {
				return nil, err
			}

			// add p to products slice
			products = append(products, p)
		}
	}

	return products, nil
}
