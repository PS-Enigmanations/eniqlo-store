package request

import (
	"enigmanations/eniqlo-store/internal/common/request"
)

type SearchProductQueryParams struct {
	Name     string `form:"name"`
	Category string `form:"category" validate:"oneof=Clothing Accessories Footwear Beverages"`
	Sku      string `form:"sku"`
	Price    string `form:"price" validate:"oneof=asc desc"`
	inStock  string `form:"inStock" validate:"oneof=true false"`

	*request.PaginationParams
}
