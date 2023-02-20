package models

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name" gorm:"type: varchar(255)"`
	Price       int    `json:"price" gorm:"type: int"`
	Stock       int    `json:"stock" gorm:"type: int"`
	Image       string `json:"image" gorm:"type: varchar(255)"`
	Description string `json:"description"`
	Status      string `json:"status" gorm:"type: varchar(255)"`
}

type ProductResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func (ProductResponse) TableName() string {
	return "products"
}
