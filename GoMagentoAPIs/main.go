package main

import (
	"GoMagentoAPIs/auth"
	"GoMagentoAPIs/apis"
	"log"
	"net/http"
)

func main() {
	authToken, err := auth.GetAuthToken()
	if err != nil {
		log.Fatalf("Error getting auth token: %v", err)
	}

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		apis.GetAllProducts(w, r, authToken)
	})
	http.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		apis.GetAllCustomers(w, r, authToken)
	})
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		apis.GetAllOrders(w, r, authToken)
	})
	http.HandleFunc("/payments", func(w http.ResponseWriter, r *http.Request) {
		apis.GetAllPayments(w, r, authToken)
	})
	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		apis.GetProductById(w, r, authToken)
	})
	http.HandleFunc("/customer", func(w http.ResponseWriter, r *http.Request) {
		apis.GetCustomerById(w, r, authToken)
	})
	http.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		apis.GetOrderById(w, r, authToken)
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}