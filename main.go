package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	db := conectaCombancoDeDados()
	defer db.Close()
	
	http.HandleFunc("/", index)

	http.ListenAndServe(":8000", nil)

}

	// Cria conexao com banco de dados e testa se tem erro
func conectaCombancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=admin host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
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
