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

func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("AuthToken")
	if !auth.ValidateToken(authToken) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	filters := utils.ParseFilters(r)
	pagination := utils.ParsePagination(r)

	orders, err := models.GetAllOrders(filters, pagination)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func GetOrderByID(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("AuthToken")
	if !auth.ValidateToken(authToken) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	order, err := models.GetOrderByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}