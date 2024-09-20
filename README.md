# Golang Project

This is a simple golang service that has followed SOLID principle of API development with automated unit tests. The project manages customers and orders, authentication based on OpenID Connect, and SMS notifications for placing orders. It also has a set of strict and comprehensive tests for every unit, CI/CD pipeline configuration and step-by-step deployment guide. On login, you will see JSON for loggedin user credentials and session token which should be needed in frontend.
## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Authentication and Authorization](#authentication-and-authorization)
- [SMS Notifications](#sms-notifications)
- [Testing](#testing)
- [Project Structure](#project-structure)
- [Urls](#urls)

## Features
This project includes the following features:
- Management of customers and orders.
- REST API for creating and retrieving customers and orders.
- Authentication and authorization via OpenID Connect (Google authentication).
- SMS notifications through Africa's Talking SMS gateway.
- Comprehensive unit tests with coverage verification.
- Relational database using Postgresql.
- CI/CD pipeline setup.
- Detailed deployment instructions.

## Prerequisites

- Go 1.23+
- Docker (optional, for database setup)
- Africa’s Talking account for SMS service
- OpenID Connect provider for authentication

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/ellyakwoyo/savannahTest.git
    cd savannahTest
    ```

2. Install Go dependencies:
    ```bash
    go mod tidy
    ```

## Configuration

Ensure you have a `.env` file in the `savannahTest` having the following information:
```env
DB_HOST = your_db_host
DB_USER = your_db_user
DB_PASSWORD = your_db_password
DB_NAME = your_db_name
DB_PORT = your_db_port
CLIENT_ID = your_google_client_id
CLIENT_SECRET = your_google_client_secret
REDIRECT_URI = https://savannah-435913.ue.r.appspot.com/app/v1/auth/google/callback
SMS_SANDBOX_API_KEY = your_sms_sandbox_api_key
SMS_SANDBOX_API_USERNAME = your_sms_sandbox_api_username
```
Also check `env.example`

## Running the Application

1. Start the project server:
    ```bash
    cd savannahTest
    go run main.go
    ```

## API Endpoints

### Customers
- **GET** `/customers` - Fetch all customers
- **POST** `/customers` - Post a new customer
    - Request body:
        ```json
        {
            "name": "Mike Dean",
            "code": "C001"
        }
        ```

### Orders
- **GET** `/orders` - Fetch all orders
  - **POST** `/orders` - Post a new order
      - Request body:
          ```json
        {
            "ProductID": 1,
            "Quantity": 4,
            "Total" : 5.0,
            "UserId" : 1
        }
          ```

## Authentication and Authorization

The application uses OpenID Connect for authentication and authorization. Ensure you have configured the OIDC provider details in the (Use google) `.env` file.

## SMS Notifications

The application sends SMS notifications using Africa’s Talking SMS gateway. Ensure you have configured your Africa’s Talking API key and username in the `.env` file.

## Testing

Run savannahTest tests:
```bash
cd savannahTest
go test ./...
```
## Project Structure
```
savannahTest/
├── app.yaml
├── authentication
│   ├── authentication.go
│   └── authentication_test.go
├── config
│   ├── config.go
│   ├── config_test.go
│   └── env.example
├── database
│   ├── connection.go
│   └── connection_test.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── handlers
│   ├── auth_handler.go
│   ├── customer_handler.go
│   ├── customer_handler_test.go
│   ├── order_handler.go
│   ├── order_handler_test.go
│   ├── product_handler.go
│   └── product_handler_test.go
├── main.go
├── main_test.go
├── Makefile
├── middlewares
│   └── auth.go
├── mocks
│   ├── mock_customer_repository.go
│   ├── mock_customer_service.go
│   ├── mock_order_repository.go
│   ├── mock_order_service.go
│   ├── mock_product_service.go
│   └── product_repository.go
├── models
│   ├── customer.go
│   ├── order.go
│   └── product.go
├── README.md
├── repositories
│   ├── customer_repository.go
│   ├── customer_repository_test.go
│   ├── order_repository_test.go
│   ├── order_respository.go
│   ├── product_repository.go
│   └── product_respository_test.go
├── routes
│   └── routes.go
├── server
│   ├── server.go
│   └── server_test.go
├── services
│   ├── customer_service.go
│   ├── customer_service_test.go
│   ├── order_service.go
│   ├── order_service_test.go
│   ├── product_service.go
│   └── product_service_test.go
└── utils
    └── sms.go
```
## Urls

- Deployed URL: [https://savannah-435913.ue.r.appspot.com/](https://savannah-435913.ue.r.appspot.com/)
- Docs URL: [https://savannah-435913.ue.r.appspot.com/docs/swagger/index.html](https://savannah-435913.ue.r.appspot.com/docs/swagger/index.html)