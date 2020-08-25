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

func handle(w http.ResponseWriter, r *http.Request, p *template.Template) {
	if r.Method == "GET" {
		p.Execute(w, params)
	}
	return
}
