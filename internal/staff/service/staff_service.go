package service

import (
	"context"
	"enigmanations/eniqlo-store/internal/staff"
	"enigmanations/eniqlo-store/internal/staff/repository"
)

type StaffService interface {
	FindById(ctx context.Context, id int) (*staff.Staff, error)
	Register(ctx context.Context, s *staff.Staff) (*staff.Staff, error)
}

type staffService struct {
	repo repository.StaffRepository
}

func NewStaffService(repo repository.StaffRepository) StaffService {
	return &staffService{repo: repo}
}

// FindById implements StaffService.
func (s *staffService) FindById(ctx context.Context, id int) (*staff.Staff, error) {
	panic("unimplemented")
}

// Register implements StaffService.
func (*staffService) Register(ctx context.Context, s *staff.Staff) (*staff.Staff, error) {
	panic("unimplemented")
}
