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
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/sync/errgroup"
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

// batch orchestrates
func (svc *transactionService) readProductDetailsBatch(
	p *request.CheckoutRequest,
	batchSize int,
) util.Result[*productDetailsBatchRet] {
	repo := svc.repo

	var batches [][]request.ProductDetail
	var g errgroup.Group
	var mutex sync.Mutex

	var totalAccumulated = 0
	var orderItems []transaction.ProductDetail
	var errs error

	// Create batches chunks
	batchChunksCh := make(chan []request.ProductDetail)
	for i := 0; i < len(p.ProductDetails); i += batchSize {
		j := i + batchSize
		if j > len(p.ProductDetails) {
			j = len(p.ProductDetails)
		}

		batches = append(batches, p.ProductDetails[i:j])
	}

	// Function to process items
	readProductDetailsEntries := func() error {
		// Listen on `batchChunksCh` to see if there is any resource pending in it.
		for chunks := range batchChunksCh {
			for _, item := range chunks {
				// When it releases the lock, another goroutine can acquire it
				// and continue working with the slice.
				defer mutex.Unlock() // <- defers the execution until the func returns (will release the lock)

				// When a goroutine acquires the lock, it can safely add items
				// to the slice without worrying about race conditions.
				mutex.Lock()

				isValidateUuid := validate.IsValidUUID(item.ProductId)
				if !isValidateUuid {
					errs = productErrs.ProductIsNotExists
					return nil
				}

				productExists, err := repo.Product.FindById(svc.context, item.ProductId)
				if err != nil {
					errs = err
					return nil
				}
				if productExists == nil {
					errs = productErrs.ProductIsNotExists
					return nil
				}
				if productExists.Stock < item.Quantity {
					errs = productErrs.StockIsNotEnough
					return nil
				}

				orderItems = append(orderItems, transaction.ProductDetail{
					ProductId: item.ProductId,
					Quantity:  item.Quantity,
				})

				totalAccumulated += int(productExists.Price * float64(item.Quantity))
			}
		}

		return nil
	}

	// Start worker goroutines
	for range batches {
		g.Go(func() error {
			return readProductDetailsEntries()
		})
	}

	// Send items to be processed
	for _, item := range batches {
		batchChunksCh <- item
	}

	// Close the channel to signal that all items have been sent
	close(batchChunksCh)

	// Wait for all goroutines to finish
	if err := g.Wait(); err != nil {
		return util.Result[*productDetailsBatchRet]{
			Error: err,
		}
	}

	if errs != nil {
		return util.Result[*productDetailsBatchRet]{
			Error: errs,
		}
	}

	return util.Result[*productDetailsBatchRet]{
		Result: &productDetailsBatchRet{
			totalOrder: totalAccumulated,
			orderItems: orderItems,
		},
	}
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

		details := svc.readProductDetailsBatch(p, 2)
		if details.Error != nil {
			fmt.Printf("Result error %v %s", details.Error, "\n")
			result <- util.ResultErr{
				Error: details.Error,
			}
			return
		}

		if float64(details.Result.totalOrder) > p.Paid {
			result <- util.ResultErr{
				Error: errs.PaidIsNotEnough,
			}

			return
		}

		validChange := p.Paid - float64(details.Result.totalOrder)
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
		newTrx, err := repo.Transaction.Save(svc.context, trx, float64(details.Result.totalOrder))
		if err != nil {
			result <- util.ResultErr{
				Error: err,
			}
			return
		}

		err = repo.Transaction.SaveDetails(svc.context, details.Result.orderItems, newTrx.TransactionId)
		if err != nil {
			result <- util.ResultErr{
				Error: err,
			}
			return
		}

		err = repo.Product.UpdateStocks(svc.context, details.Result.orderItems)
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
