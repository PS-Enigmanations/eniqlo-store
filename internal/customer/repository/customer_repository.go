package repository

import (
	"context"
	"enigmanations/eniqlo-store/internal/customer"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerRepository interface {
	Save(ctx context.Context, model customer.Customer) (*customer.Customer, error)
	FindByPhoneNumber(ctx context.Context, phoneNumber string) (*customer.Customer, error)
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