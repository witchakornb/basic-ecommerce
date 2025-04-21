package entity

type Order struct {
	ID        int  `json:"id"`
	CustomerID int  `json:"customer_id"`
	ProductID  int  `json:"product_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	DeletedAt  string  `json:"deleted_at"`
}