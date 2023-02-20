package handlers

import (
	"net/http"
	"strconv"
	"waysbeans/dto"
	"waysbeans/models"
	"waysbeans/repositories"

	"github.com/labstack/echo/v4"
)

type handler struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handler {
	return &handler{UserRepository}
}

func (h *handler) FindUsers(c echo.Context) error {
	users, err := h.UserRepository.FindUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: convertMultipleUserResponse(users)})
}

func (h *handler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: convertUserResponse(user)})
}

// func (h *handler) UpdateUser(c echo.Context) error {
// 	request := new(dto.UpdateUserRequest)
// 	if err := c.Bind(request); err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
// 	}

// 	id, _ := strconv.Atoi(c.Param("id"))

// 	user, err := h.UserRepository.GetUser(id)

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
// 	}

// 	if request.Name != "" {
// 		user.Name = request.Name
// 	}

// 	if request.Email != "" {
// 		user.Email = request.Email
// 	}

// 	if request.Password != "" {
// 		user.Password = request.Password
// 	}

// 	data, err := h.UserRepository.UpdateUser(user)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: convertResponse(data)})
// }

// func (h *handler) DeleteUser(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	user, err := h.UserRepository.GetUser(id)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
// 	}

// 	data, err := h.UserRepository.DeleteUser(user, id)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: convertResponse(data)})
// }

func convertMultipleUserResponse(users []models.User) []dto.UserResponse {
	var usersResponse []dto.UserResponse

	for _, user := range users {
		var userResponse = dto.UserResponse{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Role:     user.Role,
			Phone:    user.Phone,
			Image:    user.Image,
			Address:  user.Address,
			PostCode: user.PostCode,
		}

		usersResponse = append(usersResponse, userResponse)
	}

	return usersResponse
}

func convertUserResponse(user models.User) dto.UserResponse {
	var userResponse = dto.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Role:     user.Role,
		Phone:    user.Phone,
		Image:    user.Image,
		Address:  user.Address,
		PostCode: user.PostCode,
	}

	return userResponse
}
