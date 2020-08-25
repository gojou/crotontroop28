package handlers

import (
	"html/template"
	"net/http"
)

// NotFoundHandler TODO: Document
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	page = template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/404.html",
	))
	handle(w, r, page)
}
