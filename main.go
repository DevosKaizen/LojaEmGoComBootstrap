package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", index)

	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	// vamos mecher direto no html por aqui

	produtos := []Produto{
		{"iphone", "usado", 700, 1},
		{"tijolo", "novo", 0.50, 1000},
		{"combo de sushi", "delicioso", 50.55, 10},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
	

}
