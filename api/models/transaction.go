package models

import "time"

type Transaction struct {
	ID        int64     `json:"id"`
	UserID    int       `json:"user_id"`
	User      User      `json:"user"`
	Status    string    `json:"status"`
	Total     int       `json:"total"`
	OrderID   []int     `json:"order_id" gorm:"-"`
	Order     []Order   `json:"product" gorm:"many2many:transaction_order;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}

type TransactionResponse struct {
	ID     int64 `json:"id"`
	UserID int   `json:"user_id"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
