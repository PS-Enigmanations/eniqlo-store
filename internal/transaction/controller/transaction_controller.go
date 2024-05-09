package controller

import (
	"enigmanations/eniqlo-store/internal/transaction/service"
	"github.com/gin-gonic/gin"
	"net/http"
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
	ctx.Status(http.StatusOK)
	return
}

func (c *transactionController) SearchTransaction(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
	return
}