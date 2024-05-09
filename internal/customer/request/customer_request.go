package request

type CustomerRegisterRequest struct {
	Name 	   	string  `json:"name" validate:"required"`
	PhoneNumber string  `json:"phoneNumber" validate:"required"`
}