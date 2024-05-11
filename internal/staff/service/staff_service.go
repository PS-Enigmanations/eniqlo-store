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
	"enigmanations/eniqlo-store/util"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type StaffService interface {
	Login(ctx *gin.Context, req request.StaffLoginRequest) <-chan util.Result[*staff.Staff]
	Register(ctx *gin.Context, req request.StaffRegisterRequest) <-chan util.Result[*staff.Staff]
	GenerateAccessToken(staff *staff.Staff) (string, error)
}

type staffService struct {
	repo repository.StaffRepository
}

func NewStaffService(repo repository.StaffRepository) StaffService {
	return &staffService{repo: repo}
}

// Login implements StaffService.
func (service *staffService) Login(ctx *gin.Context, req request.StaffLoginRequest) <-chan util.Result[*staff.Staff] {
	result := make(chan util.Result[*staff.Staff])
	go func() {
		isPhoneNumberValid := country.IsValidPhoneCountryCode(req.PhoneNumber)
		if !isPhoneNumberValid {
			result <- util.Result[*staff.Staff]{
				Error: errors.New("invalid phone number"),
			}
			staffFound, err := service.repo.FindByPhoneNumber(ctx.Request.Context(), req.PhoneNumber)
			if err != nil {
				result <- util.Result[*staff.Staff]{
					Error: err,
				}
				return
			}
			if !bcrypt.CheckPasswordHash(req.Password, staffFound.Password) {
				result <- util.Result[*staff.Staff]{
					Error: errors.New("invalid password"),
				}
			}
			result <- util.Result[*staff.Staff]{
				Result: staffFound,
			}
		}
	}()
	return result
}

// Register implements StaffService.
func (service *staffService) Register(ctx *gin.Context, req request.StaffRegisterRequest) <-chan util.Result[*staff.Staff] {
	result := make(chan util.Result[*staff.Staff])
	go func() {
		hashedPassword, err := bcrypt.HashPassword(req.Password)
		if err != nil {
			result <- util.Result[*staff.Staff]{
				Error: err,
			}
			return
		}

		isPhoneNumberValid := country.IsValidPhoneCountryCode(req.PhoneNumber)
		if !isPhoneNumberValid {
			result <- util.Result[*staff.Staff]{
				Error: errors.New("invalid phone number"),
			}
			return
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
			result <- util.Result[*staff.Staff]{
				Error: err,
			}
			return
		}
		if staffFound != nil {
			result <- util.Result[*staff.Staff]{
				Error: errs.UserExist,
			}
			return
		}

		staffCreated, err := service.repo.Save(ctx.Request.Context(), &model)
		if err != nil {
			result <- util.Result[*staff.Staff]{
				Error: err,
			}
			return
		}

		result <- util.Result[*staff.Staff]{
			Result: staffCreated,
		}
	}()

	return result
}

func (service *staffService) GenerateAccessToken(staff *staff.Staff) (string, error) {
	token, err := jwt.GenerateAccessToken(staff.ID, staff)
	if err != nil {
		return "", errors.New("failed to generate access token")
	}

	return token, nil
}
