package main

import (
	"fmt"
	"github.com/lucas3034/api_pokedex/functions"
	"net/http"
)

func main() {
	rotas()
	fmt.Println("Server na porta 7000")
	http.ListenAndServe(":7000", nil)
}

func inicio(w http.ResponseWriter, r *http.Request) {
	functions.Mensagem(w)
}

func rotas() {
	http.HandleFunc("/", inicio)
	http.HandleFunc("/pokemons", rotasPokemons)
	http.HandleFunc("/pokemons/", ListasPokemons)
}

func rotasPokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMultipleChoices)
	if r.Method == "GET" {
		functions.MostrarPokemons(w)
	} else if r.Method == "POST" {
		functions.CadastrarPokemons(w, r)
	} else {
		fmt.Fprintf(w, "O Método informado está incorreto")
	}
}

func ListasPokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMultipleChoices)
	if r.Method == "GET" {
		functions.BuscarPokemons(w, r)
	} else if r.Method == "DELETE" {
		functions.DeletarPokemons(w, r)
	} else {
		fmt.Fprintf(w, "O Método informado está incorreto")
	}
}