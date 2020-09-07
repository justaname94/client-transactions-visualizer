package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"transactions/api/handlers"
	"transactions/shared/responses"
	"transactions/storage"
	"transactions/storage/helpers"

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
		return
	}

	if err := helpers.SaveEntities(rs.Db, data); err != nil {
		render.Render(w, r, responses.NewErrResponse(500, err))
		return
	}

	if err := helpers.ConnectFields(rs.Db, data); err != nil {
		render.Render(w, r, responses.NewErrResponse(500, err))
		return
	}

	// Save date so can't be loaded twice
	dateJSON, err := json.Marshal(dateType{Date: dateParam})
	if err != nil {
		render.Render(w, r, responses.NewErrResponse(400, err))
	}

	if err := storage.Save(rs.Db, dateJSON); err != nil {
		log.Println(err)
	}

	render.JSON(w, r, map[string]string{
		"success": "true",
	})
}

func (rs TransactionRs) getCustomers(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, emptyResponse{})
}

func (rs TransactionRs) getCustomer(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, emptyResponse{})
}
