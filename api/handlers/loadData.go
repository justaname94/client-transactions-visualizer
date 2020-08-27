package handlers

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	buyer "transactions/shared/models/buyer"
	product "transactions/shared/models/product"
	transaction "transactions/shared/models/transaction"
)

var (
	client = http.Client{
		Timeout: time.Second * 2,
	}
)

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

// LoadTransactions Todo
func LoadTransactions(date time.Time) ([]*transaction.Transaction, error) {
	url, _ := getURL("transactions", date.Unix())

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	reader, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	readerStr := string(reader)

	transactions, err := parseCustomFormat(readerStr)

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func getURL(endpoint string, unixTime int64) (string, error) {
	url := os.Getenv("DATA_URL")

	return fmt.Sprintf("%s/%s?date=%d", url, endpoint, unixTime), nil
}

// parseCustomFormat splits # as new line, '\x00' for inside characters
func parseCustomFormat(data string) ([]*transaction.Transaction, error) {
	items := strings.Split(data, "#")

	var transactions []*transaction.Transaction

	for idx, item := range items {
		// Empty item
		if idx == 0 {
			continue
		}
		fields := strings.Split(item, "\x00")

		if len(fields) < 4 {
			// Log and ignore transaction error
			log.Printf("invalid transaction: %s\n", fields)
			continue
		}

		id := fields[0]
		buyerID := fields[1]
		ip := fields[2]
		device := fields[3]
		// Remove parentheses before splitting
		productIds := strings.Split(fields[4][1:len(fields[4])-1], ",")

		transaction, err := transaction.NewTransaction(id,
			buyerID, ip, device, productIds)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
