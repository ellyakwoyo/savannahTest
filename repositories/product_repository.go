package repositories

import (
	"fmt"
	"savannahTest/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *models.Product) (*models.Product, error)
	FindByID(id uint) (*models.Product, error)
	Update(product *models.Product) error
	Delete(id uint) error
	FindAll() ([]models.Product, error)
}

type productRepository struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewProductRepository(db *gorm.DB, logger echo.Logger) ProductRepository {
	return &productRepository{
		db:     db,
		logger: logger,
	}
}

func (r *productRepository) Create(product *models.Product) (*models.Product, error) {
	if err := r.db.Create(product).Error; err != nil {
		r.logger.Errorf("Failed to create product: %v", err)
		return nil, fmt.Errorf("Failed to create product: %w", err)
	}
	r.logger.Infof("Product created: %v", product)
	return product, nil
}

func (r *productRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := r.db.First(&product, id).Error; err != nil {
		r.logger.Errorf("Failed to find product by ID %d: %v", id, err)
		return nil, fmt.Errorf("Failed to find product: %w", err)
	}
	r.logger.Infof("Product found: %v", product)
	return &product, nil
}

func (r *productRepository) Update(product *models.Product) error {
	if err := r.db.Save(product).Error; err != nil {
		r.logger.Errorf("Failed to update product: %v", err)
		return fmt.Errorf("Failed to update product: %w", err)
	}
	r.logger.Infof("Product updated: %v", product)
	return nil
}

func (r *productRepository) Delete(id uint) error {
	if err := r.db.Delete(&models.Product{}, id).Error; err != nil {
		r.logger.Errorf("Failed to delete product with ID %d: %v", id, err)
		return fmt.Errorf("Failed to delete product: %w", err)
	}
	r.logger.Infof("Product with ID %d deleted", id)
	return nil
}

func (r *productRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	if err := r.db.Find(&products).Error; err != nil {
		r.logger.Errorf("Failed to find all products: %v", err)
		return nil, fmt.Errorf("Failed to find products: %w", err)
	}
	r.logger.Infof("Found %d products", len(products))
	return products, nil
}
