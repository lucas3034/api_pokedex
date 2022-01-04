package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

func MostrarPokemons(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(Pokemons)
}

func CadastrarPokemons(w http.ResponseWriter, r *http.Request) {
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

//func BuscarPokemons(w http.ResponseWriter, r *http.Request){
//	//w.Header().Set("Content-Type", "application/json")
//	fmt.Println(r.URL.Path)
//	valor := strings.Split(r.URL.Path, "/")
//
//	if len(valor) > 9 {
//		w.WriteHeader(http.StatusNotFound)
//		return
//	}
//
//	id, _ := strconv.Atoi(valor[2])
//
//	for _, Pokemon := range Pokemons {
//		if Pokemons.Id == id {
//			json.NewEncoder(w).Encode(Pokemon)
//			return
//		}
//	}
//
//
//}

func Mensagem(w http.ResponseWriter) {
	fmt.Fprintf(w, "Bem-vindos a Pokedéx")
}
