package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	handlers "crotontroop28/pkg/handlers"
)


func main() {

	http.HandleFunc("/rules", handlers.RulesHandler)
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
