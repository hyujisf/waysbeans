package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	Password string `json:"password" gorm:"type: varchar(255)"`
	Image    string `json:"image" gorm:"type: varchar(255)"`
	Address  string `json:"address" gorm:"type: varchar(255)"`
	PostCode string `json:"post_code" gorm:"type: varchar(255)"`
	Role     string `json:"role" gorm:"type: varchar(255)"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	PostCode string `json:"post_code"`
	Role     string `json:"role"`
}

func (UserResponse) TableName() string {
	return "users"
}
