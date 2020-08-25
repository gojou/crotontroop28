package handlers

import (
	"html/template"
	"net/http"
)

// RulesHandler TODO: Document
func RulesHandler(w http.ResponseWriter, r *http.Request) {

	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/rules.html",
	))
	handle(w, r, page)
}
