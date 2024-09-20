package services_test

import (
	"gorm.io/gorm"
	"savannahTest/mocks"
	"savannahTest/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateOrderService(t *testing.T) {

	mockOrderService := new(mocks.MockOrderService)

	productID, quantity, userId := 1, 2, 1
	total := 200.0

	mockOrderService.On("CreateOrder", productID, quantity, total, userId).Return(nil)

	err := mockOrderService.CreateOrder(productID, quantity, total, userId)

	assert.NoError(t, err)

	mockOrderService.AssertExpectations(t)
}

func TestGetAllOrdersService(t *testing.T) {

	mockOrderService := new(mocks.MockOrderService)

	expectedOrders := []models.Order{
		{Model: gorm.Model{ID: 1}, ProductID: 1, Quantity: 2, Total: 200.0, UserId: 1},
		{Model: gorm.Model{ID: 2}, ProductID: 2, Quantity: 3, Total: 300.0, UserId: 2},
	}

	mockOrderService.On("GetAllOrders").Return(expectedOrders, nil)

	orders, err := mockOrderService.GetAllOrders()

	assert.NoError(t, err)
	assert.Equal(t, expectedOrders, orders)

	mockOrderService.AssertExpectations(t)
}

func TestGetOrderByIDService(t *testing.T) {

	mockOrderService := new(mocks.MockOrderService)

	expectedOrder := &models.Order{Model: gorm.Model{ID: 1}, ProductID: 1, Quantity: 2, Total: 200.0, UserId: 1}

	mockOrderService.On("GetOrderByID", 1).Return(expectedOrder, nil)

	order, err := mockOrderService.GetOrderByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedOrder, order)

	mockOrderService.AssertExpectations(t)
}

func TestUpdateOrderService(t *testing.T) {

	mockOrderService := new(mocks.MockOrderService)

	order := &models.Order{Model: gorm.Model{ID: 1}, ProductID: 1, Quantity: 3, Total: 300.0, UserId: 1}

	mockOrderService.On("UpdateOrder", order).Return(nil)

	err := mockOrderService.UpdateOrder(order)

	assert.NoError(t, err)

	mockOrderService.AssertExpectations(t)
}

func TestDeleteOrderService(t *testing.T) {

	mockOrderService := new(mocks.MockOrderService)

	mockOrderService.On("DeleteOrder", 1).Return(nil)

	err := mockOrderService.DeleteOrder(1)

	assert.NoError(t, err)

	mockOrderService.AssertExpectations(t)
}
