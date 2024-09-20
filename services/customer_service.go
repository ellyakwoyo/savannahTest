package services

import (
	"errors"
	"github.com/labstack/echo/v4"
	"savannahTest/models"
	"savannahTest/repositories"
)

type CustomerService interface {
	CreateCustomer(name, email string) error
	GetAllCustomers() ([]models.Customer, error)
	GetCustomerByID(id int) (*models.Customer, error)
	UpdateCustomer(id int, customerData *models.Customer) error
	DeleteCustomer(id int) error
}

type customerService struct {
	repo   repositories.CustomerRepository
	logger echo.Logger
}

func NewCustomerService(repo repositories.CustomerRepository, logger echo.Logger) CustomerService {
	return &customerService{
		repo:   repo,
		logger: logger,
	}
}

func (s *customerService) CreateCustomer(name, code string) error {
	if name == "" || code == "" {
		s.logger.Error("CreateCustomer: name and code are required")
		return errors.New("name and code are required")
	}

	customer := models.NewCustomer(name, code)
	if err := s.repo.CreateCustomer(customer); err != nil {
		s.logger.Error("CreateCustomer: error creating customer", err)
		return err
	}

	s.logger.Info("CreateCustomer: customer created successfully", customer)
	return nil
}

func (s *customerService) GetAllCustomers() ([]models.Customer, error) {
	customers, err := s.repo.GetAllCustomers()
	if err != nil {
		s.logger.Error("GetAllCustomers: error fetching customers", err)
		return nil, errors.New("could not retrieve customers")
	}

	s.logger.Info("GetAllCustomers: successfully fetched customers")
	return customers, nil
}

func (s *customerService) GetCustomerByID(id int) (*models.Customer, error) {
	customer, err := s.repo.GetCustomerByID(id)
	if err != nil {
		s.logger.Error("GetCustomerByID: customer not found", err)
		return nil, errors.New("customer not found")
	}

	s.logger.Info("GetCustomerByID: customer found", customer)
	return customer, nil
}

func (s *customerService) UpdateCustomer(id int, customerData *models.Customer) error {
	existingCustomer, err := s.repo.GetCustomerByID(id)
	if err != nil {
		s.logger.Error("UpdateCustomer: Customer not found", err)
		return errors.New("Customer not found")
	}

	// Update the existing customer with the new data
	existingCustomer.Name = customerData.Name
	existingCustomer.Code = customerData.Code

	if err := s.repo.UpdateCustomer(existingCustomer); err != nil {
		s.logger.Error("UpdateCustomer: error updating customer", err)
		return err
	}

	s.logger.Info("UpdateCustomer: customer updated successfully", existingCustomer)
	return nil
}

func (s *customerService) DeleteCustomer(id int) error {
	_, err := s.repo.GetCustomerByID(id)
	if err != nil {
		s.logger.Error("DeleteCustomer: customer not found", err)
		return errors.New("customer not found")
	}

	if err := s.repo.DeleteCustomer(id); err != nil {
		s.logger.Error("DeleteCustomer: error deleting customer", err)
		return err
	}

	s.logger.Info("DeleteCustomer: customer deleted successfully", id)
	return nil
}
