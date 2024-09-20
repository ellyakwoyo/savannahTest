package mocks

import (
	"github.com/stretchr/testify/mock"
	"savannahTest/models"
)

type MockCustomerRepository struct {
	mock.Mock
}

func (m *MockCustomerRepository) CreateCustomer(customer *models.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *MockCustomerRepository) GetAllCustomers() ([]models.Customer, error) {
	args := m.Called()
	return args.Get(0).([]models.Customer), args.Error(1)
}

func (m *MockCustomerRepository) GetCustomerByID(id int) (*models.Customer, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Customer), args.Error(1)
}

func (m *MockCustomerRepository) UpdateCustomer(id int, customer *models.Customer) error {
	args := m.Called(id, customer)
	return args.Error(0)
}

func (m *MockCustomerRepository) DeleteCustomer(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
