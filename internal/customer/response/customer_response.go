package response

import (
	"enigmanations/eniqlo-store/internal/customer"
)

const CustomerRegisterSuccMessage = "User registered successfully"

type CustomerRegisterResponse struct {
	Message string        	`json:"message"`
	Data    CustomerShow 	`json:"data"`
}

// Create response
type CustomerShow struct {
	Id       	  string `json:"userId"`
	PhoneNumber   string `json:"phoneNumber"`
	Name          string `json:"name"`
}

func CustomerToCustomerRegisterResponse(c customer.Customer) *CustomerRegisterResponse {
	return &CustomerRegisterResponse{
		Message: CustomerRegisterSuccMessage,
		Data: CustomerShow{
			Id:   c.Id,
			PhoneNumber: c.PhoneNumber,
			Name: c.Name,
		},
	}
}