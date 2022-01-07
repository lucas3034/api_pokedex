package main

import (
	"fmt"
	"net/http"

	"github.com/lucas3034/api_pokedex/functions"
)

func main() {
	routes()
	fmt.Println("Server on port: 7000")
	http.ListenAndServe(":7000", nil)
}

func Start(w http.ResponseWriter, r *http.Request) {
	functions.Message(w)
}

func routes() {
	http.HandleFunc("/", Start)
	http.HandleFunc("/pokemons", RoutesPokemons)
	http.HandleFunc("/pokemons/", ListsPokemons)
}

func RoutesPokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMultipleChoices)
	if r.Method == "GET" {
		functions.ShowPokemons(w)
	} else if r.Method == "POST" {
		functions.RegisterPokemons(w, r)
	} else {
		fmt.Fprintf(w, "The method entered is incorrect")
	}
}

func ListsPokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMultipleChoices)
	if r.Method == "GET" {
		functions.SearchPokemons(w, r)
	} else if r.Method == "DELETE" {
		functions.DeletePokemons(w, r)
	} else if r.Method == "PUT" {
		functions.EditPokemons(w, r)
	} else {
		fmt.Fprintf(w, "The method entered is incorrect")
	}
}
