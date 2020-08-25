package handlers

import (
	"html/template"
	"net/http"
)

//IndexHandler handles the default page
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	// Set the default page

	page = template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/index.html",
	))
	handle(w, r, page)
}
