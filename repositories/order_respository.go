package repositories

import (
	"savannahTest/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *models.Order) error
	GetAllOrders() ([]models.Order, error)
	GetOrderByID(id int) (*models.Order, error)
	UpdateOrder(order *models.Order) error
	DeleteOrder(id int) error
}

type orderRepository struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewOrderRepository(db *gorm.DB, logger echo.Logger) OrderRepository {
	return &orderRepository{
		db:     db,
		logger: logger,
	}
}

func (r *orderRepository) CreateOrder(order *models.Order) error {
	if err := r.db.Create(order).Error; err != nil {
		r.logger.Errorf("Failed to create order: %v", err)
		return err
	}
	return nil
}

func (r *orderRepository) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	if err := r.db.Find(&orders).Error; err != nil {
		r.logger.Errorf("Failed to fetch orders: %v", err)
		return nil, err
	}
	return orders, nil
}

func (r *orderRepository) GetOrderByID(id int) (*models.Order, error) {
	var order models.Order
	if err := r.db.First(&order, id).Error; err != nil {
		r.logger.Errorf("Failed to find order with ID %d: %v", id, err)
		return nil, err
	}
	return &order, nil
}

func (r *orderRepository) UpdateOrder(order *models.Order) error {
	if err := r.db.Save(order).Error; err != nil {
		r.logger.Errorf("Failed to update order: %v", err)
		return err
	}
	return nil
}

func (r *orderRepository) DeleteOrder(id int) error {
	if err := r.db.Delete(&models.Order{}, id).Error; err != nil {
		r.logger.Errorf("Failed to delete order with ID %d: %v", id, err)
		return err
	}
	return nil
}
