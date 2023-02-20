package models

import "time"

type Transaction struct {
	ID         string                        `json:"id" gorm:"type: varchar(255);PRIMARY_KEY"`
	MidtransID string                        `json:"midtrans_id" gorm:"type: varchar(255)"`
	OrderDate  time.Time                     `json:"order_date"`
	Total      int                           `json:"total" gorm:"type: int"`
	Status     string                        `json:"status" gorm:"type: varchar(255)"`
	UserID     int                           `json:"user_id" gorm:"type: int"`
	User       UserResponse                  `json:"users"`
	Order      []OrderResponseForTransaction `json:"products" gorm:"foreignKey:TransactionID"`
}

type TransactionResponse struct {
	ID         string       `json:"id" gorm:"type: varchar(255);PRIMARY_KEY"`
	MidtransID string       `json:"midtrans_id" gorm:"type: varchar(255)"`
	OrderDate  time.Time    `json:"order_date"`
	Total      int          `json:"total" gorm:"type: int"`
	Status     string       `json:"status" gorm:"type: varchar(255)"`
	UserID     int          `json:"user_id" gorm:"type: int"`
	User       UserResponse `json:"users"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
