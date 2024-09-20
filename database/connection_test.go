package database

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"savannahTest/config"
)

func TestDBConnection(t *testing.T) {
	// Create a new instance of Echo for logging
	e := echo.New()

	// Load the configuration for the test
	err := config.ConfigLoader()
	assert.NoError(t, err, "Config loading should not return an error")

	// Test database connection
	err = DBConnection(e.Logger)
	if err != nil {
		t.Fatalf("Database connection failed: %v", err)
	}

	// Ensure DB is not nil after a successful connection
	assert.NotNil(t, DB, "Database instance should not be nil after a successful connection")

	// Additional checks can be done here if needed
}
