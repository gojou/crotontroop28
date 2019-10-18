package handlers

import (
	"html/template"
	"net/http"
)

type TemplateParams struct {
	p1 string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	params := TemplateParams{}

	// Set the default page

	page := template.Must(template.ParseFiles(
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

	if r.Method == "GET" {
		page.Execute(w, params)
		//	return
	}

}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {

	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/products.html",
	))

	if r.Method == "GET" {
		page.Execute(w, nil)
		return
	}
}

func ClassicSprayHandler(w http.ResponseWriter, r *http.Request) {

	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/classicspray.html",
	))

	if r.Method == "GET" {
		page.Execute(w, nil)
		return
	}
}

func ClassicWreathHandler(w http.ResponseWriter, r *http.Request) {

	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/classicwreath.html",
	))

	if r.Method == "GET" {
		page.Execute(w, nil)
		return
	}
}

func VictorianSprayHandler(w http.ResponseWriter, r *http.Request) {

	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/victorianspray.html",
	))

	if r.Method == "GET" {
		page.Execute(w, nil)
		return
	}
}

func VictorianWreathHandler(w http.ResponseWriter, r *http.Request) {

	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/victorianwreath.html",
	))

	if r.Method == "GET" {
		page.Execute(w, nil)
		return
	}
}
