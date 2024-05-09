package service

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"enigmanations/eniqlo-store/internal/transaction/request"
	"enigmanations/eniqlo-store/internal/transaction/repository"
)

type TransactionService interface {
	Create(p *request.CheckoutRequest) error
}

type TransactionDependency struct {
	Transaction      repository.TransactionRepository
}

type transactionService struct {
	repo    *TransactionDependency
	pool    *pgxpool.Pool
	context context.Context
}

func NewTransactionService(ctx context.Context, pool *pgxpool.Pool, repo *TransactionDependency) TransactionService {
	return &transactionService{repo: repo, pool: pool, context: ctx}
}


func (svc *transactionService) Create(p *request.CheckoutRequest) error {
	return nil
}