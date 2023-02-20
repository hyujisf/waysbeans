package routes

import (
	"waysbeans/handlers"
	"waysbeans/pkg/middleware"
	"waysbeans/pkg/sql"
	"waysbeans/repositories"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Group) {
	authRepository := repositories.MakeRepository(sql.DB)
	h := handlers.HandlerAuth(authRepository)

	// menghandle request dengan method POST pada endpoint /register
	e.POST("/register", h.Register)

	// menghandle request dengan method POST pada endpoint /login
	e.POST("/login", h.Login)

	// endpoint untuk pengecekan status login              // add this code
	e.GET("/check-auth", middleware.Auth(h.CheckAuth)) // add this code
}
