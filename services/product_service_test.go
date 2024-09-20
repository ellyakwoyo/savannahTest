package services_test

import (
	"gorm.io/gorm"
	"savannahTest/mocks"
	"savannahTest/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProductService(t *testing.T) {

	mockProductService := new(mocks.MockProductService)

	product := &models.Product{Name: "Test Product", Description: "Test Description", Price: 100.0}
	mockProductService.On("CreateProduct", "Test Product", "Test Description", 100.0).Return(product, nil)

	createdProduct, err := mockProductService.CreateProduct("Test Product", "Test Description", 100.0)

	assert.NoError(t, err)
	assert.NotNil(t, createdProduct)
	assert.Equal(t, "Test Product", createdProduct.Name)
	assert.Equal(t, "Test Description", createdProduct.Description)
	assert.Equal(t, 100.0, createdProduct.Price)

	mockProductService.AssertExpectations(t)
}

func TestGetProductByIDService(t *testing.T) {
	mockProductService := new(mocks.MockProductService)

	product := &models.Product{Model: gorm.Model{ID: 1}, Name: "Test Product", Description: "Test Description", Price: 100.0}

	mockProductService.On("GetProductByID", uint(1)).Return(product, nil)

	foundProduct, err := mockProductService.GetProductByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, foundProduct)
	assert.Equal(t, "Test Product", foundProduct.Name)
	assert.Equal(t, "Test Description", foundProduct.Description)
	assert.Equal(t, 100.0, foundProduct.Price)

	mockProductService.AssertExpectations(t)
}

func TestUpdateProductService(t *testing.T) {

	mockProductService := new(mocks.MockProductService)

	product := &models.Product{Model: gorm.Model{ID: 1}, Name: "Updated Product", Description: "Updated Description", Price: 200.0}

	mockProductService.On("UpdateProduct", uint(1), product).Return(nil)

	err := mockProductService.UpdateProduct(1, product)

	assert.NoError(t, err)

	mockProductService.AssertExpectations(t)
}

func TestDeleteProductService(t *testing.T) {

	mockProductService := new(mocks.MockProductService)

	mockProductService.On("DeleteProduct", uint(1)).Return(nil)

	err := mockProductService.DeleteProduct(1)

	assert.NoError(t, err)

	mockProductService.AssertExpectations(t)
}

func TestGetAllProductsService(t *testing.T) {

	mockProductService := new(mocks.MockProductService)

	expectedProducts := []models.Product{
		{Model: gorm.Model{ID: 1}, Name: "Product 1", Description: "Description 1", Price: 100.0},
		{Model: gorm.Model{ID: 2}, Name: "Product 2", Description: "Description 2", Price: 150.0},
	}

	mockProductService.On("GetAllProducts").Return(expectedProducts, nil)

	products, err := mockProductService.GetAllProducts()

	assert.NoError(t, err)
	assert.Equal(t, expectedProducts, products)

	mockProductService.AssertExpectations(t)
}
