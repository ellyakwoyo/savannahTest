package main

import (
	"net/http"
	"net/http/httptest"
	"savannahTest/authentication"
	"savannahTest/config"
	"savannahTest/database"
	"savannahTest/repositories"
	"savannahTest/routes"
	"savannahTest/services"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMainFunction(t *testing.T) {
	e := echo.New()

	err := config.ConfigLoader()
	assert.NoError(t, err)

	err = database.DBConnection(e.Logger)
	assert.NoError(t, err)

	authentication.InitializeGoogleOAuth()

	productRepo := repositories.NewProductRepository(database.DB, e.Logger)
	orderRepo := repositories.NewOrderRepository(database.DB, e.Logger)
	customerRepo := repositories.NewCustomerRepository(database.DB, e.Logger)

	productService := services.NewProductService(productRepo, e.Logger)
	orderService := services.NewOrderService(orderRepo, e.Logger)
	customerService := services.NewCustomerService(customerRepo, e.Logger)

	routes.SetupRoutes(e, productService, orderService, customerService)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

}
