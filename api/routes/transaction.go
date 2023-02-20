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

	e.GET("/transactions-admin", middleware.AdminAuth(h.FindTransactions))
	e.GET("/transactions", middleware.Auth(h.FindTransactionsByUser))
	e.GET("/transaction/{id}", middleware.Auth(h.GetDetailTransaction))
	e.POST("/transaction", middleware.Auth(h.CreateTransaction))
	e.PATCH("/transaction/{id}", middleware.Auth(h.UpdateTransactionStatus))

	// membuat endpoint untuk midtrans
	e.POST("/notification", h.Notification)

}
