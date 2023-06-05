package models

import "time"

type Order struct {
	ID                int64     `json:"id"`
	CustomerID        int64     `json:"customer_id"`
	OrderStatus       string    `json:"order_status"`
	TotalPrice        float64   `json:"total_price"`
	Items             []OrderItem `json:"items"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type OrderItem struct {
	ID          int64   `json:"id"`
	OrderID     int64   `json:"order_id"`
	ProductID   int64   `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

type OrderFilter struct {
	ID          *int64     `json:"id,omitempty"`
	CustomerID  *int64     `json:"customer_id,omitempty"`
	OrderStatus *string    `json:"order_status,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type OrderPagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Total    int `json:"total"`
}