# Cimri Internship Case - User Service

This repository contains the **User Service**, which is a part of a **microservices-based architecture** developed in Go. The system consists of the following core services:

- **User Service** (this repository): Handles user-related operations, including managing user profiles, preferences, and interactions.
- **Product Service**: Manages product data and retrieves product details based on product ID.
- **Favorites Service**: Enables users to create, update, delete, and view their favorite lists. Users can also add or remove products from these lists. This service communicates with both the **User Service** and the **Product Service**.

The **User Service** is built with **Go**, **Fiber** (for HTTP routing), and **GORM** (for ORM database interactions). Unit tests are included to ensure reliability.

## 📌 Technologies Used

- **Go (Golang)**
- **Fiber** (Web Framework)
- **GORM** (Object Relational Mapping)
- **PostgreSQL**
- **Redis**
- **Docker & Docker Compose**
- **AWS EC2, RDS, S3**
- **Swagger / OpenAPI**

## 📂 Project Structure

```
/cmd
  /main.bo          # Entry point for the HTTP server
/internal
  /handler         # Handles HTTP requests and responses, including unit tests for handlers
  /service         # Contains business logic
  /repository      # Handles database interactions, including unit tests for repository
  /models          # Defines data models (using GORM)
/utils
  /envloader       # Loads environment variables
/pkg
  /redis           # Handles Redis connection
  /s3              # Handles S3 connection
  /postgres        # Handles PostgreSQL connection (with GORM)
```

### Key Points:

- **Fiber**: The application uses **Fiber** as the web framework to handle routing and HTTP requests efficiently.
- **GORM**: **GORM** is used for ORM-based database interactions with **PostgreSQL**.
- **Tests**: Unit tests for **handlers** and **repository** are written directly within the respective packages (e.g., `handler/user_handler_test.go`, `repository/user_repository_test.go`).
- **`/utils`**: Contains the **env loader**.
- **`/pkg`**: Contains utilities to handle connections to **Redis**, **S3**, and **PostgreSQL**.
- **Main Function**: The **main** function is located in the `/cmd` directory.

## 🚀 Getting Started

### 1️⃣ Install Dependencies

Ensure Go modules are up to date:

```sh
go mod tidy
```

### 2️⃣ Run with Docker

To start the service along with PostgreSQL and Redis:

```sh
docker-compose up --build
```

### 3️⃣ Run Manually (Without Docker)

If you prefer to run the service manually:

```sh
go run cmd/main.go
```

Make sure to configure your `.env` file correctly before running the service.

## ✅ Running Unit Tests

This repository includes unit tests for the repository and handler layers. To execute all tests:

```sh
go test ./...
```

To run tests for a specific package:

```sh
go test ./internal/handler
```

Unit tests utilize mock data, so no real database connection is required.

## 📖 API Documentation

API endpoints are documented using Swagger/OpenAPI. Once the service is running, access the API documentation at:

```
http://localhost:8080/swagger
```
