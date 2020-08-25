package handlers

import (
	"html/template"
	"net/http"
)

//ChiliHandler handles the default page
func ChiliHandler(w http.ResponseWriter, r *http.Request) {

	// Set the default page

	page = template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/chili.html",
	))
	handle(w, r, page)
}
