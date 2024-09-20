package database

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"savannahTest/config"
	"savannahTest/models"
)

var DB *gorm.DB

func DBConnection(logger echo.Logger) error {
	var (
		err error
		dsn string
	)

	// Prepare DSN (Data Source Name)
	dsn = fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Africa/Nairobi",
		config.Configuration.Host, config.Configuration.Port, config.Configuration.User, config.Configuration.Password, config.Configuration.Name,
	)

	// Connect to the database
	if DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		logger.Errorf("failed to connect to database: %v", err)
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Auto migrate the models
	if err := DB.AutoMigrate(&models.Customer{}, &models.Product{}, &models.Order{}); err != nil {
		logger.Errorf("failed to auto migrate models: %v", err)
		return fmt.Errorf("failed to auto migrate models: %v", err)
	}

	return nil
}
