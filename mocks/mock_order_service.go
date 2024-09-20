// mocks/order_service.go
package mocks

import (
	"github.com/stretchr/testify/mock"
	"savannahTest/models"
)

type MockOrderService struct {
	mock.Mock
}

func (m *MockOrderService) CreateOrder(productID, quantity int, total float64, userID int) error {
	args := m.Called(productID, quantity, total, userID)
	return args.Error(0)
}

func (m *MockOrderService) GetAllOrders() ([]models.Order, error) {
	args := m.Called()
	return args.Get(0).([]models.Order), args.Error(1)
}

func (m *MockOrderService) GetOrderByID(id int) (*models.Order, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Order), args.Error(1)
}

func (m *MockOrderService) UpdateOrder(order *models.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrderService) DeleteOrder(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
