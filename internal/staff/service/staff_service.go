package service

import (
	"enigmanations/eniqlo-store/internal/staff"
	"enigmanations/eniqlo-store/internal/staff/errs"
	"enigmanations/eniqlo-store/internal/staff/repository"
	"enigmanations/eniqlo-store/internal/staff/request"
	"enigmanations/eniqlo-store/pkg/bcrypt"
	"enigmanations/eniqlo-store/pkg/country"
	"enigmanations/eniqlo-store/pkg/jwt"
	"enigmanations/eniqlo-store/pkg/uuid"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type StaffService interface {
	Login(ctx *gin.Context, req request.StaffLoginRequest) (*staff.Staff, error)
	Register(ctx *gin.Context, req request.StaffRegisterRequest) (*staff.Staff, error)
	GenerateAccessToken(staff *staff.Staff) (string, error)
}

type staffService struct {
	repo repository.StaffRepository
}

func NewStaffService(repo repository.StaffRepository) StaffService {
	return &staffService{repo: repo}
}

// Login implements StaffService.
func (service *staffService) Login(ctx *gin.Context, req request.StaffLoginRequest) (*staff.Staff, error) {
	isPhoneNumberValid := country.IsValidPhoneCountryCode(req.PhoneNumber)
	if !isPhoneNumberValid {
		return nil, errors.New("invalid phone number")
	}

	staff, err := service.repo.FindByPhoneNumber(ctx.Request.Context(), req.PhoneNumber)
	if err != nil {
		return nil, err
	}
	if !bcrypt.CheckPasswordHash(req.Password, staff.Password) {
		return nil, errors.New("invalid password")
	}

	return staff, nil
}

// Register implements StaffService.
func (service *staffService) Register(ctx *gin.Context, req request.StaffRegisterRequest) (*staff.Staff, error) {
	hashedPassword, err := bcrypt.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	isPhoneNumberValid := country.IsValidPhoneCountryCode(req.PhoneNumber)
	if !isPhoneNumberValid {
		return nil, errors.New("invalid phone number")
	}

	staffId := uuid.New()
	model := staff.Staff{
		ID:          staffId,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    hashedPassword,
	}

	staffFound, err := service.repo.FindByPhoneNumber(ctx.Request.Context(), req.PhoneNumber)
	if err != nil {
		return nil, err
	}

	if staffFound != nil {
		return nil, errs.UserExist
	}

	staff, err := service.repo.Save(ctx.Request.Context(), &model)
	if err != nil {
		return nil, err
	}

	return staff, nil
}

func (service *staffService) GenerateAccessToken(staff *staff.Staff) (string, error) {
	token, err := jwt.GenerateAccessToken(staff.ID, staff)
	if err != nil {
		return "", errors.New("failed to generate access token")
	}

	return token, nil
}
