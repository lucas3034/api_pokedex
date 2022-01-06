package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Pokemon struct {
	Id   int
	Nome string
	Tipo []string
}

var Pokemons  = []Pokemon{
	{
		Id:   1,
		Nome: "Bulbassauro",
		Tipo: []string{"Grama", "Veneno"},

	},
	{
		Id:   2,
		Nome: "Ivyssauro",
		Tipo: []string{"Grama", "Veneno"},
	},
	{
		Id:   3,
		Nome: "Venossauro",
		Tipo: []string{"Grama", "Veneno"},
	},
	{
		Id:   4,
		Nome: "Charmander",
		Tipo: []string{"Fogo"},
	},
	{
		Id:   5,
		Nome: "Charmeleon",
		Tipo: []string{"Fogo"},
	},
	{
		Id:   6,
		Nome: "Charizard",
		Tipo: []string{"Fogo", "Voador"},
	},
	{
		Id:   7,
		Nome: "Squirtle",
		Tipo: []string{"Água"},
	},
	{
		Id:   8,
		Nome: "Wartortle",
		Tipo: []string{"Água"},
	},
	{
		Id:   9,
		Nome: "Blastoise",
		Tipo: []string{"Água"},
	},
}

func MostrarPokemons(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	encoder.Encode(Pokemons)
}

func CadastrarPokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var novoPokemon Pokemon
	json.Unmarshal(body, &novoPokemon)
	novoPokemon.Id = len(Pokemons) + 1
	novoPokemon.Nome = "Novo"
	novoPokemon.Tipo = []string{"Tipo"}
	Pokemons = append(Pokemons, novoPokemon)

	encoder := json.NewEncoder(w)
	encoder.Encode(Pokemons)
}

func BuscarPokemons(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Println(r.URL.Path)
		valor := strings.Split(r.URL.Path, "/")

		if len(valor) > 9 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		id, erro := strconv.Atoi(valor[2])

		if erro != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		for _, Pokemon := range Pokemons {
			if Pokemon.Id == id {
				json.NewEncoder(w).Encode(Pokemon)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
}

func DeletarPokemons(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusNoContent)
	valor := strings.Split(r.URL.Path, "/")
	id, erro := strconv.Atoi(valor[2])

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	IndicePokemon := -1
	for indice, Pokemon := range Pokemons {
		if Pokemon.Id == id {
			IndicePokemon = indice
			break
		}
	}
	if (IndicePokemon < 0) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	Esquerda := Pokemons[0:IndicePokemon]
	Direita := Pokemons[IndicePokemon + 1 : len(Pokemons)]
	Pokemons = append(Esquerda, Direita...)

}

func EditarPokemons(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusNotFound)
	valor := strings.Split(r.URL.Path, "/")
	id, erro := strconv.Atoi(valor[2])

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
}

	corpo, erroCorpo := ioutil.ReadAll(r.Body)

	if erroCorpo != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var pokemonEditado Pokemon
	erroJson := json.Unmarshal(corpo, &pokemonEditado)

	if erroJson != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	IndicePokemon := -1
	for indice, pokemon := range Pokemons {
		if pokemon.Id == id {
			IndicePokemon = indice
			break
		}
	}
	if IndicePokemon < 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	Pokemons[IndicePokemon] = pokemonEditado

	json.NewEncoder(w).Encode(pokemonEditado)
}

func Mensagem(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Bem-vindos a Pokedéx")
}

