package service

import (
	"context"

	"enigmanations/eniqlo-store/internal/customer"
	"enigmanations/eniqlo-store/internal/customer/repository"
	"enigmanations/eniqlo-store/internal/customer/request"
	"enigmanations/eniqlo-store/internal/customer/errs"
	commonErrs "enigmanations/eniqlo-store/internal/common/errs"
	"enigmanations/eniqlo-store/util"
	"enigmanations/eniqlo-store/pkg/uuid"
	"enigmanations/eniqlo-store/pkg/country"
	"strings"
	"github.com/jackc/pgx/v5/pgxpool"
	"fmt"
)

type CustomerService interface {
	GetAllByParams(p *request.CustomerGetAllQueryParams) ([]*customer.Customer, error)
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

func (svc *customerService) GetAllByParams(p *request.CustomerGetAllQueryParams) ([]*customer.Customer, error) {
	repo := svc.repo

	customers, err := repo.Customer.GetAllByParams(svc.context, p)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (svc *customerService) Create(payload *request.CustomerRegisterRequest) <-chan util.Result[*customer.Customer] {
	repo := svc.repo
	isPhoneNumberValid := false
	result := make(chan util.Result[*customer.Customer])

	go func() {
		id := uuid.New()
		model := customer.Customer{
			Id:      		id,
			Name:			payload.Name,
			PhoneNumber: 	payload.PhoneNumber,
		}

		for _, country := range country.Countries {
			if strings.HasPrefix(payload.PhoneNumber, fmt.Sprintf("%s%s", "+", country)) {
				isPhoneNumberValid = true
				break
			}
		}
		if !isPhoneNumberValid {
			result <- util.Result[*customer.Customer]{
				Error: commonErrs.InvalidPhoneNumber,
			}
			return
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