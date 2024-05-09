package repository

import (
	"context"
	"enigmanations/eniqlo-store/internal/customer"
	"enigmanations/eniqlo-store/internal/customer/request"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"strings"
)

type CustomerRepository interface {
	Save(ctx context.Context, model customer.Customer) (*customer.Customer, error)
	FindByPhoneNumber(ctx context.Context, phoneNumber string) (*customer.Customer, error)
	GetAllByParams(ctx context.Context, params *request.CustomerGetAllQueryParams) ([]*customer.Customer, error)
}

type Database struct {
	pool *pgxpool.Pool
}

func NewCustomerRepository(pool *pgxpool.Pool) CustomerRepository {
	return &Database{
		pool: pool,
	}
}

func (db *Database) FindByPhoneNumber(ctx context.Context, phoneNumber string) (*customer.Customer, error) {
	const sql = `
		SELECT id, name, phone_number FROM customers WHERE phone_number = $1 AND deleted_at IS NULL LIMIT 1;
	`
	row := db.pool.QueryRow(ctx, sql, phoneNumber)
	c := new(customer.Customer)
	err := row.Scan(
		&c.Id,
		&c.Name,
		&c.PhoneNumber,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return c, nil
}

func (db *Database) Save(ctx context.Context, model customer.Customer) (*customer.Customer, error) {
	const sql = `INSERT into customers
		("id", "name", "phone_number")
		VALUES ($1, $2, $3)
		RETURNING id, name, phone_number;`

	row := db.pool.QueryRow(
		ctx,
		sql,
		model.Id,
		model.Name,
		model.PhoneNumber,
	)

	c := new(customer.Customer)

	err := row.Scan(
		&c.Id,
		&c.Name,
		&c.PhoneNumber,
	)

	if err != nil {
		return nil, fmt.Errorf("Save %w", err)
	}

	return c, nil
}

func (db *Database) GetAllByParams(ctx context.Context, params *request.CustomerGetAllQueryParams) ([]*customer.Customer, error) {
	var (
		args  []any
		where []string
	)

	sql := fmt.Sprintf(`
		SELECT
			id,
			name,
			phone_number
		FROM customers
		`)

	// Search
	if params.Name != "" {
		args = append(args, "%"+params.Name+"%")
		where = append(where, fmt.Sprintf(`"name" ilike $%d`, len(args)))
	}


	// Phone Number
	if params.PhoneNumber != "" {
		// Ensure the phone number has a leading '%' to match any prefix
		phoneNumberWildcard := "%+" + params.PhoneNumber +"%"
	
		args = append(args, phoneNumberWildcard)
		where = append(where, fmt.Sprintf(`phone_number ilike $%d`, len(args)))
	}

	// Merge where clauses
	if len(where) > 0 {
		w := " WHERE " + strings.Join(where, " AND ") + " AND deleted_at IS NULL" // #nosec G202
		sql += w
	} else {
		w := " WHERE deleted_at IS NULL"
		sql += w
	}

	rows, err := db.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	// close rows if error ocur
	defer rows.Close()

	// iterate Rows
	var customers []*customer.Customer
	if rows != nil {
		for rows.Next() {
			// create 'c' for struct 'Cat'
			c := new(customer.Customer)

			// scan rows and place it in 'c' (cat) container
			err := rows.Scan(
				&c.Id,
				&c.Name,
				&c.PhoneNumber,
			)

			// return nil and error if scan operation fail
			if err != nil {
				fmt.Println(err)
				return nil, err
			}

			// add c to cats slice
			customers = append(customers, c)
		}
	}

	// return cats slice and nil for the error
	return customers, nil
}