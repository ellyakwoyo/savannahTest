package services

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"savannahTest/models"
	"savannahTest/repositories"
)

type ProductService interface {
	CreateProduct(name, description string, price float64) (*models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	UpdateProduct(id uint, productData *models.Product) error
	DeleteProduct(id uint) error
	GetAllProducts() ([]models.Product, error)
}

type productService struct {
	repo   repositories.ProductRepository
	logger echo.Logger
}

func NewProductService(repo repositories.ProductRepository, logger echo.Logger) ProductService {
	return &productService{
		repo:   repo,
		logger: logger,
	}
}

func (s *productService) CreateProduct(name, description string, price float64) (*models.Product, error) {
	if name == "" || description == "" || price <= 0 {
		s.logger.Error("CreateProduct: Name, description, and price are required and price must be positive")
		return nil, errors.New("name, description, and valid price are required")
	}

	product := models.NewProduct(name, description, price)

	createdProduct, err := s.repo.Create(product)
	if err != nil {
		s.logger.Errorf("CreateProduct: Failed to create product: %v", err)
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	s.logger.Infof("CreateProduct: Product created successfully: %+v", createdProduct)
	return createdProduct, nil
}

func (s *productService) GetProductByID(id uint) (*models.Product, error) {
	product, err := s.repo.FindByID(id)
	if err != nil {
		s.logger.Errorf("GetProductByID: Failed to find product with ID %d: %v", id, err)
		return nil, fmt.Errorf("product not found: %w", err)
	}

	s.logger.Infof("GetProductByID: Product found: %+v", product)
	return product, nil
}

func (s *productService) UpdateProduct(id uint, productData *models.Product) error {

	product, err := s.repo.FindByID(id)
	if err != nil {
		s.logger.Errorf("UpdateProduct: Product with ID %d not found: %v", id, err)
		return errors.New("product not found")
	}

	product.Name = productData.Name
	product.Description = productData.Description
	product.Price = productData.Price

	if err := s.repo.Update(product); err != nil {
		s.logger.Errorf("UpdateProduct: Failed to update product with ID %d: %v", id, err)
		return fmt.Errorf("failed to update product: %w", err)
	}

	s.logger.Infof("UpdateProduct: Product updated successfully: %+v", product)
	return nil
}

func (s *productService) DeleteProduct(id uint) error {

	_, err := s.repo.FindByID(id)
	if err != nil {
		s.logger.Errorf("DeleteProduct: Product with ID %d not found: %v", id, err)
		return errors.New("product not found")
	}

	if err := s.repo.Delete(id); err != nil {
		s.logger.Errorf("DeleteProduct: Failed to delete product with ID %d: %v", id, err)
		return fmt.Errorf("failed to delete product: %w", err)
	}

	s.logger.Infof("DeleteProduct: Product with ID %d deleted", id)
	return nil
}

func (s *productService) GetAllProducts() ([]models.Product, error) {
	products, err := s.repo.FindAll()
	if err != nil {
		s.logger.Errorf("GetAllProducts: Failed to retrieve products: %v", err)
		return nil, fmt.Errorf("failed to retrieve products: %w", err)
	}

	s.logger.Infof("GetAllProducts: Retrieved %d products", len(products))
	return products, nil
}
