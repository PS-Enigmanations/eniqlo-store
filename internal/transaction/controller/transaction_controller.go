package controller

import (
	custErrs "enigmanations/eniqlo-store/internal/customer/errs"
	productErrs "enigmanations/eniqlo-store/internal/product/errs"
	"enigmanations/eniqlo-store/internal/transaction/errs"
	"enigmanations/eniqlo-store/internal/transaction/request"
	"enigmanations/eniqlo-store/internal/transaction/response"
	"enigmanations/eniqlo-store/internal/transaction/service"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TransactionController interface {
	Checkout(ctx *gin.Context)
	SearchTransaction(ctx *gin.Context)
}

type transactionController struct {
	Service service.TransactionService
}

func NewTransactionController(svc service.TransactionService) TransactionController {
	return &transactionController{Service: svc}
}

func (c *transactionController) Checkout(ctx *gin.Context) {
	var reqBody request.CheckoutRequest

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	err := validate.Struct(reqBody)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	checkoutCreated := <-c.Service.Create(&reqBody)
	if checkoutCreated.Error != nil {
		fmt.Println(checkoutCreated.Error)
		switch {
		case errors.Is(checkoutCreated.Error, custErrs.CustomerIsNotExists):
			ctx.AbortWithError(http.StatusNotFound, checkoutCreated.Error)
			break
		case errors.Is(checkoutCreated.Error, productErrs.ProductIsNotExists):
			ctx.AbortWithError(http.StatusNotFound, checkoutCreated.Error)
			break
		case errors.Is(checkoutCreated.Error, errs.PaidIsNotEnough):
			ctx.AbortWithError(http.StatusBadRequest, checkoutCreated.Error)
			break
		case errors.Is(checkoutCreated.Error, errs.ChangeIsNotRight):
			ctx.AbortWithError(http.StatusBadRequest, checkoutCreated.Error)
			break
		case errors.Is(checkoutCreated.Error, productErrs.StockIsNotEnough):
			ctx.AbortWithError(http.StatusBadRequest, checkoutCreated.Error)
			break
		case errors.Is(checkoutCreated.Error, productErrs.ProductIsNotAvailable):
			ctx.AbortWithError(http.StatusBadRequest, checkoutCreated.Error)
			break
		default:
			ctx.AbortWithError(http.StatusInternalServerError, checkoutCreated.Error)
			break
		}
		return
	}

	ctx.Status(http.StatusOK)
	return
}

func (c *transactionController) SearchTransaction(ctx *gin.Context) {
	var reqQueryParams request.TransactionGetAllQueryParams

	if err := ctx.ShouldBindQuery(&reqQueryParams); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	transactions, err := c.Service.GetAllByParams(&reqQueryParams)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	transactionShows := response.ToTransactionShows(transactions)
	transactionMappedResults := response.TransactionToTransactionGetAllResponse(transactionShows)

	ctx.JSON(http.StatusOK, transactionMappedResults)
	return
}
