package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"savannahTest/handlers"
	"savannahTest/mocks"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"savannahTest/models"
)

func TestCreateProduct(t *testing.T) {
	e := echo.New()
	mockProductService := new(mocks.MockProductService)
	productHandler := handlers.NewProductHandler(mockProductService)

	mockProductService.On("CreateProduct", "TestProduct", "TestDescription", 10.0).Return(&models.Product{
		Name:        "TestProduct",
		Description: "TestDescription",
		Price:       10.0,
	}, nil)

	productJSON := `{"name": "TestProduct", "description": "TestDescription", "price": 10.0}`
	req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(productJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := productHandler.CreateProduct(c)
	if err != nil {
		t.Fatalf("CreateProduct handler returned an error: %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	mockProductService.AssertExpectations(t)
}

func TestGetAllProducts(t *testing.T) {
	e := echo.New()
	mockProductService := new(mocks.MockProductService)
	productHandler := handlers.NewProductHandler(mockProductService)

	mockProductService.On("GetAllProducts").Return([]models.Product{
		{Name: "Product1", Description: "Description1", Price: 10.0},
		{Name: "Product2", Description: "Description2", Price: 20.0},
	}, nil)

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := productHandler.GetAllProducts(c)
	if err != nil {
		t.Fatalf("GetAllProducts handler returned an error: %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	mockProductService.AssertExpectations(t)
}

func TestGetProductByID(t *testing.T) {
	e := echo.New()
	mockProductService := new(mocks.MockProductService)
	productHandler := handlers.NewProductHandler(mockProductService)

	mockProductService.On("GetProductByID", uint(1)).Return(&models.Product{
		Name:        "TestProduct",
		Description: "TestDescription",
		Price:       10.0,
	}, nil)

	req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := productHandler.GetProductByID(c)
	if err != nil {
		t.Fatalf("GetProductByID handler returned an error: %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	mockProductService.AssertExpectations(t)
}

func TestUpdateProduct(t *testing.T) {
	e := echo.New()
	mockProductService := new(mocks.MockProductService)
	productHandler := handlers.NewProductHandler(mockProductService)

	mockProductService.On("UpdateProduct", uint(1), mock.Anything).Return(nil)

	productJSON := `{"name": "UpdatedProduct", "description": "UpdatedDescription", "price": 15.0}`
	req := httptest.NewRequest(http.MethodPut, "/products/1", strings.NewReader(productJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := productHandler.UpdateProduct(c)
	if err != nil {
		t.Fatalf("UpdateProduct handler returned an error: %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	mockProductService.AssertExpectations(t)
}

func TestDeleteProduct(t *testing.T) {
	e := echo.New()
	mockProductService := new(mocks.MockProductService)
	productHandler := handlers.NewProductHandler(mockProductService)

	mockProductService.On("DeleteProduct", uint(1)).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := productHandler.DeleteProduct(c)
	if err != nil {
		t.Fatalf("DeleteProduct handler returned an error: %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	mockProductService.AssertExpectations(t)
}
