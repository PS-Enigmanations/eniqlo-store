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

type ProductRequest struct {
	Id          string  `form:"id"`
	Name        string  `form:"name" binding:"required"`
	Sku         string  `form:"sku" binding:"required"`
	Category    string  `form:"category" binding:"oneof=Clothing Accessories Footwear Beverages"`
	ImageUrl    string  `form:"imageUrl" binding:"required,url"`
	Notes       string  `form:"notes" binding:"required,min=1,max=200"`
	Price       float64 `form:"price" binding:"required"`
	Stock       int     `form:"stock" binding:"required,min=0,max=100000"`
	Location    string  `form:"location" binding:"required"`
	IsAvailable bool    `form:"isAvailable"`
}
