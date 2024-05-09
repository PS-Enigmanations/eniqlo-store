package controller

import (
	"enigmanations/eniqlo-store/internal/customer/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerController interface {
	CustomerRegister(ctx *gin.Context)
}

type customerController struct {
	Service service.CustomerService
}

func NewCustomerController(svc service.CustomerService) CustomerController {
	return &customerController{Service: svc}
}

func (c *customerController) CustomerRegister(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
	return
}
