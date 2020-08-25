package handlers

import (
	"html/template"
	"net/http"
)

// VictorianSprayHandler TODO: Document
func VictorianSprayHandler(w http.ResponseWriter, r *http.Request) {

	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/victorianspray.html",
	))
	handle(w, r, page)
}
