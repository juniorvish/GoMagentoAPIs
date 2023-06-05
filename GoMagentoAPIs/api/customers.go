package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/juniorvish/GoMagentoAPIs/auth"
	"github.com/juniorvish/GoMagentoAPIs/models"
	"github.com/juniorvish/GoMagentoAPIs/utils"
)

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("AuthToken")
	if !auth.ValidateToken(authToken) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	queryParams := r.URL.Query()
	filters := utils.ParseFilters(queryParams)
	pagination := utils.ParsePagination(queryParams)

	customers, err := models.GetAllCustomers(filters, pagination)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("AuthToken")
	if !auth.ValidateToken(authToken) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	customerID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	customer, err := models.GetCustomerByID(customerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}