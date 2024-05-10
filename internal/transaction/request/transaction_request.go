package request

type ProductDetail struct {
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity" validate:"required,min=1"`
}

type CheckoutRequest struct {
	CustomerId     string           `json:"customerId"`
	ProductDetails []ProductDetail `json:"productDetails" validate:"required,min=1,dive"`
	Paid           int              `json:"paid" validate:"required,min=1"`
	Change         int              `json:"change" validate:"min=0"`
}

type TransactionGetAllQueryParams struct {
	CustomerId string `form:"customerId"`
	Limit      string `form:"limit" default:"5"`
	Offset     string `form:"offset" default:"0"`
	CreatedAt   string `form:"createdAt"`
}