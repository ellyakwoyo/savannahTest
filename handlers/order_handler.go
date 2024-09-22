package handlers

import (
	"net/http"
	"savannahTest/services"
	"strconv"

	"github.com/labstack/echo/v4"
	"savannahTest/models"
)

type OrderHandler struct {
	orderService services.OrderService
}

func NewOrderHandler(orderService services.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

// CreateOrder creates a new order
// @Summary Create a new order
// @Description Creates a new order in the system with the provided data
// @Tags Orders
// @Accept json
// @Produce json
// @Param order body models.Order true "Order Data"
// @Success 200 {object} map[string]interface{} "Order created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /app/v1/orders [post]
func (h *OrderHandler) CreateOrder(c echo.Context) error {
	order := new(models.Order)
	if err := c.Bind(order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":  http.StatusBadRequest,
			"error": "Invalid request",
		})
	}

	if err := h.orderService.CreateOrder(order.ProductID, order.Quantity, order.Total, order.UserId); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Order created successfully",
		"data":    order,
	})
}

// GetAllOrders retrieves all orders
// @Summary Get all orders
// @Description Retrieves all orders from the system
// @Tags Orders
// @Produce json
// @Success 200 {object} map[string]interface{} "Orders fetched successfully"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /app/v1/orders [get]
func (h *OrderHandler) GetAllOrders(c echo.Context) error {
	orders, err := h.orderService.GetAllOrders()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Orders fetched successfully",
		"data":    orders,
	})
}

// GetOrderByID retrieves an order by its ID
// @Summary Get order by ID
// @Description Retrieves an order by its unique ID
// @Tags Orders
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} map[string]interface{} "Order fetched successfully"
// @Failure 400 {object} map[string]interface{} "Invalid order ID"
// @Failure 404 {object} map[string]interface{} "Order not found"
// @Router /app/v1/orders/{id} [get]
func (h *OrderHandler) GetOrderByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":  http.StatusBadRequest,
			"error": "Invalid order ID",
		})
	}

	order, err := h.orderService.GetOrderByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":  http.StatusNotFound,
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Order fetched successfully",
		"data":    order,
	})
}

// UpdateOrder updates an existing order
// @Summary Update an order
// @Description Updates an existing order by ID with new data
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param order body models.Order true "Updated order data"
// @Success 200 {object} map[string]interface{} "Order updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid ID or request"
// @Failure 404 {object} map[string]interface{} "Order not found"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /app/v1/orders/{id} [put]
func (h *OrderHandler) UpdateOrder(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":  http.StatusBadRequest,
			"error": "Invalid order ID",
		})
	}

	order, err := h.orderService.GetOrderByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":  http.StatusNotFound,
			"error": err.Error(),
		})
	}

	if err := c.Bind(order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":  http.StatusBadRequest,
			"error": "Invalid request",
		})
	}

	if err := h.orderService.UpdateOrder(order); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Order updated successfully",
		"data":    order,
	})
}

// DeleteOrder deletes an order by its ID
// @Summary Delete an order
// @Description Deletes an order from the system by its unique ID
// @Tags Orders
// @Param id path int true "Order ID"
// @Success 200 {object} map[string]interface{} "Order deleted successfully"
// @Failure 400 {object} map[string]interface{} "Invalid order ID"
// @Failure 404 {object} map[string]interface{} "Order not found"
// @Router /app/v1/orders/{id} [delete]
func (h *OrderHandler) DeleteOrder(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":  http.StatusBadRequest,
			"error": "Invalid order ID",
		})
	}

	if err := h.orderService.DeleteOrder(id); err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":  http.StatusNotFound,
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Order deleted successfully",
	})
}
