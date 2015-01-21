package web

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("templates/layout.html"))

func renderTemplate(w http.ResponseWriter, path string, p *Page) {
	clone, err := templates.Clone()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	t, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	t, err = clone.AddParseTree("main", t.Tree)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = t.ExecuteTemplate(w, "layout.html", p)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
