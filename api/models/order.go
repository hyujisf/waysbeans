package models

type Order struct {
	ID            int
	UserID        int `gorm:"type: int"`
	User          UserResponse
	TransactionID string `gorm:"type: varchar(255)"`
	Transaction   TransactionResponse
	ProductID     int `gorm:"type: int"`
	Product       ProductResponse
	OrderQty      int `gorm:"type: int"`
}

type OrderResponseForTransaction struct {
	ID            int    `json:"-"`
	TransactionID string `json:"-" gorm:"type: varchar(255)"`
	ProductID     int    `json:"-"`
	Product       ProductResponse
	OrderQty      int `json:"orderQty" gorm:"type: int"`
}

func (OrderResponseForTransaction) TableName() string {
	return "orders"
}
