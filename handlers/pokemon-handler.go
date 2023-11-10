package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"app/models"

	"github.com/gorilla/mux"
)

const apiURL = "https://pokeapi.co/api/v2/pokemon/"

func GetPokemonInfo(query string) (*models.Pokemon, error) {
	url := fmt.Sprintf("%s%s", apiURL, query)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %s", resp.Status)
	}

	var pokemon models.Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		return nil, err
	}

	return &pokemon, nil
}

func PokemonHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query := vars["query"]

	pokemon, err := GetPokemonInfo(query)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Pokemon ID: %d\n", pokemon.ID)
	fmt.Fprintf(w, "Pokemon Name: %s\n", pokemon.Name)
}
