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
	http.HandleFunc("/pokemons/", functions.BuscarPokemons)
}

func rotasPokemons(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		functions.MostrarPokemons(w)
	} else if r.Method == "POST" {
		functions.CadastrarPokemons(w, r)
	}
}
