package request

type SearchProductQueryParams struct {
	Id          string `form:"id"`
	Name        string `form:"name"`
	Category    string `form:"category" validate:"oneof=Clothing Accessories Footwear Beverages"`
	Sku         string `form:"sku"`
	Price       string `form:"price" validate:"oneof=asc desc"`
	InStock     string `form:"inStock" validate:"oneof=true false"`
	CreatedAt   string `form:"createdAt"`
	IsAvailable string `form:"isAvailable"`

	// Pagination
	Limit  string `form:"limit" default:"5"`
	Offset string `form:"offset" default:"0"`
}
