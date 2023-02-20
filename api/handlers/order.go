package handlers

import (
	"net/http"
	"strconv"
	"waysbeans/dto"
	"waysbeans/models"
	"waysbeans/repositories"

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

	userInfo := c.Get("userInfo").(jwt.MapClaims)
	idUser := int(userInfo["id"].(float64))
	orders, err := h.OrderRepository.FindOrders(idUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: convertMultipleOrderResponse(orders)})
}

// mengambil data 1 product
func (h *handlerOrder) GetOrder(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	order, err := h.OrderRepository.GetOrder(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: convertOrderResponse(order)})
}
func (h *handlerOrder) CreateOrder(c echo.Context) error {

	var request dto.AddOrderRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	userInfo := c.Get("userInfo").(jwt.MapClaims)
	idUser := int(userInfo["id"].(float64))

	// periksa order dengan product id yang sama
	order, err := h.OrderRepository.GetOrderByProduct(request.ProductID, idUser)
	if err != nil {
		// bila belum ada, maka buat baru
		newOrder := models.Order{
			UserID:    idUser,
			ProductID: request.ProductID,
			OrderQty:  1,
		}

		orderAdded, err := h.OrderRepository.CreateOrder(newOrder)
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
		}

		order, err := h.OrderRepository.GetOrder(orderAdded.ID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
		}

		return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: convertOrderResponse(order)})
	}

	// bila sudah ada, maka cukup tambahkan qty
	order.OrderQty = order.OrderQty + 1

	orderUpdated, err := h.OrderRepository.UpdateOrder(order)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	order, err = h.OrderRepository.GetOrder(orderUpdated.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: convertOrderResponse(order)})
}

func (h *handlerOrder) UpdateOrder(c echo.Context) error {
	var request dto.UpdateOrderRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	order, err := h.OrderRepository.GetOrder(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	if request.Event == "add" {
		order.OrderQty = order.OrderQty + 1
	} else if request.Event == "less" {
		order.OrderQty = order.OrderQty - 1
	}

	if request.Qty != 0 {
		order.OrderQty = request.Qty
	}

	orderUpdated, err := h.OrderRepository.UpdateOrder(order)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	order, err = h.OrderRepository.GetOrder(orderUpdated.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Status: "success",
		Data:   convertOrderResponse(order),
	})
}
func (h *handlerOrder) DeleteOrder(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: "Invalid ID"})
	}

	order, err := h.OrderRepository.GetOrder(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.ErrorResult{Status: "error", Message: "Order not found"})
	}

	orderDeleted, err := h.OrderRepository.DeleteOrder(order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: "Failed to delete order"})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: convertOrderResponse(orderDeleted)})
}

func convertMultipleOrderResponse(orders []models.Order) []dto.OrderResponse {
	var OrderResponse []dto.OrderResponse

	for _, order := range orders {
		OrderResponse = append(OrderResponse, dto.OrderResponse{
			ID:       order.ID,
			OrderQty: order.OrderQty,
			Product:  order.Product,
		})
	}

	return OrderResponse
}

func convertOrderResponse(order models.Order) dto.OrderResponse {
	return dto.OrderResponse{
		ID:       order.ID,
		OrderQty: order.OrderQty,
		Product:  order.Product,
	}
}
