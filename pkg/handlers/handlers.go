package handlers

import (
	"html/template"
	"net/http"
)

//TemplateParams contain parameters
type TemplateParams struct {
	p1 string
}

var (
	page   *template.Template
	params TemplateParams
)

//IndexHandler handles the default page
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	// Set the default page

	page = template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/index.html",
	))

	//Display 404 if it's odball URL
	if r.URL.Path != "/" {
		page = template.Must(template.ParseFiles(
			"static/html/_base.html",
			"static/html/404.html",
		))
	}
	handle(w, r, page)
}

// ProductsHandler TODO: Document
func ProductsHandler(w http.ResponseWriter, r *http.Request) {

	page = template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/products.html",
	))
	handle(w, r, page)
}

// RulesHandler TODO: Document
func RulesHandler(w http.ResponseWriter, r *http.Request) {

	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/rules.html",
	))
	handle(w, r, page)
}

// ClassicWreathHandler TODO: Document
func ClassicWreathHandler(w http.ResponseWriter, r *http.Request) {

	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/classicwreath.html",
	))
	handle(w, r, page)
}

// VictorianWreathHandler TODO: Document
func VictorianWreathHandler(w http.ResponseWriter, r *http.Request) {

	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/victorianwreath.html",
	))
	handle(w, r, page)
}

// ClassicSprayHandler TODO: Document
func ClassicSprayHandler(w http.ResponseWriter, r *http.Request) {

	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/classicspray.html",
	))
	handle(w, r, page)
}

// VictorianSprayHandler TODO: Document
func VictorianSprayHandler(w http.ResponseWriter, r *http.Request) {

	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/victorianspray.html",
	))
	handle(w, r, page)
}

func handle(w http.ResponseWriter, r *http.Request, p *template.Template) {
	if r.Method == "GET" {
		p.Execute(w, params)
	}
	return
}
