package routes

import (
	"waysbeans/handlers"
	"waysbeans/pkg/middleware"
	"waysbeans/pkg/sql"
	"waysbeans/repositories"

	"github.com/labstack/echo/v4"
)

func Transaction(e *echo.Group) {
	transactionRepository := repositories.MakeRepository(sql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	e.GET("/transactions", middleware.Auth(h.FindTransactions))
	e.GET("/transaction/{id}", middleware.Auth(h.GetTransaction))
	e.GET("/user-transaction", middleware.Auth(h.GetUserTransaction))
	e.GET("/transaction", middleware.Auth(h.CreateTransaction))
	e.PATCH("/transaction/", middleware.Auth(h.UpdateTransaction))
	e.DELETE("/transaction/{id}", middleware.Auth(h.DeleteTransaction))

	// membuat endpoint untuk midtrans
	e.POST("/notification", h.Notification)

}
