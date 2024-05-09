package repository

import (
	"context"
	"enigmanations/eniqlo-store/internal/staff"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type StaffRepository interface {
	FindById(ctx context.Context, id int) (*staff.Staff, error)
	FindByPhoneNumber(ctx context.Context, phoneNumber string) (*staff.Staff, error)
	Save(ctx context.Context, s *staff.Staff) (*staff.Staff, error)
	Login(ctx context.Context, phoeNumber string, password string) (*staff.Staff, error)
}

type staffRepository struct {
	pool *pgxpool.Pool
}

func NewStaffRepository(pool *pgxpool.Pool) StaffRepository {
	return &staffRepository{pool: pool}
}

func (r *staffRepository) FindById(ctx context.Context, id int) (*staff.Staff, error) {
	staff := &staff.Staff{}
	query := "SELECT id, name, phone_number FROM users WHERE id = $1"
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

func (r *staffRepository) FindByPhoneNumber(ctx context.Context, phoneNumber string) (*staff.Staff, error) {
	staff := &staff.Staff{}
	query := "SELECT id, name, phone_number, password FROM users WHERE phone_number = $1"
	err := r.pool.QueryRow(ctx, query, phoneNumber).Scan(
		&staff.ID,
		&staff.Name,
		&staff.PhoneNumber,
		&staff.Password,
	)
	if err != nil {
		return nil, err
	}
	return staff, nil
}

func (r *staffRepository) Login(ctx context.Context, phoneNumber string, password string) (*staff.Staff, error) {
	staff := &staff.Staff{}
	log.Println(phoneNumber, password)
	query := "SELECT id, name, phone_number FROM users WHERE phone_number = $1 AND password = $2"
	err := r.pool.QueryRow(ctx, query, phoneNumber, password).Scan(
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
	query := "INSERT INTO users (id, name, phone_number, password) VALUES ($1, $2, $3, $4) RETURNING id, name, phone_number"
	err := r.pool.QueryRow(ctx, query, s.ID, s.Name, s.PhoneNumber, s.Password).Scan(
		&staff.ID,
		&staff.Name,
		&staff.PhoneNumber,
	)
	if err != nil {
		return nil, err
	}
	return staff, nil

}
