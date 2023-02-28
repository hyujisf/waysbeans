package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"waysbeans/dto"

	"waysbeans/models"
	"waysbeans/pkg/bcrypt"
	jwtToken "waysbeans/pkg/jwt"
	"waysbeans/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(c echo.Context) error {
	request := new(dto.RegisterRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	hashedPassword, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	userData := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashedPassword,
		Role:     "customer",
		Image:    "https://api.dicebear.com/5.x/thumbs/svg?seed=" + strconv.Itoa(int(time.Now().Unix())),
	}

	registering, err := h.AuthRepository.Register(userData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: registering})
}

func (h *handlerAuth) Login(c echo.Context) error {
	request := new(dto.LoginRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	// Validasi email
	userLogin, err := h.AuthRepository.Login(user.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: "Email not registered"})
	}

	// Check password
	isValid := bcrypt.CheckPasswordHash(request.Password, userLogin.Password)
	if !isValid {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: "wrong email or password"})
	}

	// menyiapkan claims
	claims := jwt.MapClaims{}
	claims["id"] = userLogin.ID
	claims["name"] = userLogin.Name
	claims["email"] = userLogin.Email
	claims["role"] = userLogin.Role
	claims["image"] = userLogin.Image
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix() // 2 hours expired

	// menggenerate token jwt
	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	// membuat data yang akan disisipkan di response
	response := dto.LoginResponse{
		ID:    userLogin.ID,
		Name:  userLogin.Name,
		Email: userLogin.Email,
		Role:  userLogin.Role,
		Token: token,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: response})
}

func (h *handlerAuth) CheckAuth(c echo.Context) error {
	claims, ok := c.Get("userLogin").(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid user information")
	}

	authResponse := dto.LoginResponse{
		ID:    int(claims["id"].(float64)),
		Name:  claims["name"].(string),
		Email: claims["email"].(string),
		Role:  claims["role"].(string),
		Token: strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", 1),
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: authResponse})
}
