package dto

type UpdateUserRequest struct {
	Name     string `form:"name"`
	Phone    string `form:"phone"`
	Address  string `form:"address"`
	PostCode string `form:"post_code"`
	Image    string `form:"image"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	PostCode string `json:"post_code"`
	Image    string `json:"image"`
	Role     string `json:"role"`
}
