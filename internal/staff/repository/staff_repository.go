package repository

import (
	"context"
	"enigmanations/eniqlo-store/internal/staff"

	"github.com/jackc/pgx/v5/pgxpool"
)

type StaffRepository interface {
	FindById(ctx context.Context, id int) (*staff.Staff, error)
	Save(ctx context.Context, s *staff.Staff) (*staff.Staff, error)
}

type staffRepository struct {
	pool *pgxpool.Pool
}

func NewStaffRepository(pool *pgxpool.Pool) StaffRepository {
	return &staffRepository{pool: pool}
}

func (r *staffRepository) FindById(ctx context.Context, id int) (*staff.Staff, error) {
	staff := &staff.Staff{}
	query := "SELECT id, name, phone_number FROM staff WHERE id = $1"
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&staff.ID,
		&staff.Name,
		&staff.PhoneNumber,
	)
	if err != nil {
		return nil, err
	}
	return staff, nil
}

func (r *staffRepository) Save(ctx context.Context, s *staff.Staff) (*staff.Staff, error) {

	staff := &staff.Staff{} // Initialize the staff variable
	query := "INSERT INTO staff (name, phone_number, password) VALUES ($1, $2, $3) RETURNING id, name, phone_number"
	err := r.pool.QueryRow(ctx, query, s.Name, s.PhoneNumber, s.Password).Scan(
		&staff.ID,
		&staff.Name,
		&staff.PhoneNumber,
	)
	if err != nil {
		return nil, err
	}
	return staff, nil

}
