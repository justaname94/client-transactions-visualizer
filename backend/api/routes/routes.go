package routes

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"
	"transactions/api/handlers"
	messages "transactions/shared/error-messages"
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
	router.Get("/buyer", rs.getBuyers)
	router.Get("/buyer/{id}", rs.getBuyer)

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

	// Cannot load the same date twice
	dateDB, err := storage.Query(rs.Db, storage.GetDate, map[string]string{
		"$date": dateParam,
	})
	if err != nil {
		render.Render(w, r, responses.NewErrResponse(500, err))
		return
	}
	if dateDB.GetMetrics().GetNumUids()["_total"] > 0 {
		render.Render(w, r, responses.NewErrResponse(403,
			errors.New("cannot load same date twice")))
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
		render.Render(w, r, responses.NewErrResponse(500, err))
	}

	if err := storage.Save(rs.Db, dateJSON); err != nil {
		log.Println(err)
	}

	render.JSON(w, r, map[string]bool{
		"success": true,
	})
}

func (rs TransactionRs) getBuyers(w http.ResponseWriter, r *http.Request) {
	limit := 10
	page := 0

	limitParam := r.URL.Query().Get("limit")
	pageParam := r.URL.Query().Get("page")

	if limitParam != "" {
		if intVal, err := strconv.Atoi(limitParam); err == nil {
			limit = intVal
		}
	}

	if pageParam != "" {
		if intVal, err := strconv.Atoi(pageParam); err == nil {
			page = (page - 1 + intVal) * limit
		}
	}

	buyers, err := storage.Query(rs.Db, storage.AllBuyersPaginated,
		map[string]string{
			"$limit": strconv.Itoa(limit),
			"$page":  strconv.Itoa(page),
		})

	if err != nil {
		render.Render(w, r, responses.NewErrResponse(500, err))
		return
	}

	var res map[string]*json.RawMessage
	if err := json.Unmarshal(buyers.Json, &res); err != nil {
		render.Render(w, r, responses.NewErrResponse(500, err))
		return
	}

	render.JSON(w, r, res)
}

func (rs *TransactionRs) getBuyer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	info, err := storage.Query(rs.Db, storage.BuyerInfo, map[string]string{
		"$id": id,
	})
	if err != nil {
		render.Render(w, r, responses.NewErrResponse(500, err))
		return
	}

	// Buyer was not found
	if info.GetMetrics().GetNumUids()["_total"] == 0 {
		render.Render(w, r, responses.NewErrResponse(404,
			messages.ErrNotFound("buyer")))
		return
	}

	var res map[string]*json.RawMessage
	if err := json.Unmarshal(info.Json, &res); err != nil {
		render.Render(w, r, responses.NewErrResponse(500, err))
		return
	}

	render.JSON(w, r, res)
}
