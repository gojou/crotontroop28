package handlers

import (
	"html/template"
	"net/http"
)

// ProductsHandler TODO: Document
func ProductsHandler(w http.ResponseWriter, r *http.Request) {

	page = template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/products.html",
	))
	handle(w, r, page)
}
