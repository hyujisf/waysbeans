package routes

import (
	"waysbeans/handlers"
	"waysbeans/pkg/middleware"
	"waysbeans/pkg/sql"
	"waysbeans/repositories"

	"github.com/labstack/echo/v4"
)

func Order(e *echo.Group) {
	orderRepository := repositories.MakeRepository(sql.DB)
	h := handlers.HandlerOrder(orderRepository)
	// find orders
	e.GET("/orders", middleware.Auth(h.FindOrders))

	// add order
	e.POST("/order", middleware.Auth(h.CreateOrder))

	// get 1 order
	e.GET("/order/{id}", middleware.Auth(h.GetOrder))

	// update order
	e.PATCH("/order/{id}", middleware.Auth(h.UpdateOrder))

	// delete order
	e.DELETE("/order/{id}", middleware.Auth(h.DeleteOrder))
}
