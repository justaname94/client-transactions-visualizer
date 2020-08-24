package handlers

import (
	"net/http"
	"os"
	"time"
)

var (
	baseURL = os.Getenv("DATA_URL")
	client  = http.Client{
		Timeout: time.Second * 2,
	}
)

// LoadBuyers TODO
func LoadBuyers() {
}
