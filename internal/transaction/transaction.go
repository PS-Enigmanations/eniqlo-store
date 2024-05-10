package transaction

import (
	"time"
)

type Transaction struct {
	TransactionId  string          `json:"transactionId"`
	CustomerId     string          `json:"customerId"`
	ProductDetails []*ProductDetail `json:"productDetails"`
	Paid           float64             `json:"paid"`
	Change         float64             `json:"change"`
	CreatedAt 	   time.Time 		`json:"createdAt"`
}

type ProductDetail struct {
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity"`
}