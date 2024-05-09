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
}

type productController struct {
	Service service.ProductService
}

func NewProductController(svc service.ProductService) ProductController {
	return &productController{Service: svc}
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
	return
}
