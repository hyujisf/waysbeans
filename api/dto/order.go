package dto

type CreateOrder struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	ProductID int    `json:"product_id"`
	QTY       int    `json:"qty"`
	SubTotal  int    `json:"subtotal"`
	Status    string `jsom:"status"`
}

type UpdateOrder struct {
	ID       int    `json:"id"`
	QTY      int    `json:"qty"`
	SubTotal int    `json:"subtotal"`
	Status   string `jsom:"status"`
}

type OrderResponse struct {
	ID       int `json:"id"`
	QTY      int `json:"qty"`
	SubTotal int `json:"subtotal"`
}
