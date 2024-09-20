package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigLoader(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf("Error loading .env file %v", err)
	}

	fmt.Printf("Configuration: %+v\n", err)
	if err != nil {
		log.Printf("Error loading .env file %v", err)
	}

	dbHost := os.Getenv("DB_HOST")

	fmt.Printf("Host %s", dbHost)
	dbUser := readEnv("DB_USER", "localhost")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	dbName := os.Getenv("DB_NAME")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	redirectURI := os.Getenv("REDIRECT_URI")

	assert.NoError(t, err)

	err = ConfigLoader()
	if err != nil {
		log.Printf("Error loading .env file %v", err)
	}
	assert.Equal(t, dbHost, Configuration.Host)
	assert.Equal(t, dbUser, Configuration.User)
	assert.Equal(t, dbPassword, Configuration.Password)
	assert.Equal(t, dbName, Configuration.Name)
	assert.Equal(t, dbPort, Configuration.Port)
	assert.Equal(t, clientID, Configuration.GoogleClientID)
	assert.Equal(t, clientSecret, Configuration.GoogleClientSecret)
	assert.Equal(t, redirectURI, Configuration.GoogleRedirectURI)
}
