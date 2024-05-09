package request

type CustomerRegisterRequest struct {
	Name 	   	string  `json:"name" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=16,startswith=+"`
}