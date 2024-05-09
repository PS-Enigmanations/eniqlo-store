package router_v1

import (
	"context"
	"enigmanations/eniqlo-store/internal/product/controller"
	"enigmanations/eniqlo-store/internal/product/repository"
	"enigmanations/eniqlo-store/internal/product/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRouter struct {
	Controller controller.ProductController
}

func NewProductRouter(ctx context.Context, pool *pgxpool.Pool) *ProductRouter {
	productRepository := repository.NewProductRepository(pool)
	productService := service.NewProductService(ctx, pool, &service.ProductDependency{
		Product: productRepository,
	})

	return &ProductRouter{
		Controller: controller.NewProductController(productService),
	}
}
