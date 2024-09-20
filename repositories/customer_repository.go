package repositories

import (
	"savannahTest/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	CreateCustomer(customer *models.Customer) error
	GetAllCustomers() ([]models.Customer, error)
	GetCustomerByID(id int) (*models.Customer, error)
	UpdateCustomer(customer *models.Customer) error
	DeleteCustomer(id int) error
}

type customerRepository struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewCustomerRepository(db *gorm.DB, logger echo.Logger) CustomerRepository {
	return &customerRepository{
		db:     db,
		logger: logger,
	}
}

func (r *customerRepository) CreateCustomer(customer *models.Customer) error {
	if err := r.db.Create(customer).Error; err != nil {
		r.logger.Errorf("Failed to create customer: %v", err)
		return err
	}
	return nil
}

func (r *customerRepository) GetAllCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	if err := r.db.Find(&customers).Error; err != nil {
		r.logger.Errorf("Failed to fetch customers: %v", err)
		return nil, err
	}
	return customers, nil
}

func (r *customerRepository) GetCustomerByID(id int) (*models.Customer, error) {
	var customer models.Customer
	if err := r.db.First(&customer, id).Error; err != nil {
		r.logger.Errorf("Failed to find customer with ID %d: %v", id, err)
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) UpdateCustomer(customer *models.Customer) error {
	if err := r.db.Save(customer).Error; err != nil {
		r.logger.Errorf("Failed to update customer: %v", err)
		return err
	}
	return nil
}

func (r *customerRepository) DeleteCustomer(id int) error {
	if err := r.db.Delete(&models.Customer{}, id).Error; err != nil {
		r.logger.Errorf("failed to delete customer with ID %d: %v", id, err)
		return err
	}
	return nil
}
