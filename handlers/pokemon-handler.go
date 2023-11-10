package handlers

import (
	"app/models"
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

// URL base de la API de Pokémon
const apiURL = "https://pokeapi.co/api/v2/pokemon/"

// Función para obtener información de un Pokémon mediante su nombre o número
func GetPokemonInfo(query string) (*models.Pokemon, error) {
	// Construir la URL de la API
	url := fmt.Sprintf("%s%s", apiURL, query)

	// Realizar una solicitud HTTP GET a la API
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Comprobar si la respuesta de la API es exitosa (código 200)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %s", resp.Status)
	}

	// Decodificar la respuesta JSON en un struct de Pokémon
	var pokemon models.Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		return nil, err
	}

	// Devolver el struct del Pokémon obtenido
	return &pokemon, nil
}

// Controlador para manejar solicitudes relacionadas con Pokémon
func PokemonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Si es una solicitud GET, servir el archivo de plantilla HTML desde el directorio estático

		// Cargar la plantilla HTML
		tmpl, err := template.ParseFiles("static/templates/index.html")
		if err != nil {
			http.Error(w, "Error al cargar la plantilla", http.StatusInternalServerError)
			return
		}

		// Renderizar la plantilla y enviarla como respuesta
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		// Si es una solicitud POST (por ejemplo, después de enviar un formulario)

		// Procesar el formulario enviado y obtener el número de Pokémon ingresado por el usuario
		query := r.FormValue("pokemonNumber")

		// Obtener información del Pokémon llamando a la función GetPokemonInfo
		pokemon, err := GetPokemonInfo(query)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusInternalServerError)
			return
		}

		// Cargar la plantilla HTML
		tmpl, err := template.ParseFiles("static/templates/index.html")
		if err != nil {
			http.Error(w, "Error al cargar la plantilla", http.StatusInternalServerError)
			return
		}

		// Crear una estructura de datos que contiene la información del Pokémon para renderizar en la plantilla
		data := struct {
			PokemonInfo *models.Pokemon
		}{
			PokemonInfo: pokemon,
		}

		// Renderizar la plantilla con la información del Pokémon y enviarla como respuesta
		tmpl.Execute(w, data)
	}
}
