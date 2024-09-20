package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"savannahTest/handlers"
	"savannahTest/mocks"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"savannahTest/models"
)

func TestCreateCustomer(t *testing.T) {
	e := echo.New()
	mockCustomerService := new(mocks.MockCustomerService)
	customerHandler := handlers.NewCustomerHandler(mockCustomerService)

	mockCustomerService.On("CreateCustomer", "John Doe", "C123").Return(nil)

	customerJSON := `{"name": "John Doe", "code": "C123"}`
	req := httptest.NewRequest(http.MethodPost, "/customers", strings.NewReader(customerJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := customerHandler.CreateCustomer(c)
	if err != nil {
		t.Fatalf("CreateCustomer handler returned an error: %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	mockCustomerService.AssertExpectations(t)
}

func TestGetAllCustomers(t *testing.T) {
	e := echo.New()
	mockCustomerService := new(mocks.MockCustomerService)
	customerHandler := handlers.NewCustomerHandler(mockCustomerService)

	mockCustomerService.On("GetAllCustomers").Return([]models.Customer{
		{Name: "John Doe", Code: "C123"},
		{Name: "Jane Smith", Code: "C456"},
	}, nil)

	req := httptest.NewRequest(http.MethodGet, "/customers", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := customerHandler.GetAllCustomers(c)
	if err != nil {
		t.Fatalf("GetAllCustomers handler returned an error: %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	mockCustomerService.AssertExpectations(t)
}

func TestGetCustomerByID(t *testing.T) {
	e := echo.New()
	mockCustomerService := new(mocks.MockCustomerService)
	customerHandler := handlers.NewCustomerHandler(mockCustomerService)

	customerID := 1
	mockCustomerService.On("GetCustomerByID", customerID).Return(&models.Customer{
		Name: "John Doe", Code: "C123",
	}, nil)

	req := httptest.NewRequest(http.MethodGet, "/customers/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(customerID))

	err := customerHandler.GetCustomerByID(c)
	if err != nil {
		t.Fatalf("GetCustomerByID handler returned an error: %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	mockCustomerService.AssertExpectations(t)
}

func TestUpdateCustomer(t *testing.T) {
	e := echo.New()
	mockCustomerService := new(mocks.MockCustomerService)
	customerHandler := handlers.NewCustomerHandler(mockCustomerService)

	customerID := 1
	customer := &models.Customer{Name: "John Doe", Code: "C123"}
	mockCustomerService.On("UpdateCustomer", customerID, customer).Return(nil)

	customerJSON := `{"name": "John Doe", "code": "C123"}`
	req := httptest.NewRequest(http.MethodPut, "/customers/1", strings.NewReader(customerJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(customerID))

	err := customerHandler.UpdateCustomer(c)
	if err != nil {
		t.Fatalf("UpdateCustomer handler returned an error: %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	mockCustomerService.AssertExpectations(t)
}

func TestDeleteCustomer(t *testing.T) {
	e := echo.New()
	mockCustomerService := new(mocks.MockCustomerService)
	customerHandler := handlers.NewCustomerHandler(mockCustomerService)

	customerID := 1
	mockCustomerService.On("DeleteCustomer", customerID).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/customers/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(customerID))

	err := customerHandler.DeleteCustomer(c)
	if err != nil {
		t.Fatalf("DeleteCustomer handler returned an error: %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	mockCustomerService.AssertExpectations(t)
}
