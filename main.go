package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"savannahTest/authentication"
	"savannahTest/config"
	"savannahTest/database"
	_ "savannahTest/docs"
	"savannahTest/repositories"
	"savannahTest/routes"
	"savannahTest/server"
	"savannahTest/services"
)

// @title SavannahTest API
// @version 1.0
// @contact.email support@swagger.io
// SetupRoutes @host localhost:8080

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	err := config.ConfigLoader()
	if err != nil {
		e.Logger.Error(err.Error())
	}

	if err := database.DBConnection(e.Logger); err != nil {
		e.Logger.Fatal(err.Error())
	}

	authentication.InitializeGoogleOAuth()

	productRepo := repositories.NewProductRepository(database.DB, e.Logger)
	orderRepo := repositories.NewOrderRepository(database.DB, e.Logger)
	customerRepo := repositories.NewCustomerRepository(database.DB, e.Logger)

	productService := services.NewProductService(productRepo, e.Logger)
	orderService := services.NewOrderService(orderRepo, e.Logger)
	customerService := services.NewCustomerService(customerRepo, e.Logger)

	routes.SetupRoutes(e, productService, orderService, customerService)

	server.RunServer(e)
}
