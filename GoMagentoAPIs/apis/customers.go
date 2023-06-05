package apis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"GoMagentoAPIs/auth"
)

type Customer struct {
	ID           int       `json:"id"`
	CreatedDate  time.Time `json:"created_date"`
	UpdatedDate  time.Time `json:"updated_date"`
	CustomerName string    `json:"customer_name"`
	CustomerCode string    `json:"customer_code"`
}

type customersResponse struct {
	Items []Customer `json:"items"`
}

func getAllCustomers(authToken string, createdDate, updatedDate, customerName, customerCode string, page, pageSize int) ([]Customer, error) {
	magentoBaseURL := "https://example.com/magento/rest/V1/customers/search"

	queryParams := url.Values{}
	queryParams.Add("searchCriteria[currentPage]", strconv.Itoa(page))
	queryParams.Add("searchCriteria[pageSize]", strconv.Itoa(pageSize))

	if createdDate != "" {
		queryParams.Add("searchCriteria[filterGroups][0][filters][0][field]", "created_at")
		queryParams.Add("searchCriteria[filterGroups][0][filters][0][value]", createdDate)
		queryParams.Add("searchCriteria[filterGroups][0][filters][0][conditionType]", "eq")
	}

	if updatedDate != "" {
		queryParams.Add("searchCriteria[filterGroups][1][filters][0][field]", "updated_at")
		queryParams.Add("searchCriteria[filterGroups][1][filters][0][value]", updatedDate)
		queryParams.Add("searchCriteria[filterGroups][1][filters][0][conditionType]", "eq")
	}

	if customerName != "" {
		queryParams.Add("searchCriteria[filterGroups][2][filters][0][field]", "name")
		queryParams.Add("searchCriteria[filterGroups][2][filters][0][value]", customerName)
		queryParams.Add("searchCriteria[filterGroups][2][filters][0][conditionType]", "eq")
	}

	if customerCode != "" {
		queryParams.Add("searchCriteria[filterGroups][3][filters][0][field]", "code")
		queryParams.Add("searchCriteria[filterGroups][3][filters][0][value]", customerCode)
		queryParams.Add("searchCriteria[filterGroups][3][filters][0][conditionType]", "eq")
	}

	req, err := http.NewRequest("GET", magentoBaseURL+"?"+queryParams.Encode(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+authToken)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response customersResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Items, nil
}

func getCustomerById(authToken string, customerId int) (*Customer, error) {
	magentoBaseURL := "https://example.com/magento/rest/V1/customers/"

	req, err := http.NewRequest("GET", magentoBaseURL+strconv.Itoa(customerId), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+authToken)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var customer Customer
	err = json.Unmarshal(body, &customer)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}