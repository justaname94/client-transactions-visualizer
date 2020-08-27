package routes

import (
	"net/http"
	"time"
	"transactions/shared/responses"

	"github.com/go-chi/render"

	"github.com/go-chi/chi"
)

// TransactionRs provides a collection of resources for the transactions API.
type TransactionRs struct{}

type emptyResponse struct{}

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

func (rs TransactionRs) loadData(w http.ResponseWriter, r *http.Request) {
	dateParam := chi.URLParam(r, "date")

	if dateParam == "" {
		dateParam = time.Now().Format(dateFormat)
	}

	date, err := time.Parse(dateFormat, dateParam)

	if err != nil {
		render.Render(w, r, responses.NewErrResponse(400, err))
		return
	}

	// buyers, err := handlers.LoadBuyers(date)

	// products, err := handlers.LoadProducts(date)

	// transactions, err := handlers.LoadTransactions(date)

	if err != nil {
		render.Render(w, r, responses.NewErrResponse(400, err))
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"date": date,
	})
}

func (rs TransactionRs) getCustomers(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, emptyResponse{})
}

func (rs TransactionRs) getCustomer(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, emptyResponse{})
}
