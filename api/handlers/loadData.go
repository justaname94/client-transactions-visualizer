package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	models "transactions/shared/models/buyer"
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
func LoadBuyers(date time.Time) ([]*models.Buyer, error) {
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

	if err := json.NewDecoder(res.Body).Decode(&buyersInterface); err != nil {
		return nil, err
	}

	var buyers []*models.Buyer

	switch v := buyersInterface.(type) {
	case []interface{}:
		for _, buyer := range v {
			id := buyer.(map[string]interface{})["id"].(string)
			name := buyer.(map[string]interface{})["name"].(string)
			age := int(buyer.(map[string]interface{})["age"].(float64))

			newBuyer, err := models.NewBuyer(id, name, age)
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
