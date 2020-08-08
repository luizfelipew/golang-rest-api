package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul, top de linha", Preco: 39, Quantidade: 2},
		{"Tenis", "Confortavel", 89, 3},
		{"Fone", "Top de linha", 59, 2},
		{"Produto Novo", "Top mesmo", 1.99, 1},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
