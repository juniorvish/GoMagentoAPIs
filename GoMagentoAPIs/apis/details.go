package apis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/juniorvish/GoMagentoAPIs/auth"
)

const (
	productDetailsEndpoint = "/V1/products/%s"
	customerDetailsEndpoint = "/V1/customers/%d"
	orderDetailsEndpoint   = "/V1/orders/%d"
)

func getProductById(authToken string, productId string) (*Product, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", magentoBaseURL+fmt.Sprintf(productDetailsEndpoint, productId), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+authToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var product Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func getCustomerById(authToken string, customerId int) (*Customer, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", magentoBaseURL+fmt.Sprintf(customerDetailsEndpoint, customerId), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+authToken)
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

func getOrderById(authToken string, orderId int) (*Order, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", magentoBaseURL+fmt.Sprintf(orderDetailsEndpoint, orderId), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+authToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var order Order
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}