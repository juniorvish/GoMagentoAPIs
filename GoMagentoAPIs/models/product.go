package models

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	SKU         string  `json:"sku"`
	ImageURL    string  `json:"image_url"`
	CategoryID  int     `json:"category_id"`
	Stock       int     `json:"stock"`
}

type ProductFilter struct {
	Name        *string `json:"name,omitempty"`
	PriceMin    *float64 `json:"price_min,omitempty"`
	PriceMax    *float64 `json:"price_max,omitempty"`
	CategoryID  *int     `json:"category_id,omitempty"`
}

type ProductPagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}