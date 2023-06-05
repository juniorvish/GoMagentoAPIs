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

type Product struct {
	ID          int       `json:"id"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
}

type ProductsResponse struct {
	Products []Product `json:"products"`
}

func getAllProducts(authToken string, createdDate, updatedDate, productName, productCode string, page, pageSize int) (ProductsResponse, error) {
	magentoBaseURL := "https://example.com/magento/api"
	client := &http.Client{}
	req, err := http.NewRequest("GET", magentoBaseURL+"/products", nil)
	if err != nil {
		return ProductsResponse{}, err
	}

	q := url.Values{}
	q.Add("page", strconv.Itoa(page))
	q.Add("pageSize", strconv.Itoa(pageSize))

	if createdDate != "" {
		q.Add("createdDate", createdDate)
	}
	if updatedDate != "" {
		q.Add("updatedDate", updatedDate)
	}
	if productName != "" {
		q.Add("productName", productName)
	}
	if productCode != "" {
		q.Add("productCode", productCode)
	}

	req.URL.RawQuery = q.Encode()
	req.Header.Add("Authorization", "Bearer "+authToken)

	resp, err := client.Do(req)
	if err != nil {
		return ProductsResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ProductsResponse{}, err
	}

	var productsResponse ProductsResponse
	err = json.Unmarshal(body, &productsResponse)
	if err != nil {
		return ProductsResponse{}, err
	}

	return productsResponse, nil
}

func getProductById(authToken string, id int) (Product, error) {
	magentoBaseURL := "https://example.com/magento/api"
	client := &http.Client{}
	req, err := http.NewRequest("GET", magentoBaseURL+"/products/"+strconv.Itoa(id), nil)
	if err != nil {
		return Product{}, err
	}

	req.Header.Add("Authorization", "Bearer "+authToken)

	resp, err := client.Do(req)
	if err != nil {
		return Product{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Product{}, err
	}

	var product Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}