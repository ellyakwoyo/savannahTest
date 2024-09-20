// mocks/customer_service.go
package mocks

import (
	"github.com/stretchr/testify/mock"
	"savannahTest/models"
)

type MockCustomerService struct {
	mock.Mock
}

func (m *MockCustomerService) CreateCustomer(name, email string) error {
	args := m.Called(name, email)
	return args.Error(0)
}

func (m *MockCustomerService) GetAllCustomers() ([]models.Customer, error) {
	args := m.Called()
	return args.Get(0).([]models.Customer), args.Error(1)
}

func (m *MockCustomerService) GetCustomerByID(id int) (*models.Customer, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Customer), args.Error(1)
}

func (m *MockCustomerService) UpdateCustomer(id int, customerData *models.Customer) error {
	args := m.Called(id, customerData)
	return args.Error(0)
}

func (m *MockCustomerService) DeleteCustomer(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
