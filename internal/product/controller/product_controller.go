package controller

import (
	"enigmanations/eniqlo-store/internal/product/request"
	"enigmanations/eniqlo-store/internal/product/response"
	"enigmanations/eniqlo-store/internal/product/service"
	"fmt"
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

	products, err := c.Service.SearchProducts(&reqQueryParams)
	if err != nil {
		fmt.Print("err", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Mapping data from service to response
	productShows := response.ToProductShows(products)
	productMappedResults := response.ProductToSearchProductsResponse(productShows)

	ctx.JSON(http.StatusOK, productMappedResults)
	return
}
