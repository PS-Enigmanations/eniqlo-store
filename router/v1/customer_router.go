package router_v1

import (
	"context"
	"enigmanations/eniqlo-store/internal/customer/controller"
	"enigmanations/eniqlo-store/internal/customer/repository"
	"enigmanations/eniqlo-store/internal/customer/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerRouter struct {
	Controller controller.CustomerController
}

func NewCustomerRouter(ctx context.Context, pool *pgxpool.Pool) *CustomerRouter {
	customerRepository := repository.NewCustomerRepository(pool)

	customerService := service.NewCustomerService(
		ctx,
		pool,
		&service.CustomerDependency{
			Customer:      customerRepository,
		},
	)

	return &CustomerRouter{
		Controller: controller.NewCustomerController(customerService),
	}
}
