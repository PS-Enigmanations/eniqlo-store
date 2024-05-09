package service

import (
	"context"

	"enigmanations/eniqlo-store/internal/customer"
	"enigmanations/eniqlo-store/internal/customer/repository"
	"enigmanations/eniqlo-store/internal/customer/request"
	"enigmanations/eniqlo-store/internal/customer/errs"
	"enigmanations/eniqlo-store/util"
	"enigmanations/eniqlo-store/pkg/uuid"
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
func NewCustomerService(ctx context.Context, pool *pgxpool.Pool, repo *CustomerDependency) CustomerService {
	return &customerService{repo: repo, pool: pool, context: ctx}
}

func (svc *customerService) Create(payload *request.CustomerRegisterRequest) <-chan util.Result[*customer.Customer] {
	repo := svc.repo

	result := make(chan util.Result[*customer.Customer])

	go func() {
		id := uuid.New()
		model := customer.Customer{
			Id:      		id,
			Name:			payload.Name,
			PhoneNumber: 	payload.PhoneNumber,
		}

		// call FindByPhoneNumber if exists
		customerFound, err := repo.Customer.FindByPhoneNumber(svc.context, payload.PhoneNumber)
		if customerFound != nil {
			result <- util.Result[*customer.Customer]{
				Error: errs.CustomerExist,
			}
			return
		}

		// if error occur, return nil for the response as well as return the error
		if err != nil {
			result <- util.Result[*customer.Customer]{
				Error: err,
			}
			return
		}

		// call Create from repository/ datastore
		customerCreated, err := repo.Customer.Save(svc.context, model)

		// if error occur, return nil for the response as well as return the error
		if err != nil {
			result <- util.Result[*customer.Customer]{
				Error: err,
			}
			return
		}

		result <- util.Result[*customer.Customer]{
			Result: customerCreated,
		}
		close(result)
	}()

	return result
}