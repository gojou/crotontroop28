package main

import (
	"net/http"

	"github.com/gojou/crotontroop28/pkg/handlers"
	"github.com/gorilla/mux"
)

func routes(r *mux.Router) {
	r.HandleFunc("/cancel", handlers.CancelHandler)
	r.HandleFunc("/thanks", handlers.ThanksHandler)
	r.HandleFunc("/dues", handlers.DuesHandler)
	r.HandleFunc("/products/classicwreath", handlers.ClassicWreathHandler)
	r.HandleFunc("/products/victorianwreath", handlers.VictorianWreathHandler)
	r.HandleFunc("/products/classicspray", handlers.ClassicSprayHandler)
	r.HandleFunc("/products/victorianspray", handlers.VictorianSprayHandler)
	r.HandleFunc("/products", handlers.ProductsHandler)
	r.HandleFunc("/", handlers.IndexHandler)
	r.HandleFunc("/chili", handlers.ChiliHandler)
	r.HandleFunc("/chili/rules", handlers.RulesHandler)
	r.HandleFunc("/dues", handlers.DuesHandler)

	r.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)

}
