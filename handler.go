package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

type Product struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Picture     string  `json:"picture"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

var conn *pgx.Conn

func createProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	json.NewDecoder(r.Body).Decode(&product)

	insertQuery := "INSERT INTO products (name, type, picture, price, description) VALUES ($1, $2, $3, $4, $5) RETURNING id"

	err := conn.QueryRow(context.Background(), insertQuery, product.Name, product.Type, product.Picture, product.Price, product.Description).Scan(&product.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var product Product
	productQuery := "SELECT id, name, type, picture, price, description FROM products WHERE id=$1"
	err = conn.QueryRow(context.Background(), productQuery, id).
		Scan(&product.ID, &product.Name, &product.Type, &product.Picture, &product.Price, &product.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, name, type, picture, price, description FROM products"
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, "Products not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Picture, &product.Price, &product.Description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var product Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = conn.QueryRow(context.Background(),
		`UPDATE products SET name = $1, type = $2, picture = $3, price = $4, description = $5
        WHERE id = $6 RETURNING id`, product.Name, product.Type, product.Picture, product.Price, product.Description, id).Scan(&product.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func searchProducts(w http.ResponseWriter, r *http.Request) {
	nameParam := r.URL.Query().Get("name")
	if nameParam == "" {
		http.Error(w, "Name parameter is required", http.StatusBadRequest)
		return
	}

	var products []Product
	searchQuery := `SELECT id, name, type, picture, price, description FROM products WHERE name ILIKE $1`

	rows, err := conn.Query(context.Background(), searchQuery, "%"+nameParam+"%")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		fmt.Println("SEARCH API Rows INSIDE FOR LOOP: ", rows)
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Picture, &product.Price, &product.Description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Product Management System!!!"))
}
