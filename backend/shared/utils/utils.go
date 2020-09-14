package utils

import (
	"net/http"
	"time"
	buyer "transactions/shared/models/buyer"
	product "transactions/shared/models/product"
	transaction "transactions/shared/models/transaction"
)

// Entities represents an array of the used models
type Entities struct {
	Buyers       []*buyer.Buyer
	Products     []*product.Product
	Transactions []*transaction.Transaction
}

// RetryRequest recursively uses exponential backoff to retry a request a set amount
// of times. Inspired from https://upgear.io/blog/simple-golang-retry-function/
func RetryRequest(client http.Client, attempts int, sleep time.Duration,
	req *http.Request) (*http.Response, error) {
	response, err := client.Do(req)

	if err != nil {
		if attempts--; attempts > 0 {
			time.Sleep(sleep)
			return RetryRequest(client, attempts, 2*sleep, req)
		}
		return &http.Response{}, err
	}
	return response, nil
}
