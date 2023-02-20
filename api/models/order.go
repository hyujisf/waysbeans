package models

import "time"

type Order struct {
	ID        int             `json:"id" gorm:"primary_key:auto_increment"`
	QTY       int             `json:"qty"`
	SubTotal  int             `json:"subtotal"`
	ProductID int             `json:"product_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product   ProductResponse `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    int             `json:"user_id"`
	User      UserResponse    `json:"user"`
	Status    string          `json:"status"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
}

type OrderResponse struct {
	ID        int             `json:"id"`
	Qty       int             `json:"qty"`
	SubTotal  int             `json:"subtotal"`
	ProductID int             `json:"product_id"`
	Product   ProductResponse `json:"product"`
	Status    string          `json:"status"`
}

func (OrderResponse) TableName() string {
	return "orders"
}
