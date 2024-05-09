package router_v1

import (
	"context"
	"enigmanations/eniqlo-store/internal/staff/controller"
	"enigmanations/eniqlo-store/internal/staff/repository"
	"enigmanations/eniqlo-store/internal/staff/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

type StaffRouter struct {
	Controller controller.StaffController
}

func NewStaffRouter(ctx context.Context, pool *pgxpool.Pool) *StaffRouter {
	staffRepository := repository.NewStaffRepository(pool)

	staffService := service.NewStaffService(
		staffRepository,
	)

	return &StaffRouter{
		Controller: controller.NewStaffController(staffService),
	}
}
