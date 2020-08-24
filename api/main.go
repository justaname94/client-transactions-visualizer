package main

import (
	"flag"
	"log"
	"net/http"
	"transactions/api/routes"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
)

var port = flag.String("port", "8080",
	"Port in which the server is going to run")

func main() {
	flag.Parse()

	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST"},
	})

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.Compress(5),
		middleware.Recoverer,
		cors.Handler,
	)

	router.Mount("/", routes.TransactionRs{}.Routes())

	log.Fatal(http.ListenAndServe(":"+*port, router))
}
