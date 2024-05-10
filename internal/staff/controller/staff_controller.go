package controller

import (
	"enigmanations/eniqlo-store/internal/staff/errs"
	"enigmanations/eniqlo-store/internal/staff/request"
	"enigmanations/eniqlo-store/internal/staff/response"
	"enigmanations/eniqlo-store/internal/staff/service"
	"net/http"
	"errors"

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
		if errors.Is(err, errs.UserExist) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "invalid phone number" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := controller.Service.GenerateAccessToken(staffCreated)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"data": response.StaffResponse{
			ID:          staffCreated.ID,
			Name:        staffCreated.Name,
			PhoneNumber: staffCreated.PhoneNumber,
			AccessToken: accessToken,
		},
	})
}

func (controller *staffController) Login(c *gin.Context) {
	var requestBody request.StaffLoginRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	staffLoggedIn, err := controller.Service.Login(c, requestBody)
	if err != nil {
		if err.Error() == "invalid phone number" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := controller.Service.GenerateAccessToken(staffLoggedIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Message": "User logged in successfully",
		"Data": response.StaffResponse{
			ID:          staffLoggedIn.ID,
			Name:        staffLoggedIn.Name,
			PhoneNumber: staffLoggedIn.PhoneNumber,
			AccessToken: accessToken,
		},
	})
}
