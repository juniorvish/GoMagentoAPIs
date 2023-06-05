package apis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"GoMagentoAPIs/auth"
)

type Payment struct {
	OrderID       string `json:"order_id"`
	PaymentMethod string `json:"payment_method"`
	AmountPaid    string `json:"amount_paid"`
}

type PaymentsResponse struct {
	Items []Payment `json:"items"`
}

func getAllPayments(authToken string, filters map[string]string, page int, pageSize int) (PaymentsResponse, error) {
	magentoBaseURL := "https://your-magento-instance.com/rest/V1/"
	client := &http.Client{}
	req, err := http.NewRequest("GET", magentoBaseURL+"payments", nil)
	if err != nil {
		return PaymentsResponse{}, err
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
		return PaymentsResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return PaymentsResponse{}, err
	}

	var paymentsResponse PaymentsResponse
	err = json.Unmarshal(body, &paymentsResponse)
	if err != nil {
		return PaymentsResponse{}, err
	}

	return paymentsResponse, nil
}