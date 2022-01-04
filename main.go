package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/lucas3034/api_pokedex/functions"

	//"io/ioutil"
	"net/http"
)

func main() {
	rotas()
	fmt.Println("Server na porta 7000")
	http.ListenAndServe(":7000", nil)
}

type Pokemon struct {
	Id   int
	Nome string
	Tipo string
}

var Pokemons []Pokemon = []Pokemon{
	{
		Id:   1,
		Nome: "Bulbassauro",
		Tipo: "Grama/Veneno",
	},
	{
		Id:   2,
		Nome: "Ivyssauro",
		Tipo: "Grama/Veneno",
	},
	{
		Id:   3,
		Nome: "Venossauro",
		Tipo: "Grama/Veneno",
	},
	{
		Id:   4,
		Nome: "Charmander",
		Tipo: "Fogo",
	},
	{
		Id:   5,
		Nome: "Charmeleon",
		Tipo: "Fogo",
	},
	{
		Id:   6,
		Nome: "Charizard",
		Tipo: "Fogo/Voador",
	},
	{
		Id:   7,
		Nome: "Squirtle",
		Tipo: "Água",
	},
	{
		Id:   8,
		Nome: "Wartortle",
		Tipo: "Água",
	},
	{
		Id:   9,
		Nome: "Blastoise",
		Tipo: "Água",
	},
}

func inicio(w http.ResponseWriter, r *http.Request) {
	functions.Mensagem(w, r)
}

func mostrarPokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(Pokemons)
}

 func cadastrarPokemons(w http.ResponseWriter, r *http.Request) {
 	w.Header().Set("Content-Type", "application/json")
 	w.WriteHeader(http.StatusCreated)

 	body, _ := ioutil.ReadAll(r.Body)
 	var novoPokemon Pokemon
 	json.Unmarshal(body, &novoPokemon)
 	novoPokemon.Id = len(Pokemons) + 1
 	novoPokemon.Nome = "Novo"
 	novoPokemon.Tipo = "Tipo"
 	Pokemons = append(Pokemons, novoPokemon)

 	encoder := json.NewEncoder(w)
 	encoder.Encode(Pokemons)
 }

func rotas() {
	http.HandleFunc("/", inicio)
	http.HandleFunc("/pokemons", rotasPokemons)
}

func rotasPokemons(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		mostrarPokemons(w, r)
	} else if r.Method == "POST" {
		cadastrarPokemons(w, r)
	}
}
