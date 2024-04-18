package controllers

import (
	"net/http"
	"html/template"
)

type IndexControlleur struct {
	tmpl *template.Template
}

func NewIndexControlleur() *IndexControlleur{
	return &IndexControlleur{
		tmpl: template.Must(template.ParseFiles("index.html")),
	}
}

func (c *IndexControlleur) ServeHTTP(w http.ResponseWriter, r *http.Request){
	var data struct{}
	c.tmpl.Execute(w,data)
}
