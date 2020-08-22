package routes

import (
	"errors"
	"net/http"
	"transactions/shared/responses"

	"github.com/go-chi/render"

	"github.com/go-chi/chi"
	// "transactions/shared/responses"
)

// TransactionRs provides a collection of resources for the transactions API.
type TransactionRs struct{}

type emptyResponse struct{}

// Routes creates a REST router for the api resources
func (rs TransactionRs) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/load", rs.loadData)
	router.Get("/customers", rs.getCustomers)
	router.Get("/customers/{id}", rs.getCustomer)

	return router
}

func (rs TransactionRs) loadData(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, responses.NewErrorResponse(500, errors.New("test error")))
}

func (rs TransactionRs) getCustomers(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, emptyResponse{})
}

func (rs TransactionRs) getCustomer(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, emptyResponse{})
}
