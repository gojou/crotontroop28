package handlers

import (
	"html/template"
	"net/http"
)

// ClassicSprayHandler TODO: Document
func ClassicSprayHandler(w http.ResponseWriter, r *http.Request) {

	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/classicspray.html",
	))
	handle(w, r, page)
}
