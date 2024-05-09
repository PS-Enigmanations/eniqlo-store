package request

type SearchProductQueryParams struct {
	Name     string `form:"name"`
	Category string `form:"category" validate:"oneof=Clothing Accessories Footwear Beverages"`
	Sku      string `form:"sku"`
	Price    string `form:"price" validate:"oneof=asc desc"`
	InStock  string `form:"inStock" validate:"oneof=true false"`

	// Pagination
	Limit  string `form:"limit" default:"5"`
	Offset string `form:"offset" default:"0"`
}
