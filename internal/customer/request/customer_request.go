package request

type CustomerRegisterRequest struct {
	Name 	   	string `json:"name" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=16,startswith=+"`
}

type CustomerGetAllQueryParams struct {
	PhoneNumber string `form:"phoneNumber"`
    Name        string `form:"name"`
}