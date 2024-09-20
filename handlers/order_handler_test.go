package handlers_test

import (
	"bytes"
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"savannahTest/handlers"
	"savannahTest/mocks"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"savannahTest/models"
)

func TestCreateOrder(t *testing.T) {
	e := echo.New()
	mockOrderService := new(mocks.MockOrderService)
	orderHandler := handlers.NewOrderHandler(mockOrderService)

	mockOrderService.On("CreateOrder", 1, 2, 100.0, 1).Return(nil)

	order := &models.Order{ProductID: 1, Quantity: 2, Total: 100.0, UserId: 1}
	orderJSON, _ := json.Marshal(order)
	req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(orderJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := orderHandler.CreateOrder(c)
	if err != nil {
		t.Fatalf("CreateOrder handler returned an error: %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Order created successfully")
	mockOrderService.AssertExpectations(t)
}

func TestGetAllOrders(t *testing.T) {
	e := echo.New()
	mockOrderService := new(mocks.MockOrderService)
	orderHandler := handlers.NewOrderHandler(mockOrderService)

	mockOrderService.On("GetAllOrders").Return([]models.Order{
		{Model: gorm.Model{ID: 1}, ProductID: 1, Quantity: 2, Total: 100.0, UserId: 1},
		{Model: gorm.Model{ID: 2}, ProductID: 2, Quantity: 1, Total: 50.0, UserId: 2},
	}, nil)

	req := httptest.NewRequest(http.MethodGet, "/orders", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := orderHandler.GetAllOrders(c)
	if err != nil {
		t.Fatalf("GetAllOrders handler returned an error: %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Orders fetched successfully")
	mockOrderService.AssertExpectations(t)
}
func TestGetOrderByID(t *testing.T) {
	e := echo.New()
	mockOrderService := new(mocks.MockOrderService)
	orderHandler := handlers.NewOrderHandler(mockOrderService)

	orderID := 1
	expectedOrder := &models.Order{
		Model: gorm.Model{ID: uint(orderID)}, ProductID: 1, Quantity: 2, Total: 100.0, UserId: 1,
	}
	mockOrderService.On("GetOrderByID", orderID).Return(expectedOrder, nil)

	req := httptest.NewRequest(http.MethodGet, "/orders/"+strconv.Itoa(orderID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(orderID))

	err := orderHandler.GetOrderByID(c)
	if err != nil {
		t.Fatalf("GetOrderByID handler returned an error: %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)

	type OrderResponse struct {
		Code    int          `json:"code"`
		Message string       `json:"message"`
		Data    models.Order `json:"data"`
	}

	var actualResponse OrderResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &actualResponse); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	expectedResponse := OrderResponse{
		Code:    http.StatusOK,
		Message: "Order fetched successfully",
		Data:    *expectedOrder,
	}

	assert.Equal(t, expectedResponse, actualResponse)

	// Verify mock expectations
	mockOrderService.AssertExpectations(t)
}

func TestUpdateOrder(t *testing.T) {
	e := echo.New()
	mockOrderService := new(mocks.MockOrderService)
	orderHandler := handlers.NewOrderHandler(mockOrderService)

	orderID := 1
	order := &models.Order{Model: gorm.Model{ID: uint(orderID)}, ProductID: 1, Quantity: 2, Total: 150.0, UserId: 1}
	mockOrderService.On("GetOrderByID", orderID).Return(order, nil)
	mockOrderService.On("UpdateOrder", order).Return(nil)

	orderJSON, _ := json.Marshal(order)
	req := httptest.NewRequest(http.MethodPut, "/orders/"+strconv.Itoa(orderID), bytes.NewBuffer(orderJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(orderID))

	// Call handler
	err := orderHandler.UpdateOrder(c)
	if err != nil {
		t.Fatalf("UpdateOrder handler returned an error: %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Order updated successfully")
	mockOrderService.AssertExpectations(t)
}

func TestDeleteOrder(t *testing.T) {
	e := echo.New()
	mockOrderService := new(mocks.MockOrderService)
	orderHandler := handlers.NewOrderHandler(mockOrderService)

	orderID := 1
	mockOrderService.On("DeleteOrder", orderID).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/orders/"+strconv.Itoa(orderID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(orderID))

	err := orderHandler.DeleteOrder(c)
	if err != nil {
		t.Fatalf("DeleteOrder handler returned an error: %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Order deleted successfully")
	mockOrderService.AssertExpectations(t)
}
