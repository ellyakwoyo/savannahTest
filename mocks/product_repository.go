package mocks

import (
	"github.com/stretchr/testify/mock"
	"savannahTest/models"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) Create(product *models.Product) (*models.Product, error) {
	args := m.Called(product)
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductRepository) FindByID(id uint) (*models.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductRepository) Update(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockProductRepository) FindAll() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}
