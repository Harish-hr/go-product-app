package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database
	conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	// allow all origins
	c := cors.New(cors.Options{AllowedOrigins: []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true})

	// Initialize the router
	r := mux.NewRouter()

	// create product
	r.HandleFunc("/products", createProduct).Methods("POST")

	// get all products
	r.HandleFunc("/products", getProducts).Methods("GET")

	// search product by name
	r.HandleFunc("/products/search", searchProducts).Methods("GET")

	// get product by ID
	r.HandleFunc("/products/{id}", getProduct).Methods("GET")

	// update product
	r.HandleFunc("/products/{id}", updateProduct).Methods("PUT")

	// home page
	r.HandleFunc("/", homeHandler).Methods("GET")
	handler := c.Handler(r)

	// query to create table "products"
	createProductQuery := `
	CREATE TABLE IF NOT EXISTS products (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
		type VARCHAR(100) NOT NULL,
		picture VARCHAR(255),
		price NUMERIC(10, 2) NOT NULL,
		description TEXT);`

	_, err = conn.Exec(context.Background(), createProductQuery)
	if err != nil {
		log.Fatal(err)
	}

	// start server
	log.Fatal(http.ListenAndServe(":8080", handler))
}
