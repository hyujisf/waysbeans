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
	request.Price, _ = strconv.Atoi(c.FormValue("price"))

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

// func (h *handlerProduct) UpdateProduct(c echo.Context) error {
// 	var err error

// 	id, _ := strconv.Atoi(c.Param("id"))

// 	var categoriesId []int
// 	categoryIdString := c.FormValue("category_id")
// 	err = json.Unmarshal([]byte(categoryIdString), &categoriesId)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
// 	}

// 	request := dto.UpdateProductRequest{
// 		Name:        c.FormValue("name"),
// 		Description: c.FormValue("description"),
// 		Image:       c.Get("image").(string),
// 	}
// 	request.Stock, _ = strconv.Atoi(c.FormValue("stock"))
// 	request.Price, _ = strconv.Atoi(c.FormValue("price"))

// 	updateProduct, err := h.ProductRepository.GetProduct(id)
// 	if err != nil {
// 		return c.JSON(http.StatusNotFound, dto.ErrorResult{
// 			Status:  "error",
// 			Message: err.Error(),
// 		})
// 	}

// 	validation := validator.New()
// 	err = validation.Struct(request)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
// 	}

// 	id, _ := strconv.Atoi(c.Param("id"))

// 	product, err := h.ProductRepository.GetProduct(id)

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
// 	}

// 	if request.Name != "" {
// 		product.Name = request.Name
// 	}

// 	if request.Desc != "" {
// 		product.Desc = request.Desc
// 	}

// 	if request.Price != 0 {
// 		product.Price = request.Price
// 	}

// 	if request.Image != "" {
// 		product.Image = request.Image
// 	}

// 	if request.Qty != 0 {
// 		product.Qty = request.Qty
// 	}

// 	if len(request.CategoryID) == 0 {
// 		data, err := h.ProductRepository.DeleteProductCategoryByProductId(product)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
// 		}

// 		return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: data})
// 	}

// 	categories, _ := h.ProductRepository.FindCategoriesById(request.CategoryID)
// 	product.Category = categories

// 	data, err := h.ProductRepository.UpdateProduct(product)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: data})
// }

// func (h *handlerProduct) DeleteProduct(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	product, err := h.ProductRepository.GetProduct(id)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
// 	}

// 	data, err := h.ProductRepository.DeleteProduct(product, id)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: data})
// }

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
