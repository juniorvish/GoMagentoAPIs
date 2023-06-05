package main

import (
	"GoMagentoAPIs/auth"
	"GoMagentoAPIs/api"
	"log"
	"net/http"
)

func main() {
	authMiddleware := auth.AuthMiddleware()

	http.HandleFunc("/api/products", authMiddleware(api.GetAllProducts))
	http.HandleFunc("/api/customers", authMiddleware(api.GetAllCustomers))
	http.HandleFunc("/api/orders", authMiddleware(api.GetAllOrders))
	http.HandleFunc("/api/payments", authMiddleware(api.GetAllPayments))
	http.HandleFunc("/api/product/", authMiddleware(api.GetProductByID))
	http.HandleFunc("/api/customer/", authMiddleware(api.GetCustomerByID))
	http.HandleFunc("/api/order/", authMiddleware(api.GetOrderByID))

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}