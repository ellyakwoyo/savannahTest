package services

import (
	"errors"
	"github.com/labstack/echo/v4"
	"savannahTest/config"
	"savannahTest/models"
	"savannahTest/repositories"
	"savannahTest/utils"
)

type OrderService interface {
	CreateOrder(productID, quantity int, total float64, userID int) error
	GetAllOrders() ([]models.Order, error)
	GetOrderByID(id int) (*models.Order, error)
	UpdateOrder(order *models.Order) error
	DeleteOrder(id int) error
}

type orderService struct {
	repo   repositories.OrderRepository
	logger echo.Logger
}

func NewOrderService(repo repositories.OrderRepository, logger echo.Logger) OrderService {
	return &orderService{
		repo:   repo,
		logger: logger,
	}
}

func (s *orderService) CreateOrder(productID, quantity int, total float64, userID int) error {
	if productID <= 0 || quantity <= 0 || total <= 0 {
		s.logger.Error("CreateOrder: Invalid order details")
		return errors.New("Invalid order details")
	}

	order := models.NewOrder(productID, quantity, total, userID)
	if err := s.repo.CreateOrder(order); err != nil {
		s.logger.Error("CreateOrder: Error creating order", err)
		return err
	}

	apiKey := config.Configuration.SMSSandboxAPIKey
	username := config.Configuration.SMSSandboxUserName
	message := "Your order has been successfully created."

	userPhone := "+254716533839"

	if err := utils.SendSMS(apiKey, username, userPhone, message); err != nil {
		s.logger.Error("CreateOrder: Error sending SMS", err)
		return err
	}

	s.logger.Info("CreateOrder: Order created successfully", order)
	return nil
}

func (s *orderService) GetAllOrders() ([]models.Order, error) {
	orders, err := s.repo.GetAllOrders()
	if err != nil {
		s.logger.Error("GetAllOrders: error fetching orders", err)
		return nil, errors.New("could not retrieve orders")
	}

	s.logger.Info("GetAllOrders: successfully fetched orders")
	return orders, nil
}

func (s *orderService) GetOrderByID(id int) (*models.Order, error) {
	order, err := s.repo.GetOrderByID(id)
	if err != nil {
		s.logger.Error("GetOrderByID: order not found", err)
		return nil, errors.New("order not found")
	}

	s.logger.Info("GetOrderByID: order found", order)
	return order, nil
}

func (s *orderService) UpdateOrder(order *models.Order) error {
	if err := s.repo.UpdateOrder(order); err != nil {
		s.logger.Error("UpdateOrder: error updating order", err)
		return err
	}

	s.logger.Info("UpdateOrder: order updated successfully", order)
	return nil
}

func (s *orderService) DeleteOrder(id int) error {
	_, err := s.repo.GetOrderByID(id)
	if err != nil {
		s.logger.Error("DeleteOrder: order not found", err)
		return errors.New("order not found")
	}

	if err := s.repo.DeleteOrder(id); err != nil {
		s.logger.Error("DeleteOrder: error deleting order", err)
		return err
	}

	s.logger.Info("DeleteOrder: order deleted successfully", id)
	return nil
}
