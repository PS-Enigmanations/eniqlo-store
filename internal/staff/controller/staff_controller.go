package controller

import (
	"enigmanations/eniqlo-store/internal/staff/request"
	"enigmanations/eniqlo-store/internal/staff/response"
	"enigmanations/eniqlo-store/internal/staff/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StaffController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type staffController struct {
	Service service.StaffService
}

func NewStaffController(service service.StaffService) StaffController {
	return &staffController{
		Service: service,
	}
}

func (controller *staffController) Register(c *gin.Context) {
	var requestBody request.StaffRegisterRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	staffCreated, err := controller.Service.Register(c, requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response.StaffResponse{
		ID:          staffCreated.ID,
		Name:        staffCreated.Name,
		PhoneNumber: staffCreated.PhoneNumber,
	})
}

func (*staffController) Login(c *gin.Context) {
	panic("implement me")
}
