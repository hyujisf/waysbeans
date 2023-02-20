package dto

import "waysbeans/models"

type ProductRequestForTransaction struct {
	ID        int `json:"id"`
	ProductID int `json:"product_id"`
	OrderQty  int `json:"orderQty"`
}

type ProductResponseForTransaction struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	Description string `json:"description"`
	OrderQty    int    `json:"orderQty"`
}

type CreateTransactionRequest struct {
	Total    int                            `json:"total" validate:"required"`
	UserID   int                            `json:"user_id" validate:"required"`
	Products []ProductRequestForTransaction `json:"products" validate:"required"`
}
type UpdateTransactionRequest struct {
	Total    int                            `json:"total"`
	UserID   int                            `json:"user_id"`
	Products []ProductRequestForTransaction `json:"products"`
	Status   string                         `json:"status"`
}

type TransactionResponse struct {
	ID         string                          `json:"id"`
	MidtransID string                          `json:"midtrans_id"`
	OrderDate  string                          `json:"order_date"`
	Total      int                             `json:"total"`
	Status     string                          `json:"status"`
	User       models.UserResponse             `json:"user"`
	Products   []ProductResponseForTransaction `json:"products"`
}
