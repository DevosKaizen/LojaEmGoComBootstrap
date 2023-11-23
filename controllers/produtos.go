package controllers

import (
	"LOJAEMGO/models"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)

}

func New(w http.ResponseWriter, r *http.Request) {

	// todosOsProdutos := models.BuscaTodosProdutos()
	// temp.ExecuteTemplate(w, "New", todosOsProdutos)

	temp.ExecuteTemplate(w, "New", nil)
}
