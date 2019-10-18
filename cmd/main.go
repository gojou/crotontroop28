package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"crotontroop28/pkg/handlers"
)


func main() {
	http.HandleFunc("/classicwreath", handlers.ClassicWreathHandler)
	http.HandleFunc("/victorianwreath", handlers.VictorianWreathHandler)
	http.HandleFunc("/classicspray", handlers.ClassicSprayHandler)
	http.HandleFunc("/victorianspray", handlers.VictorianSprayHandler)
	http.HandleFunc("/products", handlers.ProductsHandler)
	http.HandleFunc("/", handlers.IndexHandler)


	port := os.Getenv("PORT")
if port == "" {
	port = "8080"
	log.Printf("Defaulting to port %s", port)
}

log.Printf("Listening on port %s", port)
log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
// [END setting_port]

}
