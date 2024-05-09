package request

type ProductDetail struct {
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity" validate:"required,min=1"`
}

type CheckoutRequest struct {
	CustomerId     string           `json:"customerId"`
	ProductDetails []ProductDetail `json:"productDetails" validate:"required,min=1,dive"`
	Paid           int              `json:"paid" validate:"required,min=1"`
	Change         int              `json:"change" validate:"required,min=0"`
}