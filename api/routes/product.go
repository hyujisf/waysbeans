package routes

import (
	"waysbeans/handlers"
	"waysbeans/pkg/middleware"
	"waysbeans/pkg/sql"
	"waysbeans/repositories"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Group) {
	productRepository := repositories.MakeRepository(sql.DB)
	h := handlers.HandlerProduct(productRepository)

	// mengambil semua data product
	e.GET("/products", h.FindProducts)

	// mengambil satu data product
	e.GET("/product/:id", h.GetProduct)

	// menambahkan product
	e.POST("/product", middleware.AdminAuth(middleware.UploadFile(h.AddProduct)))
	// e.DELETE("/product/:id", middleware.Auth(h.DeleteProduct))
	// e.PATCH("/product/:id", middleware.Auth(middleware.UploadFile(h.UpdateProduct)))
}
