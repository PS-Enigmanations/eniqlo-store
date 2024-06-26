package router_v1

import (
	"context"
	"enigmanations/eniqlo-store/internal/transaction/controller"
	"enigmanations/eniqlo-store/internal/transaction/repository"
	custRepository "enigmanations/eniqlo-store/internal/customer/repository"
	productRepository "enigmanations/eniqlo-store/internal/product/repository"
	"enigmanations/eniqlo-store/internal/transaction/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionRouter struct {
	Controller controller.TransactionController
}

func NewTransactionRouter(ctx context.Context, pool *pgxpool.Pool) *TransactionRouter {
	trxRepository := repository.NewTransactionRepository(pool)
	customerRepository := custRepository.NewCustomerRepository(pool)
	productRepository := productRepository.NewProductRepository(pool)

	trxService := service.NewTransactionService(
		ctx,
		pool,
		&service.TransactionDependency{
			Transaction:  trxRepository,
			Customer:	  customerRepository,
			Product: productRepository,
		},
	)

	return &TransactionRouter{
		Controller: controller.NewTransactionController(trxService),
	}
}
