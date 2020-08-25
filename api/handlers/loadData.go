package handlers

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	buyer "transactions/shared/models/buyer"
	product "transactions/shared/models/product"
)

var (
	client = http.Client{
		Timeout: time.Second * 2,
	}
)

func getURL(endpoint string, unixTime int64) (string, error) {
	url := os.Getenv("DATA_URL")

	return fmt.Sprintf("%s/%s?date=%d", url, endpoint, unixTime), nil
}

// LoadBuyers TODO
func LoadBuyers(date time.Time) ([]*buyer.Buyer, error) {
	url, _ := getURL("buyers", date.Unix())

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var buyersInterface interface{}

	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&buyersInterface); err != nil {
		return nil, err
	}

	var buyers []*buyer.Buyer

	switch v := buyersInterface.(type) {
	case []interface{}:
		for _, data := range v {
			id := data.(map[string]interface{})["id"].(string)
			name := data.(map[string]interface{})["name"].(string)
			age := int(data.(map[string]interface{})["age"].(float64))

			newBuyer, err := buyer.NewBuyer(id, name, age)
			if err != nil {
				// Log and ignore incomplete buyers
				log.Println(err)
			}
			buyers = append(buyers, newBuyer)
		}
	default:
		return nil, errors.New("an error ocurred obtaining data from the endpoint")
	}

	return buyers, nil
}

// LoadProducts TODO
func LoadProducts(date time.Time) ([]*product.Product, error) {
	url, _ := getURL("products", date.Unix())

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	reader := csv.NewReader(res.Body)
	reader.Comma = '\''

	if err != nil {
		return nil, err
	}

	var products []*product.Product

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			// Log error and ignore product record
			log.Println(err)
		}

		id := string(record[0])
		name := string(record[1])

		price, _ := strconv.Atoi(record[2])

		product, err := product.NewProduct(id, name, price)

		if err != nil {
			// Log error and ignore product record
			log.Println(err)
		}

		products = append(products, product)
	}

	return products, nil
}
