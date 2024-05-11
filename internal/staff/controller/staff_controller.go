package controller

import (
	"enigmanations/eniqlo-store/internal/staff/errs"
	"enigmanations/eniqlo-store/internal/staff/request"
	"enigmanations/eniqlo-store/internal/staff/response"
	"enigmanations/eniqlo-store/internal/staff/service"
	"errors"
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

	staffCreated := <-controller.Service.Register(c, requestBody)
	if staffCreated.Error != nil {
		if errors.Is(staffCreated.Error, errs.UserExist) {
			c.JSON(http.StatusConflict, gin.H{"error": staffCreated.Error.Error()})
			return
		}
		if errors.Is(staffCreated.Error, errs.ErrInvalidPhoneNumber) {
			c.JSON(http.StatusConflict, gin.H{"error": staffCreated.Error.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": staffCreated.Error.Error()})
		return
	}

	accessToken, err := controller.Service.GenerateAccessToken(staffCreated.Result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"data": response.StaffResponse{
			ID:          staffCreated.Result.ID,
			Name:        staffCreated.Result.Name,
			PhoneNumber: staffCreated.Result.PhoneNumber,
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

	staffLoggedIn := <-controller.Service.Login(c, requestBody)
	if staffLoggedIn.Error != nil {
		if errors.Is(staffLoggedIn.Error, errs.ErrInvalidPhoneNumber) {
			c.JSON(http.StatusBadRequest, gin.H{"error": staffLoggedIn.Error.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": staffLoggedIn.Error.Error()})
		return
	}

	accessToken, err := controller.Service.GenerateAccessToken(staffLoggedIn.Result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in successfully",
		"data": response.StaffResponse{
			ID:          staffLoggedIn.Result.ID,
			Name:        staffLoggedIn.Result.Name,
			PhoneNumber: staffLoggedIn.Result.PhoneNumber,
			AccessToken: accessToken,
		},
	})
}
