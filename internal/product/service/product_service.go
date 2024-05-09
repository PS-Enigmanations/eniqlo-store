package service

import (
	"context"
	"enigmanations/eniqlo-store/internal/product"
	"enigmanations/eniqlo-store/internal/product/repository"
	"enigmanations/eniqlo-store/internal/product/request"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductService interface {
	SearchProducts(p *request.SearchProductQueryParams) ([]*product.Product, error)
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

func (svc *productService) SearchProducts(p *request.SearchProductQueryParams) ([]*product.Product, error) {
	repo := svc.repo

	products, err := repo.Product.SearchProducts(svc.context, p)
	if err != nil {
		return nil, err
	}

	return products, nil
}
