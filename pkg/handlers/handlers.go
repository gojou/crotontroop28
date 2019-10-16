package handlers

import (
	"html/template"
	"net/http"
)
type TemplateParams struct {
	p1 string
}

// func IndexHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		http.NotFound(w, r)
// 		return
// 	}
// 	fmt.Fprint(w, "Hello, World!")
// }

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

func RulesHandler(w http.ResponseWriter, r *http.Request) {

	page := template.Must(template.ParseFiles(
		"static/html/_base.html",
		"static/html/rules.html",
	))

	if r.Method == "GET" {
		page.Execute(w, nil)
		return
	}
}
