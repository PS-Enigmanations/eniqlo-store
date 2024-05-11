package service

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"enigmanations/eniqlo-store/internal/transaction"
	"enigmanations/eniqlo-store/internal/transaction/request"
	"enigmanations/eniqlo-store/internal/transaction/repository"
	custRepository "enigmanations/eniqlo-store/internal/customer/repository"
	productRepository "enigmanations/eniqlo-store/internal/product/repository"
	custErrs "enigmanations/eniqlo-store/internal/customer/errs"
	productErrs "enigmanations/eniqlo-store/internal/product/errs"
	"enigmanations/eniqlo-store/internal/transaction/errs"
	"enigmanations/eniqlo-store/pkg/validate"
	"enigmanations/eniqlo-store/util"
	"enigmanations/eniqlo-store/pkg/uuid"
)

type TransactionService interface {
	Create(p *request.CheckoutRequest) <-chan util.Result[interface{}]
	GetAllByParams(p *request.TransactionGetAllQueryParams) ([]*transaction.Transaction, error)
}

type TransactionDependency struct {
	Transaction      repository.TransactionRepository
	Product      	 productRepository.ProductRepository
	Customer      	 custRepository.CustomerRepository
}

type transactionService struct {
	repo    *TransactionDependency
	pool    *pgxpool.Pool
	context context.Context
}

func NewTransactionService(ctx context.Context, pool *pgxpool.Pool, repo *TransactionDependency) TransactionService {
	return &transactionService{repo: repo, pool: pool, context: ctx}
}


func (svc *transactionService) Create(p *request.CheckoutRequest) <-chan util.Result[interface{}] {
	repo := svc.repo
	result := make(chan util.Result[interface{}])
	go func() {
		customerFound, err := repo.Customer.FindById(svc.context, p.CustomerId)
		if customerFound == nil {
			result <- util.Result[interface{}]{
				Error: custErrs.CustomerIsNotExists,
			}
			return
		}

		if err != nil {
			result <- util.Result[interface{}]{
				Error: err,
			}
			return
		}

		total := 0

		var details []transaction.ProductDetail

		for _, detail := range p.ProductDetails {
			validateUuid := validate.IsValidUUID(detail.ProductId)

			if !validateUuid {
				result <- util.Result[interface{}]{
					Error: productErrs.ProductIsNotExists,
				}
				return
			}

			productExists, err := repo.Product.FindById(svc.context, detail.ProductId)
			if err != nil {
				result <- util.Result[interface{}]{
					Error: err,
				}
				return
			}

			if productExists == nil {
				result <- util.Result[interface{}]{
					Error: productErrs.ProductIsNotExists,
				}
				return
			}

			if productExists.Stock <= detail.Quantity {
				result <- util.Result[interface{}]{
					Error: productErrs.StockIsNotEnough,
				}
				return
			}

			d := transaction.ProductDetail{
				ProductId: detail.ProductId,
				Quantity: detail.Quantity,
			}

			details = append(details, d)

			total += int(productExists.Price * float64(detail.Quantity))
		}

		if float64(total) > p.Paid {
			result <- util.Result[interface{}]{
				Error: errs.PaidIsNotEnough,
			}
			return
		}

		validChange := p.Paid - float64(total)
		if validChange != p.Change {
			result <- util.Result[interface{}]{
				Error: errs.ChangeIsNotRight,
			}
			return
		}

		id := uuid.New()
		trx := transaction.Transaction{
			TransactionId:  id,
			CustomerId:		p.CustomerId,
			Paid: 			float64(p.Paid),
			Change: 		float64(p.Change),
		}
		newTrx, err := repo.Transaction.Save(svc.context, trx, float64(total))
		if err != nil {
			result <- util.Result[interface{}]{
				Error: err,
			}
			return
		}

		err = repo.Transaction.SaveDetails(svc.context, details, newTrx.TransactionId)
		if err != nil {
			result <- util.Result[interface{}]{
				Error: err,
			}
			return
		}

		err = repo.Product.UpdateStocks(svc.context, details)
		if err != nil {
			result <- util.Result[interface{}]{
				Error: err,
			}
			return
		}

		result <- util.Result[interface{}]{}
		close(result)
	}()
	
	return result
}

func (svc *transactionService) GetAllByParams(p *request.TransactionGetAllQueryParams) ([]*transaction.Transaction, error) {
	repo := svc.repo

	transactions, err := repo.Transaction.GetAllByParams(svc.context, p)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}