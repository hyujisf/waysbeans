package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"waysbeans/dto"
	"waysbeans/models"
	"waysbeans/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerProduct struct {
	ProductRepository repositories.ProductRepository
}

func HandlerProduct(ProductRepository repositories.ProductRepository) *handlerProduct {
	return &handlerProduct{ProductRepository}
}

// mengambil data semua product
func (h *handlerProduct) FindProducts(c echo.Context) error {
	products, err := h.ProductRepository.FindProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: convertMultipleProductResponse(products)})
}

// mengambil data 1 product
func (h *handlerProduct) GetProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var product models.Product
	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: convertProductResponse(product)})
}

// menambahkan product baru
func (h *handlerProduct) AddProduct(c echo.Context) error {
	var err error

	dataFile := c.Get("dataFile").(string)
	fmt.Println("ini data", dataFile)

	request := dto.CreateProductRequest{
		Name:        c.FormValue("name"),
		Description: c.FormValue("description"),
		Image:       dataFile,
	}
	request.Stock, _ = strconv.Atoi(c.FormValue("stock"))
	request.Price, _ = strconv.ParseFloat(c.FormValue("price"), 64)

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	newProduct := models.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Image:       request.Image,
		Stock:       request.Stock,
		Status:      "active",
	}

	productAdded, err := h.ProductRepository.AddProduct(newProduct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	product, _ := h.ProductRepository.GetProduct(productAdded.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: convertProductResponse(product)})
}

func (h *handlerProduct) UpdateProduct(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)

	request := dto.UpdateProductRequest{
		Name:        c.FormValue("name"),
		Image:       dataFile,
		Description: c.FormValue("description"),
	}
	request.Stock, _ = strconv.Atoi(c.FormValue("stock"))
	request.Price, _ = strconv.ParseFloat(c.FormValue("price"), 64)

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.ProductRepository.GetProduct(int(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	if request.Name != "" {
		product.Name = request.Name
	}

	if request.Price != 0 {
		product.Price = request.Price
	}

	if dataFile != "false" {
		product.Image = request.Image
	}

	if request.Stock != 0 {
		product.Stock = request.Stock
	}

	if request.Description != "" {
		product.Description = request.Description
	}

	product, err = h.ProductRepository.UpdateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: convertProductResponse(product)})
}
func (h *handlerProduct) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	updateProduct, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.ErrorResult{
			Status:  "error",
			Message: err.Error(),
		})
	}

	updateProduct.Status = "inactive"

	productUpdated, err := h.ProductRepository.UpdateProduct(updateProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Status:  "error",
			Message: err.Error(),
		})
	}

	response := dto.SuccessResult{
		Status: "success",
		Data:   fmt.Sprintf("Product dengan id %d berhasil di non-aktifkan", productUpdated.ID),
	}
	return c.JSON(http.StatusOK, response)
}
func convertMultipleProductResponse(products []models.Product) []dto.ProductResponse {
	var productsResponse []dto.ProductResponse

	for _, p := range products {
		productsResponse = append(productsResponse, dto.ProductResponse(p))
	}

	return productsResponse
}

func convertProductResponse(product models.Product) dto.ProductResponse {
	return dto.ProductResponse(product)
}
