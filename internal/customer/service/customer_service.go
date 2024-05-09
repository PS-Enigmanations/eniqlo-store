package service

import (
	"context"

	"enigmanations/eniqlo-store/internal/customer"
	"enigmanations/eniqlo-store/internal/customer/repository"
	"enigmanations/eniqlo-store/internal/customer/request"
	"enigmanations/eniqlo-store/util"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerService interface {
	Create(p *request.CustomerRegisterRequest) <-chan util.Result[*customer.Customer]
}

type CustomerDependency struct {
	Customer      repository.CustomerRepository
}

type customerService struct {
	repo    *CustomerDependency
	pool    *pgxpool.Pool
	context context.Context
}

// NewService creates an API service.
func NewCatService(ctx context.Context, pool *pgxpool.Pool, repo *CustomerDependency) CustomerService {
	return &customerService{repo: repo, pool: pool, context: ctx}
}

func (svc *customerService) Create(payload *request.CustomerRegisterRequest) <-chan util.Result[*customer.Customer] {
	return nil
}