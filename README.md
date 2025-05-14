# GoLang-Gin-boilerplate

Starter template for scalable Go applications using the Gin web framework.

## File Structure


GoLang-Gin-boilerplate/
│── cmd/
│   └── server/
│       └── main.go         # Entry point of the app
│── compose/
│   └── local/
│       ├── api/
│       └── postgres
│── config/
│   └── config.go       # Database & environment config
│── internal/
│   ├── controllers/    # Handle HTTP requests
│   ├── database/
│   │   ├── models/       # Database models
│   │   ├── repositories/ # DB interactions
│   ├── services/       # Business logic
│   ├── middleware/
│── scripts/
│── test/
│── routes/
│   └── routes.go       # API route definitions
│── .env                # Environment variables
│── go.mod              # Go module file
│── go.sum              # Dependency checksum
│── docker-compose.yml  # Docker Compose file
│── README.md           # Documentation


## Prerequisites

* **Go:** Make sure you have Go installed. You can check your Go version by running:

    ```bash
    go version
    ```

## Installation

1.  **Create a New Project:**

    ```bash
    go mod init [github.com/yourusername/$PROJECT_NAME]
    In my Case
    go mod init [github.com/inikhildubey/GoLang-GinCURD](https://github.com/inikhildubey/GoLang-GinCURD)
    ```

2.  **Install Gin Framework:**

    ```bash
    go get -u [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
    ```

3.  **Install godotenv:**

    ```bash
    go get [github.com/joho/godotenv](https://github.com/joho/godotenv)
    ```

4.  **Install GORM:**

    ```bash
    go get -u gorm.io/gorm
    ```

5.  **Install Postgres Driver for GORM:**

    ```bash
    go get -u gorm.io/driver/postgres
    ```

## Configuration

* **.env:** This file is used to store environment variables.  Make sure to create a `.env` file in the root directory of your project.  Example:

    ```
    PORT=8080
    DB_HOST=localhost
    DB_USER=your_user
    DB_PASSWORD=your_password
    DB_NAME=your_database
    ```

* **config/config.go:** This file handles application configuration, including database connections and environment settings.

## Running the Application

1.  **Using Docker Compose (Recommended):**

    * Make sure you have Docker and Docker Compose installed.
    * Run the application using Docker Compose:

        ```bash
        docker-compose up -d
        ```

2.  **Without Docker:**
    * Install Postgres and create a database.
    * cd into `cmd/server`
    * Set the required environment variables.
    * Run the `main.go` file:

        ```bash
        go run main.go
        ```

## Project Structure Details

* **cmd/server/main.go:** This is the entry point of the application. It initializes the Gin router, loads configurations, and starts the server.
* **compose/local/:** Contains Docker Compose configurations for local development, including setting up the application and a Postgres database.
* **config/:** Handles application configuration, such as database connections, and loading environment variables.
* **internal/:**
    * **controllers/:** Contains the handlers for HTTP requests.  These handlers process incoming requests, interact with services, and return responses.
    * **database/models/:** Defines the structure of the database tables using GORM.
    * **database/repositories:** Contains the logic for interacting with the database.
    * **services/:** Contains the business logic of the application.  Controllers use services to perform operations.
    * **middleware/:** Contains middleware functions that intercept HTTP requests, for example, for authentication or logging.
* **routes/routes.go:** Defines the API routes for the application using the Gin router.
* **scripts/:** (Optional)  Can contain scripts for database migrations, setup, or other tasks.
* **test/:** (Optional) Contains tests.

## Dependencies

* [Gin](https://github.com/gin-gonic/gin):  A high-performance web framework for Go.
* [godotenv](https://github.com/joho/godotenv):  For loading environment variables from `.env` files.
* [GORM](https://gorm.io/): The fantastic ORM library for Go, aims to be developer friendly.
* [Postgres Driver](https://github.com/jackc/pgx): For connecting to the PostgreSQL database.