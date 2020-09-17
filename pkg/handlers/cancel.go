package handlers

import (
	"html/template"
	"net/http"
)

// CancelHandler TODO: Document
func CancelHandler(w http.ResponseWriter, r *http.Request) {

	page = template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/cancel.html",
	))
	handle(w, r, page)
}
