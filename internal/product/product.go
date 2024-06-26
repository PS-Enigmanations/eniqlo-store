package product

import (
	"slices"
	"time"
)

type Product struct {
	Id          string   `json:"id"`
	Name        string   `json:"name" validate:"required,min=1,max=30"`
	Sku         string   `json:"sku" validate:"required,min=1,max=30"`
	Category    Category `json:"category" validate:"required,oneof=Clothing Accessories Footwear Beverages"`
	ImageUrl    string   `json:"imageUrl" validate:"required,dive,url"`
	Notes       string   `json:"notes" validate:"required,min=1,max=200"`
	Price       float64  `json:"price" validate:"required,min=1"`
	Stock       int      `json:"stock" validate:"required,min=0,max=100000"`
	Location    string   `json:"location" validate:"required,min=1,max=200"`
	IsAvailable bool     `json:"isAvailable"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Category string

// @TODO: Need separate table to store categories `product_types`, `product_type_products`
const (
	Clothing    Category = "Clothing"
	Accessories Category = "Accessories"
	Footwear    Category = "Footwear"
	Beverages   Category = "Beverages"
)

var ImageFormats = []string{".jpg", ".jpeg", ".png", ".webp"}

type CategoryItem struct {
	Name   string
	Search []string
}

var Categories = []CategoryItem{
	{Name: "Clothing", Search: []string{"clothing", "Clothing"}},
	{Name: "Accessories", Search: []string{"accessories", "Accessories"}},
	{Name: "Footwear", Search: []string{"footwear", "Footwear"}},
	{Name: "Beverages", Search: []string{"beverages", "Beverages"}},
}

func HasCategory(cat string) bool {
	found := false
	for _, item := range Categories {
		if slices.Contains(item.Search, cat) {
			found = true
			break
		}
	}

	return found
}
