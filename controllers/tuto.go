package controllers

import (
	"net/http"
	"html/template"
)

type TutoControlleur struct {
	tmpl *template.Template
}

func NewTutoControlleur() *TutoControlleur{
	return &TutoControlleur{
		tmpl: template.Must(template.ParseFiles("tuto.html")),
	}
}

func (c *TutoControlleur) ServeHTTP(w http.ResponseWriter, r *http.Request){
	var data struct{}
	c.tmpl.Execute(w,data)
}
