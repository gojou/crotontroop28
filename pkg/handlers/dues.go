package handlers

import (
	"html/template"
	"net/http"
)

// DuesHandler TODO: Document
func DuesHandler(w http.ResponseWriter, r *http.Request) {

	page = template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/dues.html",
	))
	handle(w, r, page)
}
