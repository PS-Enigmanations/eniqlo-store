package repository

import (
	"context"
	"encoding/json"
	"enigmanations/eniqlo-store/internal/transaction"
	"enigmanations/eniqlo-store/internal/transaction/request"
	"enigmanations/eniqlo-store/util"
	"enigmanations/eniqlo-store/pkg/uuid"
	"enigmanations/eniqlo-store/pkg/validate"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionRepository interface {
	Save(ctx context.Context, model transaction.Transaction, total float64) (*transaction.Transaction, error)
	SaveDetails(ctx context.Context, models []transaction.ProductDetail, transactionId string) error
	GetAllByParams(ctx context.Context, params *request.TransactionGetAllQueryParams) ([]*transaction.Transaction, error)
}

type Database struct {
	pool *pgxpool.Pool
}

func NewTransactionRepository(pool *pgxpool.Pool) TransactionRepository {
	return &Database{
		pool: pool,
	}
}

func (db *Database) Save(ctx context.Context, model transaction.Transaction, total float64) (*transaction.Transaction, error) {
	const sql = `INSERT into transactions
		("id", "customer_id", "total", "paid", "change")
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;`

	row := db.pool.QueryRow(
		ctx,
		sql,
		model.TransactionId,
		model.CustomerId,
		total,
		model.Paid,
		model.Change,
	)

	c := new(transaction.Transaction)

	err := row.Scan(
		&c.TransactionId,
	)

	if err != nil {
		return nil, fmt.Errorf("Save %w", err)
	}

	return c, nil
}

func (db *Database) SaveDetails(ctx context.Context, models []transaction.ProductDetail, transactionId string) error {
	const sql = `INSERT into transaction_details
		("id", "transaction_id", "product_id", "quantity")
		VALUES ($1, $2, $3, $4);`

	for _, model := range models {
		id := uuid.New()
		_, err := db.pool.Exec(ctx, sql, id, transactionId, model.ProductId, model.Quantity)
		if err != nil {
			return fmt.Errorf("Save Transaction Detail %w", err)
		}
	}

	return nil
}

func (db *Database) GetAllByParams(ctx context.Context, params *request.TransactionGetAllQueryParams) ([]*transaction.Transaction, error) {
	var (
		args  []any
		where []string
		order []string
	)

	sql := fmt.Sprintf(`
		SELECT
			t.id AS transactionId,
			t.customer_id AS customerId,
			json_agg(json_build_object(
				'productId', td.product_id,
				'quantity', td.quantity
			)) AS productDetails,
			t.paid,
			t.change,
			t.created_at
		FROM transactions t JOIN
			transaction_details td ON t.id = td.transaction_id
		`)

	// Filter customer id
	if params.CustomerId != "" {
		args = append(args, params.CustomerId)
		where = append(where, fmt.Sprintf(`"customer_id" = $%d`, len(args)))
	}

	// Merge where clauses
	if len(where) > 0 {
		w := " WHERE " + strings.Join(where, " AND ") + " GROUP BY t.id " // #nosec G202
		sql += w
	} else {
		w := " GROUP BY t.id "
		sql += w
	}

	// Order by created at
	if params.CreatedAt != "" && validate.IsStrSortType(params.CreatedAt) {
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

	var transactions []*transaction.Transaction
	if rows != nil {
		for rows.Next() {
			var productDetailsJSON []byte
			c := new(transaction.Transaction)

			err := rows.Scan(
				&c.TransactionId,
				&c.CustomerId,
				&productDetailsJSON, // Scan into a byte slice
				&c.Paid,
				&c.Change,
				&c.CreatedAt,
			)

			if err != nil {
				fmt.Println(err)
				return nil, err
			}

			// Unmarshal the JSON data into the slice of ProductDetail pointers
			var productDetails []*transaction.ProductDetail
			err = json.Unmarshal(productDetailsJSON, &productDetails)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			c.ProductDetails = productDetails

			transactions = append(transactions, c)
		}
	}
	return transactions, nil
}
