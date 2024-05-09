package customer

import (
	"time"
)

type Customer struct {
	Id         	string  `json:"id"`
	Name 	   	string  `json:"name"`
	PhoneNumber string  `json:"phone_number"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}