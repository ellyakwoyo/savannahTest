package repositories_test

import (
	"gorm.io/gorm"
	"savannahTest/mocks"
	"savannahTest/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	mockRepo := new(mocks.MockOrderRepository)

	order := &models.Order{
		ProductID: 1,
		Quantity:  2,
		Total:     200.0,
		UserId:    1,
	}

	mockRepo.On("CreateOrder", order).Return(nil)

	err := mockRepo.CreateOrder(order)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestGetAllOrders(t *testing.T) {
	mockRepo := new(mocks.MockOrderRepository)

	expectedOrders := []models.Order{
		{
			Model:     gorm.Model{ID: 1},
			ProductID: 1,
			Quantity:  2,
			Total:     200.0,
			UserId:    1,
		},
		{
			Model:     gorm.Model{ID: 2},
			ProductID: 2,
			Quantity:  1,
			Total:     150.0,
			UserId:    1,
		},
	}

	mockRepo.On("GetAllOrders").Return(expectedOrders, nil)

	orders, err := mockRepo.GetAllOrders()

	assert.NoError(t, err)
	assert.Equal(t, expectedOrders, orders)

	mockRepo.AssertExpectations(t)
}

func TestGetOrderByID(t *testing.T) {
	mockRepo := new(mocks.MockOrderRepository)

	expectedOrder := &models.Order{
		Model:     gorm.Model{ID: 1},
		ProductID: 1,
		Quantity:  2,
		Total:     200.0,
		UserId:    1,
	}

	mockRepo.On("GetOrderByID", 1).Return(expectedOrder, nil)

	order, err := mockRepo.GetOrderByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedOrder, order)

	mockRepo.AssertExpectations(t)
}

func TestUpdateOrder(t *testing.T) {
	mockRepo := new(mocks.MockOrderRepository)

	order := &models.Order{
		Model:     gorm.Model{ID: 1},
		ProductID: 1,
		Quantity:  3,
		Total:     300.0,
		UserId:    1,
	}

	mockRepo.On("UpdateOrder", order).Return(nil)

	err := mockRepo.UpdateOrder(order)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteOrder(t *testing.T) {
	mockRepo := new(mocks.MockOrderRepository)

	mockRepo.On("DeleteOrder", 1).Return(nil)

	err := mockRepo.DeleteOrder(1)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
