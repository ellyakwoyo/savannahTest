package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Host               string
	User               string
	Password           string
	Name               string
	Port               int
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURI  string
	SMSSandboxAPIKey   string
	SMSSandboxUserName string
}

var Configuration Config

func ConfigLoader() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	dbPort, err := strconv.Atoi(readEnv("DB_PORT", "5432"))
	if err != nil {
		log.Fatalf("Invalid port number for DB_PORT: %v", err)
	}

	Configuration = Config{
		Host:               readEnv("DB_HOST", "localhost"),
		User:               readEnv("DB_USER", "postgres"),
		Password:           readEnv("DB_PASSWORD", "dev"),
		Port:               dbPort,
		Name:               readEnv("DB_NAME", "savannah-test"),
		SMSSandboxAPIKey:   readEnv("SMS_SANDBOX_API_KEY", ""),
		SMSSandboxUserName: readEnv("SMS_SANDBOX_API_USERNAME", ""),
		GoogleClientID:     readEnv("CLIENT_ID", ""),
		GoogleClientSecret: readEnv("CLIENT_SECRET", ""),
		GoogleRedirectURI:  readEnv("REDIRECT_URI", ""),
		//Secret:             getEnv("SECRET", ""),
	}

	log.Printf("Configuration loaded successfully %+v", Configuration)
	return nil
}

func readEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
