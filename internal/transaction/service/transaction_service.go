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
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/sync/errgroup"
)

type TransactionService interface {
	Create(p *request.CheckoutRequest) <-chan util.Result[interface{}]
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
	detail transaction.ProductDetail
}

func (svc *transactionService) readProductDetailsEntries(items []request.ProductDetail) error {
	repo := svc.repo

	total := 0
	var details []transaction.ProductDetail

	for _, item := range items {
		isValidateUuid := validate.IsValidUUID(item.ProductId)
		if !isValidateUuid {
			return productErrs.ProductIsNotExists
		}

		productExists, err := repo.Product.FindById(svc.context, item.ProductId)
		if err != nil {
			return err
		}
		if productExists == nil {
			return productErrs.ProductIsNotExists
		}
		if productExists.Stock < item.Quantity {
			return productErrs.StockIsNotEnough
		}

		detail := transaction.ProductDetail{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		}
		details = append(details, detail)
		total += int(productExists.Price * float64(detail.Quantity))
	}

	return nil
}

// batch orchestrates
func (svc *transactionService) readProductDetailsBatch(ps []request.ProductDetail, batchSize int) error {

	batches := make(chan []request.ProductDetail)
	var g errgroup.Group

	// Function to process items

	// Start worker goroutines
	// Listen on `jobsCh` to see if there is any resource pending in it.
	for batch := range batches {
		g.Go(func() error {
			return svc.readProductDetailsEntries(batch)
		})
	}

	// Send items to be processed and push into `batches` channel
	for i := 0; i < len(ps); i += batchSize {
		j := i + batchSize
		if j > len(ps) {
			j = len(ps)
		}

		// Send items to be processed
		batches <- ps[i:j]
	}

	// Close the channel to signal that all items have been sent
	close(batches)

	// Wait for all goroutines to finish
	if err := g.Wait(); err != nil {
		return fmt.Errorf("readProductDetailsBatch result error: %v\n", err)
	}

	return nil
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
			isValidateUuid := validate.IsValidUUID(detail.ProductId)
			if !isValidateUuid {
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

			if productExists.Stock < detail.Quantity {
				result <- util.Result[interface{}]{
					Error: productErrs.StockIsNotEnough,
				}
				return
			}

			d := transaction.ProductDetail{
				ProductId: detail.ProductId,
				Quantity:  detail.Quantity,
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
		if validChange != *p.Change {
			result <- util.Result[interface{}]{
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
