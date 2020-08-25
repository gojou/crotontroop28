package handlers

import (
	"html/template"
	"net/http"
)

// ClassicWreathHandler TODO: Document
func ClassicWreathHandler(w http.ResponseWriter, r *http.Request) {

	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/classicwreath.html",
	))
	handle(w, r, page)
}
