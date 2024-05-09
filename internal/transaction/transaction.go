package transaction

type Transaction struct {
	TransactionId  string          `json:"transactionId"`
	CustomerId     string          `json:"customerId"`
	ProductDetails []*ProductDetail `json:"productDetails"`
	Paid           float64             `json:"paid"`
	Change         float64             `json:"change"`
}

type ProductDetail struct {
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity"`
}