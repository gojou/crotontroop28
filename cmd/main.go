package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gojou/crotontroop28/pkg/handlers"
)

func main() {
	http.HandleFunc("/products/classicwreath", handlers.ClassicWreathHandler)
	http.HandleFunc("/products/victorianwreath", handlers.VictorianWreathHandler)
	http.HandleFunc("/products/classicspray", handlers.ClassicSprayHandler)
	http.HandleFunc("/products/victorianspray", handlers.VictorianSprayHandler)
	http.HandleFunc("/products", handlers.ProductsHandler)
	http.HandleFunc("/", handlers.IndexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}
