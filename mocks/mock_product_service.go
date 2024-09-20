package mocks

import (
	"github.com/stretchr/testify/mock"
	"savannahTest/models"
)

type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) CreateProduct(name, description string, price float64) (*models.Product, error) {
	args := m.Called(name, description, price)
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductService) GetProductByID(id uint) (*models.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductService) UpdateProduct(id uint, productData *models.Product) error {
	args := m.Called(id, productData)
	return args.Error(0)
}

func (m *MockProductService) DeleteProduct(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockProductService) GetAllProducts() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}
