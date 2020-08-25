package handlers

import (
	"html/template"
	"net/http"
)

// VictorianWreathHandler TODO: Document
func VictorianWreathHandler(w http.ResponseWriter, r *http.Request) {

	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/victorianwreath.html",
	))
	handle(w, r, page)
}
