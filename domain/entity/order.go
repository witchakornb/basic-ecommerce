package entity

type Order struct {
	ID        string  `json:"id"`
	CustomerID string  `json:"customer_id"`
	ProductID  string  `json:"product_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	DeletedAt  string  `json:"deleted_at"`
}