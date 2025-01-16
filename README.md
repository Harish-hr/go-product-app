# Go Product App

This application is a simple product management system built with Go, utilizing a PostgreSQL database and the Gin framework for RESTful API functionalities.  It allows users to perform CRUD operations on product data.

## Features

* **Create Product:** Add new products with details like name, type, picture URL, price, and description.
* **Read Product:** Retrieve product details by ID.
* **Read All Products:** List all available products.
* **Update Product:** Modify existing product information.
* **Search Product:** Find products by name using partial or full string matching.
* **Home Page:** A simple welcome message.

## Steps to run application in local environment

* **Step 1:** Docker desktop should be installed on local machine. [docker --version (Docker version 27.4.0, build bde2b89)]
* **Step 2:** Download or clone project into local machine. Navigate to the directory containing your docker-compose.yml file.
* **Step 3:** Run the command: "docker compose up -d"
* **Step 4:** Open browser normal or in-cognoto mode hit "http://localhost:8080"
* **Step 5:** Message "Welcome to the Product Management System!!!" confirms that backend server up and runnning

## Technologies Used

* **Go:** The programming language used for developing the application logic.
* **PostgreSQL:** The database used for storing product data.
* **Gin:** A high-performance HTTP web framework for Go.
* **Gorilla Mux:** A powerful URL router and dispatcher for matching incoming requests to their respective handler.
* **pgx:** A PostgreSQL driver for Go.
* **godotenv:** For loading environment variables.
* **CORS:**  Handles Cross-Origin Resource Sharing, enabling requests from different domains.
