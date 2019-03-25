package main

import (
	//	"fmt"
	"html/template"
	"net/http"
	//	"time"
	//	"strconv"

	// "google.golang.org/appengine/datastore"
	// "google.golang.org/appengine/log"

	"google.golang.org/appengine"
)

type TemplateParams struct {
	p1    string

}

func main() {
	http.HandleFunc("/", indexHandler)
	appengine.Main() // Starts the server to receive requests
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	params := TemplateParams{}

	// Set the default page

	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/index.html",
	))

	//Display 404 if it's odball URL
	if r.URL.Path != "/" {
		page = template.Must(template.ParseFiles(
			"static/html/_base.html",
			"static/html/404.html",
		))

	}

	if r.Method == "GET" {
		page.Execute(w, params)
		//	return
	}

}
