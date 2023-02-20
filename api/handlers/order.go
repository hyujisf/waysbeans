package handlers

import (
	"net/http"
	"strconv"
	"waysbeans/dto"
	"waysbeans/models"
	"waysbeans/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerOrder struct {
	OrderRepository repositories.OrderRepository
}

func HandlerOrder(orderRepository repositories.OrderRepository) *handlerOrder {
	return &handlerOrder{orderRepository}
}

// mengambil data semua product

func (h *handlerOrder) FindOrders(c echo.Context) error {

	orders, err := h.OrderRepository.FindOrders()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: orders})
}

// mengambil data 1 product
func (h *handlerOrder) GetOrder(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	order, err := h.OrderRepository.GetOrder(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: order})
}
func (h *handlerOrder) CreateOrder(c echo.Context) error {
	userLogin := c.Get("userLogin").(jwt.MapClaims)
	id := int(userLogin["id"].(float64))

	request := new(dto.CreateOrder)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	order := models.Order{
		ProductID: request.ProductID,
		UserID:    id,
		QTY:       request.QTY,
		SubTotal:  request.SubTotal,
		Status:    "on",
	}

	data, err := h.OrderRepository.CreateOrder(order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: data})
}
func (h *handlerOrder) UpdatesOrder(c echo.Context) error {
	request := new(dto.UpdateOrder)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	userLogin := c.Get("userLogin").(jwt.MapClaims)
	id := int(userLogin["id"].(float64))

	order, err := h.OrderRepository.FindOrdersTransaction(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	for i := range order {
		order[i].Status = "off"
	}

	data, err := h.OrderRepository.UpdatesOrder(order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: data})
}

func (h *handlerOrder) UpdateOrder(c echo.Context) error {
	request := new(dto.UpdateOrder)
	if err := c.Bind(request); err != nil {
		response := dto.ErrorResult{Status: "error", Message: err.Error()}
		return c.JSON(http.StatusBadRequest, response)
	}

	id, _ := strconv.Atoi(c.Param("id"))

	order, err := h.OrderRepository.GetOrder(id)
	if err != nil {
		response := dto.ErrorResult{Status: "error", Message: err.Error()}
		return c.JSON(http.StatusInternalServerError, response)
	}

	if request.QTY != 0 {
		order.QTY = request.QTY
	}

	data, err := h.OrderRepository.UpdateOrder(order)
	if err != nil {
		response := dto.ErrorResult{Status: "error", Message: err.Error()}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := dto.SuccessResult{Status: "success", Data: data}
	return c.JSON(http.StatusOK, response)
}
func (h *handlerOrder) DeleteOrder(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := h.OrderRepository.GetOrder(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}
	data, err := h.OrderRepository.DeleteOrder(order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: data})
}
func (h *handlerOrder) FindOrdersByID(c echo.Context) error {
	userLogin := c.Get("userLogin").(jwt.MapClaims)
	id := int(userLogin["id"].(float64))
	order, err := h.OrderRepository.FindOrdersTransaction(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	for i, p := range order {
		order[i].Product.Image = p.Product.Image
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: order})
}

func convertResponseOrder(u models.Order) models.Order {
	return models.Order{
		ID:       u.ID,
		QTY:      u.QTY,
		SubTotal: u.SubTotal,
		Product:  u.Product,
	}
}
