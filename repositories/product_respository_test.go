package repositories_test

import (
	"gorm.io/gorm"
	"savannahTest/mocks"
	"savannahTest/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	mockRepo := new(mocks.MockProductRepository)

	product := &models.Product{Name: "Test Product", Price: 100.0}
	expectedProduct := &models.Product{Model: gorm.Model{ID: 1}, Name: "Test Product", Price: 100.0} // Ensure ID is properly initialized

	mockRepo.On("Create", product).Return(expectedProduct, nil)

	createdProduct, err := mockRepo.Create(product)

	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, createdProduct)

	mockRepo.AssertExpectations(t)
}

func TestFindProductByID(t *testing.T) {
	mockRepo := new(mocks.MockProductRepository)

	expectedProduct := &models.Product{Model: gorm.Model{ID: 1}, Name: "Test Product", Price: 100.0} // Ensure ID is properly initialized

	mockRepo.On("FindByID", uint(1)).Return(expectedProduct, nil)

	product, err := mockRepo.FindByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, product)

	mockRepo.AssertExpectations(t)
}

func TestUpdateProduct(t *testing.T) {
	mockRepo := new(mocks.MockProductRepository)

	product := &models.Product{Model: gorm.Model{ID: 1}, Name: "Updated Product", Price: 150.0} // Ensure ID is properly initialized

	mockRepo.On("Update", product).Return(nil)

	err := mockRepo.Update(product)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteProduct(t *testing.T) {
	mockRepo := new(mocks.MockProductRepository)

	mockRepo.On("Delete", uint(1)).Return(nil)

	err := mockRepo.Delete(1)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestFindAllProducts(t *testing.T) {
	mockRepo := new(mocks.MockProductRepository)

	expectedProducts := []models.Product{
		{Model: gorm.Model{ID: 1}, Name: "Product 1", Price: 100.0}, // Ensure ID is properly initialized
		{Model: gorm.Model{ID: 2}, Name: "Product 2", Price: 150.0}, // Ensure ID is properly initialized
	}

	mockRepo.On("FindAll").Return(expectedProducts, nil)

	products, err := mockRepo.FindAll()

	assert.NoError(t, err)
	assert.Equal(t, expectedProducts, products)

	mockRepo.AssertExpectations(t)
}
