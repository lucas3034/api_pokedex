package routes

import (
	"fmt"
	"net/http"

	"github.com/lucas3034/api_pokedex/functions"
)

func RoutesPokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMultipleChoices)
	if r.Method == "GET" {
		functions.ShowPokemons(w)
	} else if r.Method == "POST" {
		functions.RegisterPokemons(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
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
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "The method entered is incorrect")
	}
}
