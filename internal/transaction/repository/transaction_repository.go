package repository

import (
	"enigmanations/eniqlo-store/internal/transaction"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionRepository interface {
	Save(ctx context.Context, model transaction.Transaction) (*transaction.Transaction, error)
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