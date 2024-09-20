package handlers

import (
	"net/http"
	"savannahTest/services"
	"strconv"

	"github.com/labstack/echo/v4"
	"savannahTest/models"
)

type CustomerHandler struct {
	customerService services.CustomerService
}

func NewCustomerHandler(customerService services.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
	}
}

// CreateCustomer godoc
// @Summary Create a new customer
// @Description Create a customer with a name and code.
// @Tags Customer
// @Accept json
// @Produce json
// @Param customer body models.Customer true "Customer data"
// @Success 200 body models.Customer "Customer created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /customers [post]
func (h *CustomerHandler) CreateCustomer(c echo.Context) error {
	customer := new(models.Customer)
	if err := c.Bind(customer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":  http.StatusBadRequest,
			"error": "Invalid request",
		})
	}

	if err := h.customerService.CreateCustomer(customer.Name, customer.Code); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Customer created successfully",
		"data":    customer,
	})
}

// GetAllCustomers godoc
// @Summary Retrieve all customers
// @Description Fetches all customers from the database.
// @Tags Customer
// @Produce json
// @Success 200 {object} map[string]interface{} "Customers fetched successfully"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /customers [get]
func (h *CustomerHandler) GetAllCustomers(c echo.Context) error {
	customers, err := h.customerService.GetAllCustomers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Customers fetched successfully",
		"data":    customers,
	})
}

// GetCustomerByID godoc
// @Summary Retrieve a customer by ID
// @Description Fetches a single customer using their ID.
// @Tags Customer
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {object} map[string]interface{} "Customer fetched successfully"
// @Failure 404 {object} map[string]interface{} "Customer not found"
// @Router /customers/{id} [get]
func (h *CustomerHandler) GetCustomerByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	customer, err := h.customerService.GetCustomerByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":  http.StatusNotFound,
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Customer fetched successfully",
		"data":    customer,
	})
}

// UpdateCustomer godoc
// @Summary Update an existing customer
// @Description Updates the customer details by ID.
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Param customer body models.Customer true "Updated customer data"
// @Success 200 {object} map[string]interface{} "Customer updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /customers/{id} [put]
func (h *CustomerHandler) UpdateCustomer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	customer := new(models.Customer)
	if err := c.Bind(customer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":  http.StatusBadRequest,
			"error": "Invalid request",
		})
	}

	if err := h.customerService.UpdateCustomer(id, customer); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Customer updated successfully",
		"data":    customer,
	})
}

// DeleteCustomer godoc
// @Summary Delete a customer
// @Description Deletes a customer by ID.
// @Tags Customer
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {object} map[string]interface{} "Customer deleted successfully"
// @Failure 404 {object} map[string]interface{} "Customer not found"
// @Router /customers/{id} [delete]
func (h *CustomerHandler) DeleteCustomer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.customerService.DeleteCustomer(id); err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":  http.StatusNotFound,
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Customer deleted successfully",
	})
}
