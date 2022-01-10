package main

import (
	"fmt"
	"net/http"

	"github.com/lucas3034/api_pokedex/functions"
	"github.com/lucas3034/api_pokedex/routes"
)

func main() {
	localRoutes()
	fmt.Println("Server on port: 7000")
	http.ListenAndServe(":7000", nil)
}

func Start(w http.ResponseWriter, r *http.Request) {
	functions.Message(w)
}

func localRoutes() {
	http.HandleFunc("/", Start)
	http.HandleFunc("/pokemons", routes.RoutesPokemons)
	http.HandleFunc("/pokemons/", routes.ListsPokemons)
}

