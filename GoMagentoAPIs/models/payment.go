package models

type Payment struct {
	ID          int64   `json:"id"`
	OrderID     int64   `json:"order_id"`
	Amount      float64 `json:"amount"`
	PaymentType string  `json:"payment_type"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}