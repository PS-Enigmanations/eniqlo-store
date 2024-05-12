package service

import (
	"context"
	custErrs "enigmanations/eniqlo-store/internal/customer/errs"
	custRepository "enigmanations/eniqlo-store/internal/customer/repository"
	productErrs "enigmanations/eniqlo-store/internal/product/errs"
	productRepository "enigmanations/eniqlo-store/internal/product/repository"
	"enigmanations/eniqlo-store/internal/transaction"
	"enigmanations/eniqlo-store/internal/transaction/errs"
	"enigmanations/eniqlo-store/internal/transaction/repository"
	"enigmanations/eniqlo-store/internal/transaction/request"
	"enigmanations/eniqlo-store/pkg/uuid"
	"enigmanations/eniqlo-store/pkg/validate"
	"enigmanations/eniqlo-store/util"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionService interface {
	Create(p *request.CheckoutRequest) <-chan util.ResultErr
	GetAllByParams(p *request.TransactionGetAllQueryParams) ([]*transaction.Transaction, error)
}

type TransactionDependency struct {
	Transaction repository.TransactionRepository
	Product     productRepository.ProductRepository
	Customer    custRepository.CustomerRepository
}

type transactionService struct {
	repo    *TransactionDependency
	pool    *pgxpool.Pool
	context context.Context
}

func NewTransactionService(ctx context.Context, pool *pgxpool.Pool, repo *TransactionDependency) TransactionService {
	return &transactionService{repo: repo, pool: pool, context: ctx}
}

type productDetailsBatchRet struct {
	totalOrder int
	orderItems []transaction.ProductDetail
}

func (svc *transactionService) Create(p *request.CheckoutRequest) <-chan util.ResultErr {
	repo := svc.repo

	result := make(chan util.ResultErr)
	go func() {
		customerFound, err := repo.Customer.FindById(svc.context, p.CustomerId)
		if customerFound == nil {
			result <- util.ResultErr{
				Error: custErrs.CustomerIsNotExists,
			}
			return
		}

		if err != nil {
			result <- util.ResultErr{
				Error: err,
			}
			return
		}

		total := 0.0

		var details []transaction.ProductDetail

		for _, detail := range p.ProductDetails {
			validateUuid := validate.IsValidUUID(detail.ProductId)

			if !validateUuid {
				result <- util.ResultErr{
					Error: productErrs.ProductIsNotExists,
				}
				return
			}

			productExists, err := repo.Product.FindById(svc.context, detail.ProductId)
			if err != nil {
				result <- util.ResultErr{
					Error: err,
				}
				return
			}

			if productExists == nil {
				result <- util.ResultErr{
					Error: productErrs.ProductIsNotExists,
				}
				return
			}

			if productExists.Stock < detail.Quantity {
				result <- util.ResultErr{
					Error: productErrs.StockIsNotEnough,
				}
				return
			}

			if productExists.IsAvailable != true {
				result <- util.ResultErr{
					Error: productErrs.ProductIsNotAvailable,
				}
				return
			}

			d := transaction.ProductDetail{
				ProductId: detail.ProductId,
				Quantity:  detail.Quantity,
			}

			details = append(details, d)

			total += float64(productExists.Price * float64(detail.Quantity))
		}

		if total > p.Paid {
			result <- util.ResultErr{
				Error: errs.PaidIsNotEnough,
			}
			return
		}

		validChange := p.Paid - total
		if validChange != *p.Change {
			result <- util.ResultErr{
				Error: errs.ChangeIsNotRight,
			}
			return
		}

		id := uuid.New()
		trx := transaction.Transaction{
			TransactionId: id,
			CustomerId:    p.CustomerId,
			Paid:          float64(p.Paid),
			Change:        float64(*p.Change),
		}
		newTrx, err := repo.Transaction.Save(svc.context, trx, float64(total))
		if err != nil {
			result <- util.ResultErr{
				Error: err,
			}
			return
		}

		err = repo.Transaction.SaveDetails(svc.context, details, newTrx.TransactionId)
		if err != nil {
			result <- util.ResultErr{
				Error: err,
			}
			return
		}

		err = repo.Product.UpdateStocks(svc.context, details)
		if err != nil {
			result <- util.ResultErr{
				Error: err,
			}
			return
		}

		result <- util.ResultErr{Error: nil}
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
