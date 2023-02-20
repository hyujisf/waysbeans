package routes

import (
	"waysbeans/handlers"
	"waysbeans/pkg/sql"
	"waysbeans/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.MakeRepository(sql.DB)
	h := handlers.HandlerUser(userRepository)

	e.GET("/users", h.FindUsers)
	e.GET("/user/:id", h.GetUser)
	// e.POST("/user", h.CreateUser)
	// e.PATCH("/user/:id", h.UpdateUser)
	// e.DELETE("/user/:id", h.DeleteUser)
}
