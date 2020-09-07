package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"transactions/api/handlers"
	"transactions/shared/responses"
	"transactions/storage"

	"github.com/dgraph-io/dgo/v2"
	"github.com/go-chi/render"

	"github.com/go-chi/chi"
)

// TransactionRs provides a collection of resources for the transactions API.
type TransactionRs struct {
	Db *dgo.Dgraph
}

type emptyResponse struct{}

type dateType struct {
	Date string `json:"date,omitempty"`
}

const dateFormat = "2006-01-02"

// Routes creates a REST router for the api resources
func (rs TransactionRs) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/load", rs.loadData)
	router.Get("/load/{date}", rs.loadData)
	router.Get("/customers", rs.getCustomers)
	router.Get("/customers/{id}", rs.getCustomer)

	return router
}

func (rs *TransactionRs) loadData(w http.ResponseWriter, r *http.Request) {
	dateParam := chi.URLParam(r, "date")

	if dateParam == "" {
		dateParam = time.Now().Format(dateFormat)
	}

	date, err := time.Parse(dateFormat, dateParam)

	if err != nil {
		render.Render(w, r, responses.NewErrResponse(400, err))
		return
	}

	data, err := handlers.Load(date)

	if err != nil {
		render.Render(w, r, responses.NewErrResponse(500, err))
	}

	buyersJSON, err := json.Marshal(data.Buyers)
	if err != nil {
		log.Println(err)
	}

	if err := storage.Save(rs.Db, buyersJSON); err != nil {
		log.Println(err)
	}

	productsJSON, err := json.Marshal(data.Products)
	if err != nil {
		log.Println(err)
	}
	if err := storage.Save(rs.Db, productsJSON); err != nil {
		log.Println(err)
	}

	transactionsJSON, err := json.Marshal(data.Transactions)
	if err != nil {
		log.Println(err)
	}
	if err := storage.Save(rs.Db, transactionsJSON); err != nil {
		log.Println(err)
	}

	if err != nil {
		render.Render(w, r, responses.NewErrResponse(400, err))
		return
	}

	var buyerArr []string

	for _, buyer := range data.Buyers {
		buyerArr = append(buyerArr, buyer.ID)
	}

	if err := storage.BulkConnect(rs.Db, "id", "buyerID", "transaction",
		buyerArr); err != nil {
		render.Render(w, r, responses.NewErrResponse(400, err))
	}

	var productArr []string

	for _, product := range data.Products {
		productArr = append(productArr, product.ID)
	}

	if err := storage.BulkConnect(rs.Db, "productIDs", "id", "product",
		productArr); err != nil {
		render.Render(w, r, responses.NewErrResponse(400, err))
	}

	// Save date so can't be loaded twice
	dateJSON, err := json.Marshal(dateType{Date: dateParam})
	if err != nil {
		render.Render(w, r, responses.NewErrResponse(400, err))
	}

	if err := storage.Save(rs.Db, dateJSON); err != nil {
		log.Println(err)
	}

	render.JSON(w, r, dateJSON)
}

func (rs TransactionRs) getCustomers(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, emptyResponse{})
}

func (rs TransactionRs) getCustomer(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, emptyResponse{})
}
