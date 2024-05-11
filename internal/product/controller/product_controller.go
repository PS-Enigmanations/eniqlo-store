package controller

import (
	"enigmanations/eniqlo-store/internal/common/errs"
	"enigmanations/eniqlo-store/internal/product/request"
	"enigmanations/eniqlo-store/internal/product/response"
	"enigmanations/eniqlo-store/internal/product/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	SearchProducts(ctx *gin.Context)
	Index(ctx *gin.Context)
	CreateProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
}

type productController struct {
	Service service.ProductService
}

func NewProductController(svc service.ProductService) ProductController {
	return &productController{Service: svc}
}

func (c *productController) Index(ctx *gin.Context) {
	var reqQueryParams request.SearchProductQueryParams
	if err := ctx.ShouldBindQuery(&reqQueryParams); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	products := <-c.Service.GetProducts(&reqQueryParams)
	if products.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, products.Error)
		return
	}

	// Mapping data from service to response
	productShows := response.ToProductShows(products.Result)
	productMappedResults := response.ProductToSearchProductsResponse(productShows)

	ctx.JSON(http.StatusOK, productMappedResults)
}

func (c *productController) SearchProducts(ctx *gin.Context) {
	var reqQueryParams request.SearchProductQueryParams
	if err := ctx.ShouldBindQuery(&reqQueryParams); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	products := <-c.Service.SearchProducts(&reqQueryParams)
	if products.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, products.Error)
		return
	}

	// Mapping data from service to response
	productShows := response.ToProductShows(products.Result)
	productMappedResults := response.ProductToSearchProductsResponse(productShows)

	ctx.JSON(http.StatusOK, productMappedResults)
}

func (c *productController) CreateProduct(ctx *gin.Context) {
	var reqBody request.ProductRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productCreated := <-c.Service.SaveProduct(&reqBody)
	if productCreated.Error != nil {
		if productCreated.Error == errs.ErrImageUrlInvalid {
			ctx.AbortWithError(http.StatusBadRequest, productCreated.Error)
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, productCreated.Error)
		return
	}

	productCreatedMappedResult := response.ProductToProductCreateResponse(productCreated.Result)
	ctx.JSON(http.StatusCreated, productCreatedMappedResult)
}

func (c *productController) UpdateProduct(ctx *gin.Context) {
	var reqBody request.ProductRequest
	reqBody.Id = ctx.Param("id")

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productSaved := <-c.Service.SaveProduct(&reqBody)
	if productSaved.Error != nil {
		if productSaved.Error == errs.ErrProductNotFound {
			ctx.AbortWithError(http.StatusNotFound, productSaved.Error)
			return
		}
		if productSaved.Error == errs.ErrImageUrlInvalid {
			ctx.AbortWithError(http.StatusBadRequest, productSaved.Error)
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, productSaved.Error)
		return

	}

	ctx.Status(http.StatusOK)
}

func (c *productController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.Service.DeleteProduct(id)
	if err != nil {
		if err == errs.ErrProductNotFound {
			ctx.AbortWithError(http.StatusNotFound, err)
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}
