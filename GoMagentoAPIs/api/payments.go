package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/juniorvish/GoMagentoAPIs/auth"
	"github.com/juniorvish/GoMagentoAPIs/models"
	"github.com/juniorvish/GoMagentoAPIs/utils"
)

func GetAllPayments(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("AuthToken")
	if !auth.ValidateToken(authToken) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	queryParams := r.URL.Query()
	filters := utils.ParseFilters(queryParams)
	pagination := utils.ParsePagination(queryParams)

	payments, err := models.GetAllPayments(filters, pagination)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payments)
}

func GetPaymentByID(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("AuthToken")
	if !auth.ValidateToken(authToken) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	payment, err := models.GetPaymentByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payment)
}