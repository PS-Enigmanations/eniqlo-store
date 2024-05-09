package response

import (
	"enigmanations/eniqlo-store/internal/transaction"
)

type TransactionGetAllResponse struct {
	Message string  	  	`json:"message"`
	Data    []TransactionShow 	`json:"data"`
}

type TransactionShow transaction.Transaction

func ToTransactionShows(c []*transaction.Transaction) []TransactionShow {
	list := make([]TransactionShow, len(c))
	for i, item := range c {
		list[i] = TransactionShow{
			TransactionId:      item.TransactionId,
			CustomerId:       	item.CustomerId,
			ProductDetails: 	item.ProductDetails,
			Paid:        		item.Paid,
			Change: 			item.Change,
		}
	}

	return list
}

func TransactionToTransactionGetAllResponse(data []TransactionShow) *TransactionGetAllResponse {
	return &TransactionGetAllResponse{
		Message: "success",
		Data:    data,
	}
}