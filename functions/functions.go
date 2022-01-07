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
	Name string
	Type []string
}

var Pokemons = []Pokemon{
	{
		Id:   1,
		Name: "Bulbassaur",
		Type: []string{"Grass", "Poison"},
	},
	{
		Id:   2,
		Name: "Ivyssaur",
		Type: []string{"Grass", "Poison"},
	},
	{
		Id:   3,
		Name: "Venossaur",
		Type: []string{"Grass", "Poison"},
	},
	{
		Id:   4,
		Name: "Charmander",
		Type: []string{"Fire"},
	},
	{
		Id:   5,
		Name: "Charmeleon",
		Type: []string{"Fire"},
	},
	{
		Id:   6,
		Name: "Charizard",
		Type: []string{"Fire", "Flying"},
	},
	{
		Id:   7,
		Name: "Squirtle",
		Type: []string{"Water"},
	},
	{
		Id:   8,
		Name: "Wartortle",
		Type: []string{"Water"},
	},
	{
		Id:   9,
		Name: "Blastoise",
		Type: []string{"Water"},
	},
}

func ShowPokemons(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	encoder.Encode(Pokemons)
}

func RegisterPokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	body, error := ioutil.ReadAll(r.Body)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var NewPokemon Pokemon
	json.Unmarshal(body, &NewPokemon)
	NewPokemon.Id = len(Pokemons) + 1
	NewPokemon.Name = "New"
	NewPokemon.Type = []string{"Type"}
	Pokemons = append(Pokemons, NewPokemon)

	encoder := json.NewEncoder(w)
	encoder.Encode(Pokemons)
}

func SearchPokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println(r.URL.Path)
	Value := strings.Split(r.URL.Path, "/")

	if len(Value) > 9 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	id, error := strconv.Atoi(Value[2])

	if error != nil {
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

func DeletePokemons(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	Value := strings.Split(r.URL.Path, "/")
	id, error := strconv.Atoi(Value[2])

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	IndexPokemon := -1
	for Index, Pokemon := range Pokemons {
		if Pokemon.Id == id {
			IndexPokemon = Index
			break
		}
	}
	if IndexPokemon < 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	Left := Pokemons[0:IndexPokemon]
	Right := Pokemons[IndexPokemon+1 : len(Pokemons)]
	Pokemons = append(Left, Right...)

}

func EditPokemons(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	Value := strings.Split(r.URL.Path, "/")
	id, error := strconv.Atoi(Value[2])

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, errorBody := ioutil.ReadAll(r.Body)

	if errorBody != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var pokemonEdited Pokemon
	errorJson := json.Unmarshal(body, &pokemonEdited)

	if errorJson != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	IndexPokemon := -1
	for Index, pokemon := range Pokemons {
		if pokemon.Id == id {
			IndexPokemon = Index
			break
		}
	}
	if IndexPokemon < 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	Pokemons[IndexPokemon] = pokemonEdited

	json.NewEncoder(w).Encode(pokemonEdited)
}

func Message(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to Pokedex")
}
