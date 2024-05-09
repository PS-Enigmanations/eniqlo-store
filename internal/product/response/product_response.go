package response

import (
	"enigmanations/eniqlo-store/internal/product"
	"time"
)

// Search response
type ProductShow struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Sku       string    `json:"sku"`
	Category  string    `json:"category"`
	ImageUrl  string    `json:"imageUrl"`
	Stock     int       `json:"stock"`
	Price     float64   `json:"price"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"created_at"`
}

type ProductShows []ProductShow

type SearchProductsResponse struct {
	Message string       `json:"message"`
	Data    ProductShows `json:"data"`
}

func ToProductShows(p []*product.Product) ProductShows {
	list := make(ProductShows, len(p))
	for i, item := range p {
		list[i] = ProductShow{
			Id:        item.Id,
			Name:      item.Name,
			Sku:       item.Sku,
			Category:  string(item.Category),
			ImageUrl:  item.ImageUrl,
			Stock:     item.Stock,
			Price:     item.Price,
			Location:  item.Location,
			CreatedAt: item.CreatedAt,
		}
	}

	return list
}

const SearchProductsSuccMessage = "Successfully get products"

func ProductToSearchProductsResponse(data ProductShows) *SearchProductsResponse {
	return &SearchProductsResponse{
		Message: SearchProductsSuccMessage,
		Data:    data,
	}
}
