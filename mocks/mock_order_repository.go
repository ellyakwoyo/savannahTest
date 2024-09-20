package mocks

import (
	"github.com/stretchr/testify/mock"
	"savannahTest/models"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) CreateOrder(order *models.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrderRepository) GetAllOrders() ([]models.Order, error) {
	args := m.Called()
	return args.Get(0).([]models.Order), args.Error(1)
}

func (m *MockOrderRepository) GetOrderByID(id int) (*models.Order, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Order), args.Error(1)
}

func (m *MockOrderRepository) UpdateOrder(order *models.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrderRepository) DeleteOrder(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
