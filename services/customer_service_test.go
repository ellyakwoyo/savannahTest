package services_test

import (
	"gorm.io/gorm"
	"savannahTest/mocks"
	"savannahTest/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCustomer(t *testing.T) {
	mockCustomerService := new(mocks.MockCustomerService)

	name := "Test Customer"
	code := "CUST123"

	mockCustomerService.On("CreateCustomer", name, code).Return(nil)

	err := mockCustomerService.CreateCustomer(name, code)

	assert.NoError(t, err)

	mockCustomerService.AssertExpectations(t)
}

func TestGetAllCustomers(t *testing.T) {
	mockCustomerService := new(mocks.MockCustomerService)

	expectedCustomers := []models.Customer{
		{Model: gorm.Model{ID: 1}, Name: "Customer 1", Code: "CUST001"},
		{Model: gorm.Model{ID: 2}, Name: "Customer 2", Code: "CUST002"},
	}

	mockCustomerService.On("GetAllCustomers").Return(expectedCustomers, nil)

	customers, err := mockCustomerService.GetAllCustomers()

	assert.NoError(t, err)
	assert.Equal(t, expectedCustomers, customers)

	mockCustomerService.AssertExpectations(t)
}

func TestGetCustomerByID(t *testing.T) {
	mockCustomerService := new(mocks.MockCustomerService)

	expectedCustomer := &models.Customer{Model: gorm.Model{ID: 1}, Name: "Customer 1", Code: "CUST001"}

	mockCustomerService.On("GetCustomerByID", 1).Return(expectedCustomer, nil)

	customer, err := mockCustomerService.GetCustomerByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedCustomer, customer)

	mockCustomerService.AssertExpectations(t)
}

func TestUpdateCustomer(t *testing.T) {
	mockCustomerService := new(mocks.MockCustomerService)

	customerData := &models.Customer{Name: "Updated Customer", Code: "CUST123"}

	mockCustomerService.On("UpdateCustomer", 1, customerData).Return(nil)

	err := mockCustomerService.UpdateCustomer(1, customerData)

	assert.NoError(t, err)

	mockCustomerService.AssertExpectations(t)
}

func TestDeleteCustomer(t *testing.T) {
	mockCustomerService := new(mocks.MockCustomerService)

	mockCustomerService.On("DeleteCustomer", 1).Return(nil)

	err := mockCustomerService.DeleteCustomer(1)

	assert.NoError(t, err)

	mockCustomerService.AssertExpectations(t)
}
