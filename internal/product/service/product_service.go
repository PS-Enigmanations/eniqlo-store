package service

import (
	"context"
	"enigmanations/eniqlo-store/internal/product"
	"enigmanations/eniqlo-store/internal/product/repository"
	"enigmanations/eniqlo-store/internal/product/request"
	"enigmanations/eniqlo-store/util"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductService interface {
	SearchProducts(p *request.SearchProductQueryParams) <-chan util.Result[[]*product.Product]
	GetProducts(p *request.SearchProductQueryParams) <-chan util.Result[[]*product.Product]
}

type ProductDependency struct {
	Product repository.ProductRepository
}

type productService struct {
	repo    *ProductDependency
	pool    *pgxpool.Pool
	context context.Context
}

func NewProductService(ctx context.Context, pool *pgxpool.Pool, repo *ProductDependency) ProductService {
	return &productService{
		repo:    repo,
		pool:    pool,
		context: ctx,
	}
}

func (svc *productService) SearchProducts(p *request.SearchProductQueryParams) <-chan util.Result[[]*product.Product] {
	repo := svc.repo

	result := make(chan util.Result[[]*product.Product])
	go func() {
		products, err := repo.Product.SearchProducts(svc.context, p, true)
		if err != nil {
			result <- util.Result[[]*product.Product]{
				Error: err,
			}
			return
		}

		result <- util.Result[[]*product.Product]{
			Result: products,
		}
		close(result)
	}()

	return result
}

func (svc *productService) GetProducts(p *request.SearchProductQueryParams) <-chan util.Result[[]*product.Product] {
	repo := svc.repo

	result := make(chan util.Result[[]*product.Product])
	go func() {
		products, err := repo.Product.SearchProducts(svc.context, p, false)
		if err != nil {
			result <- util.Result[[]*product.Product]{
				Error: err,
			}
			return
		}

		result <- util.Result[[]*product.Product]{
			Result: products,
		}
		close(result)
	}()

	return result
}
