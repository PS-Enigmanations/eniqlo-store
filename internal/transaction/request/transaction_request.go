package request

type ProductDetail struct {
	ProductId string `json:"productId" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required,min=1"`
}

type CheckoutRequest struct {
	CustomerId     string           `json:"customerId" validate:"required"`
	ProductDetails []ProductDetail 	`json:"productDetails" validate:"required,min=1,dive"`
	Paid           float64          `json:"paid" validate:"required,min=1"`
	Change         *float64          `json:"change" validate:"min=0"`
}

type TransactionGetAllQueryParams struct {
	CustomerId string `form:"customerId"`
	Limit      string `form:"limit" default:"5"`
	Offset     string `form:"offset" default:"0"`
	CreatedAt   string `form:"createdAt"`
}