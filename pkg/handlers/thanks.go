package handlers

import (
	"html/template"
	"net/http"
)

// ThanksHandler TODO: Document
func ThanksHandler(w http.ResponseWriter, r *http.Request) {

	page = template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/thanks.html",
	))
	handle(w, r, page)
}
