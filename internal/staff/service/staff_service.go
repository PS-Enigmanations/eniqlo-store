package service

import (
	"enigmanations/eniqlo-store/internal/staff"
	"enigmanations/eniqlo-store/internal/staff/repository"
	"enigmanations/eniqlo-store/internal/staff/request"
	"enigmanations/eniqlo-store/pkg/bcrypt"

	"github.com/gin-gonic/gin"
)

type StaffService interface {
	FindById(ctx *gin.Context, id int) (*staff.Staff, error)
	Register(ctx *gin.Context, req request.StaffRegisterRequest) (*staff.Staff, error)
}

type staffService struct {
	repo repository.StaffRepository
}

func NewStaffService(repo repository.StaffRepository) StaffService {
	return &staffService{repo: repo}
}

// FindById implements StaffService.
func (service *staffService) FindById(ctx *gin.Context, id int) (*staff.Staff, error) {
	panic("unimplemented")
}

// Register implements StaffService.
func (service *staffService) Register(ctx *gin.Context, req request.StaffRegisterRequest) (*staff.Staff, error) {
	hashedPassword, err := bcrypt.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	model := staff.Staff{
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    hashedPassword,
	}

	staff, err := service.repo.Save(ctx.Request.Context(), &model)
	if err != nil {
		return nil, err
	}

	return staff, nil
}
