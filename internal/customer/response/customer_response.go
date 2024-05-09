package response

import (
	"enigmanations/eniqlo-store/internal/customer"
)

const CustomerRegisterSuccMessage = "User registered successfully"

type CustomerRegisterResponse struct {
	Message string        	`json:"message"`
	Data    CustomerShow 	`json:"data"`
}

type CustomerGetAllResponse struct {
	Message string  	  	`json:"message"`
	Data    []CustomerShow 	`json:"data"`
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

func ToCustomerShows(c []*customer.Customer) []CustomerShow {
	list := make([]CustomerShow, len(c))
	for i, item := range c {
		list[i] = CustomerShow{
			Id:          item.Id,
			Name:        item.Name,
			PhoneNumber: item.PhoneNumber,
		}
	}

	return list
}

func CustomerToCustomerGetAllResponse(data []CustomerShow) *CustomerGetAllResponse {
	return &CustomerGetAllResponse{
		Message: "success",
		Data:    data,
	}
}