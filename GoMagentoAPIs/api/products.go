package api

import (
	"encoding/json"
	"net/http"

	"github.com/juniorvish/GoMagentoAPIs/auth"
	"github.com/juniorvish/GoMagentoAPIs/models"
	"github.com/juniorvish/GoMagentoAPIs/utils"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("AuthToken")
	if !auth.ValidateToken(authToken) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	filters := utils.ParseFilters(r)
	pagination := utils.ParsePagination(r)

	products, err := models.GetProducts(filters, pagination)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("AuthToken")
	if !auth.ValidateToken(authToken) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	productID := r.URL.Query().Get("id")
	if productID == "" {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	product, err := models.GetProductByID(productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}