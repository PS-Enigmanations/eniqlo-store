package controller

import (
	"enigmanations/eniqlo-store/internal/customer/service"
	"enigmanations/eniqlo-store/internal/customer/request"
	"enigmanations/eniqlo-store/internal/customer/response"
	"enigmanations/eniqlo-store/internal/customer/errs"
	commonErrs "enigmanations/eniqlo-store/internal/common/errs"
	"net/http"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CustomerController interface {
	Register(ctx *gin.Context)
	SearchCustomer(ctx *gin.Context)
}

type customerController struct {
	Service service.CustomerService
}

func NewCustomerController(svc service.CustomerService) CustomerController {
	return &customerController{Service: svc}
}

func (c *customerController) Register(ctx *gin.Context) {
	var reqBody request.CustomerRegisterRequest

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

	// send data to service layer to further process (create record)
	customerCreated := <-c.Service.Create(&reqBody)
	if customerCreated.Error != nil {
		switch {
		case errors.Is(customerCreated.Error, errs.CustomerExist):
			ctx.AbortWithError(http.StatusConflict, customerCreated.Error)
			break
		case errors.Is(customerCreated.Error, commonErrs.InvalidPhoneNumber):
			ctx.AbortWithError(http.StatusBadRequest,customerCreated.Error)
			break
		default:
			ctx.AbortWithError(http.StatusInternalServerError, customerCreated.Error)
			break
		}
		return
	}

	// Mapping data from service to response
	customerCreatedMappedResult := response.CustomerToCustomerRegisterResponse(*customerCreated.Result)
	ctx.JSON(http.StatusCreated, customerCreatedMappedResult)
	return
}

func (c *customerController) SearchCustomer(ctx *gin.Context) {
	var reqQueryParams request.CustomerGetAllQueryParams

	if err := ctx.ShouldBindQuery(&reqQueryParams); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	customers, err := c.Service.GetAllByParams(&reqQueryParams)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Mapping data from service to response
	customerShows := response.ToCustomerShows(customers)
	customerMappedResults := response.CustomerToCustomerGetAllResponse(customerShows)

	ctx.JSON(http.StatusOK, customerMappedResults)
	return
}
