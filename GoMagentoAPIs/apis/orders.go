package apis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"GoMagentoAPIs/auth"
)

type Order struct {
	ID           int       `json:"id"`
	CreatedDate  time.Time `json:"created_date"`
	UpdatedDate  time.Time `json:"updated_date"`
	CustomerCode string    `json:"customer_code"`
	CustomerName string    `json:"customer_name"`
}

type ordersResponse struct {
	Orders      []Order `json:"orders"`
	TotalCount  int     `json:"total_count"`
	CurrentPage int     `json:"current_page"`
	PageSize    int     `json:"page_size"`
}

func getAllOrders(authToken string, filters map[string]string, page int, pageSize int) (ordersResponse, error) {
	magentoBaseURL := "https://example.com/magento/rest/V1"
	client := &http.Client{}
	req, err := http.NewRequest("GET", magentoBaseURL+"/orders", nil)
	if err != nil {
		return ordersResponse{}, err
	}

	req.Header.Add("Authorization", "Bearer "+authToken)
	req.Header.Add("Content-Type", "application/json")

	query := req.URL.Query()
	for key, value := range filters {
		query.Add(key, value)
	}
	query.Add("page", strconv.Itoa(page))
	query.Add("pageSize", strconv.Itoa(pageSize))
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return ordersResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ordersResponse{}, err
	}

	var ordersResp ordersResponse
	err = json.Unmarshal(body, &ordersResp)
	if err != nil {
		return ordersResponse{}, err
	}

	return ordersResp, nil
}

func getOrderById(authToken string, orderId int) (Order, error) {
	magentoBaseURL := "https://example.com/magento/rest/V1"
	client := &http.Client{}
	req, err := http.NewRequest("GET", magentoBaseURL+"/orders/"+strconv.Itoa(orderId), nil)
	if err != nil {
		return Order{}, err
	}

	req.Header.Add("Authorization", "Bearer "+authToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return Order{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Order{}, err
	}

	var order Order
	err = json.Unmarshal(body, &order)
	if err != nil {
		return Order{}, err
	}

	return order, nil
}