package service

import (
	"context"
	"enigmanations/eniqlo-store/internal/product"
	"enigmanations/eniqlo-store/internal/product/repository"
	"enigmanations/eniqlo-store/internal/product/request"
	"enigmanations/eniqlo-store/pkg/uuid"
	"enigmanations/eniqlo-store/util"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductService interface {
	SearchProducts(p *request.SearchProductQueryParams) <-chan util.Result[[]*product.Product]
	GetProducts(p *request.SearchProductQueryParams) <-chan util.Result[[]*product.Product]
	SaveProduct(p *request.ProductRequest) (*product.Product, error)
	DeleteProduct(id string) error
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

func (svc *productService) SaveProduct(p *request.ProductRequest) (*product.Product, error) {
	repo := svc.repo

	productId := uuid.New()
	if p.Id != "" {
		productId = p.Id
	}

	product := &product.Product{
		Id:          productId,
		Name:        p.Name,
		Sku:         p.Sku,
		Category:    product.Category(p.Category),
		ImageUrl:    p.ImageUrl,
		Notes:       p.Notes,
		Price:       p.Price,
		Stock:       p.Stock,
		Location:    p.Location,
		IsAvailable: p.IsAvailable,
	}

	product, err := repo.Product.SaveProduct(svc.context, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (svc *productService) DeleteProduct(id string) error {
	repo := svc.repo

	product, err := repo.Product.SearchProducts(svc.context, &request.SearchProductQueryParams{Id: id}, false)
	if err != nil {
		return err
	}

	if len(product) == 0 {
		return errors.New("product not found")
	}

	err = repo.Product.DeleteProduct(svc.context, id)
	if err != nil {
		return err
	}

	return nil
}
