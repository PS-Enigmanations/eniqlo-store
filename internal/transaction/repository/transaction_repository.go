package repository

import (
	"enigmanations/eniqlo-store/internal/transaction"
	"enigmanations/eniqlo-store/internal/transaction/request"
	"enigmanations/eniqlo-store/util"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"fmt"
	"encoding/json"
	"strings"
)

type TransactionRepository interface {
	Save(ctx context.Context, model transaction.Transaction) (*transaction.Transaction, error)
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

func (db *Database) Save(ctx context.Context, model transaction.Transaction) (*transaction.Transaction, error) {
	return nil, nil
}

func (db *Database) GetAllByParams(ctx context.Context, params *request.TransactionGetAllQueryParams) ([]*transaction.Transaction, error) {
	var (
		args  []any
		where []string
		order []string
	)

	fmt.Println(params.CustomerId)
	
	 
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

	fmt.Println(sql)
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