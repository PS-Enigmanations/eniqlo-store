package controller

import (
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
	// send data to service layer to further process (create record)
	productCreated, err := c.Service.SaveProduct(&reqBody)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Mapping data from service to response
	productCreatedMappedResult := response.ProductToProductCreateResponse(productCreated)
	ctx.JSON(http.StatusCreated, productCreatedMappedResult)
}

func (c *productController) UpdateProduct(ctx *gin.Context) {
	var reqBody request.ProductRequest

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reqBody.Id = ctx.Param("id")
	_, err := c.Service.SaveProduct(&reqBody)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product updated successfully",
	})
}
