package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	servidor()
}

type Pokemon struct {
	Id    int
	Nome  string
	Tipo1 string
	Tipo2 string
}

var Pokemons []Pokemon = []Pokemon{
	Pokemon{
		Id:    1,
		Nome:  "Bulbassauro",
		Tipo1: "Grama",
		Tipo2: "Veneno",
	},
	Pokemon{
		Id:    2,
		Nome:  "Ivyssauro",
		Tipo1: "Grama",
		Tipo2: "Veneno",
	},
	Pokemon{
		Id:    3,
		Nome:  "Venossauro",
		Tipo1: "Grama",
		Tipo2: "Veneno",
	},
	Pokemon{
		Id:    4,
		Nome:  "Charmander",
		Tipo1: "Fogo",
		Tipo2: "",
	},
	Pokemon{
		Id:    5,
		Nome:  "Charmeleon",
		Tipo1: "Fogo",
		Tipo2: "",
	},
	Pokemon{
		Id:    6,
		Nome:  "Charizard",
		Tipo1: "Fogo",
		Tipo2: "Voador",
	},
	Pokemon{
		Id:    7,
		Nome:  "Squirtle",
		Tipo1: "Água",
		Tipo2: "",
	},
	Pokemon{
		Id:    8,
		Nome:  "Wartortle",
		Tipo1: "Água",
		Tipo2: "",
	},
	Pokemon{
		Id:    9,
		Nome:  "Blastoise",
		Tipo1: "Água",
		Tipo2: "",
	},
}

func inicio(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem-vindos a Pokedéx")
}

func Mostrar(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(Pokemons)
}

func rotas() {
	http.HandleFunc("/", inicio)
	http.HandleFunc("/pokemons", Mostrar)
}

func servidor() {
	rotas()

	fmt.Println("Server na porta 7000")
	http.ListenAndServe(":7000", nil)
}
