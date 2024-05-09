package controller

import (
	"enigmanations/eniqlo-store/internal/transaction/service"
	"enigmanations/eniqlo-store/internal/transaction/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
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
		fmt.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
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

	ctx.Status(http.StatusOK)
	return
}