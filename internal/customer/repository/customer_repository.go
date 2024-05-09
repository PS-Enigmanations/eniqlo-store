package repository

import (
	"context"
	"enigmanations/eniqlo-store/internal/customer"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerRepository interface {
	Save(ctx context.Context, model customer.Customer) (*customer.Customer, error)
}

type Database struct {
	pool *pgxpool.Pool
}

func NewCustomerRepository(pool *pgxpool.Pool) CustomerRepository {
	return &Database{
		pool: pool,
	}
}

func (db *Database) Save(ctx context.Context, model customer.Customer) (*customer.Customer, error) {
	return nil, nil
}