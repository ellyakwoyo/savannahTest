package repositories_test

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"savannahTest/mocks"
	"savannahTest/models"
	"testing"
)

func TestCreateCustomer(t *testing.T) {
	mockRepo := new(mocks.MockCustomerRepository)

	customer := &models.Customer{Name: "John Doe", Code: "JD123"}

	mockRepo.On("CreateCustomer", customer).Return(nil)

	err := mockRepo.CreateCustomer(customer)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestGetAllCustomers(t *testing.T) {
	mockRepo := new(mocks.MockCustomerRepository)

	expectedCustomers := []models.Customer{
		{
			Model: gorm.Model{ID: 1},
			Name:  "John Doe",
			Code:  "JD123",
		},
		{
			Model: gorm.Model{ID: 2},
			Name:  "Jane Doe",
			Code:  "JD456",
		},
	}

	mockRepo.On("GetAllCustomers").Return(expectedCustomers, nil)

	customers, err := mockRepo.GetAllCustomers()

	assert.NoError(t, err)
	assert.Equal(t, expectedCustomers, customers)

	mockRepo.AssertExpectations(t)
}

func TestGetCustomerByID(t *testing.T) {
	mockRepo := new(mocks.MockCustomerRepository)

	expectedCustomer := &models.Customer{
		Model: gorm.Model{ID: 1},
		Name:  "John Doe",
		Code:  "JD123",
	}

	mockRepo.On("GetCustomerByID", 1).Return(expectedCustomer, nil)

	customer, err := mockRepo.GetCustomerByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedCustomer, customer)

	mockRepo.AssertExpectations(t)
}

func TestUpdateCustomer(t *testing.T) {
	mockRepo := new(mocks.MockCustomerRepository)

	customer := &models.Customer{
		Model: gorm.Model{ID: 1},
		Name:  "Updated Name",
		Code:  "UC123",
	}

	mockRepo.On("UpdateCustomer", 1, customer).Return(nil)

	err := mockRepo.UpdateCustomer(1, customer)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteCustomer(t *testing.T) {
	mockRepo := new(mocks.MockCustomerRepository)

	mockRepo.On("DeleteCustomer", 1).Return(nil)

	err := mockRepo.DeleteCustomer(1)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
