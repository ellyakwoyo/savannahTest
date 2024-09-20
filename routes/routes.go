package routes

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"savannahTest/handlers"
	"savannahTest/middlewares"
	"savannahTest/services"
)

func SetupRoutes(e *echo.Echo, productService services.ProductService, orderService services.OrderService, customerService services.CustomerService) {
	e.GET("/docs/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", handlers.Home)

	apiV1 := e.Group("/app/v1")

	apiV1.GET("/login", handlers.Login)
	apiV1.GET("/auth/google/callback", handlers.HandleGoogleCallback)

	productHandler := handlers.NewProductHandler(productService)
	products := apiV1.Group("/products")
	products.GET("", productHandler.GetAllProducts, middlewares.AuthenticateMiddleware())
	products.GET("/:id", productHandler.GetProductByID, middlewares.AuthenticateMiddleware())
	products.POST("", productHandler.CreateProduct, middlewares.AuthenticateMiddleware())
	products.PUT("/:id", productHandler.UpdateProduct, middlewares.AuthenticateMiddleware())
	products.DELETE("/:id", productHandler.DeleteProduct, middlewares.AuthenticateMiddleware())

	orderHandler := handlers.NewOrderHandler(orderService)
	orders := apiV1.Group("/orders")
	orders.GET("", orderHandler.GetAllOrders, middlewares.AuthenticateMiddleware())
	orders.GET("/:id", orderHandler.GetOrderByID, middlewares.AuthenticateMiddleware())
	orders.POST("", orderHandler.CreateOrder, middlewares.AuthenticateMiddleware())
	orders.PUT("/:id", orderHandler.UpdateOrder, middlewares.AuthenticateMiddleware())
	orders.DELETE("/:id", orderHandler.DeleteOrder, middlewares.AuthenticateMiddleware())

	customerHandler := handlers.NewCustomerHandler(customerService)
	customers := apiV1.Group("/customers")
	customers.GET("", customerHandler.GetAllCustomers, middlewares.AuthenticateMiddleware())
	customers.GET("/:id", customerHandler.GetCustomerByID, middlewares.AuthenticateMiddleware())
	customers.POST("", customerHandler.CreateCustomer, middlewares.AuthenticateMiddleware())
	customers.PUT("/:id", customerHandler.UpdateCustomer, middlewares.AuthenticateMiddleware())
	customers.DELETE("/:id", customerHandler.DeleteCustomer, middlewares.AuthenticateMiddleware())
}
