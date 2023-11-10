package models

// Definición de la estructura Pokemon
type Pokemon struct {
	ID     int    `json:"id"`    // Campo ID que representa el identificador del Pokémon
	Name   string `json:"name"`  // Campo Name que representa el nombre del Pokémon
	Types  []Type `json:"types"` // Campo Types que es una lista de tipos del Pokémon
	Sprite struct {
		FrontDefault string `json:"front_default"` // Campo Sprite que almacena la URL de la imagen del Pokémon
	} `json:"sprites"`
}

// Definición de la estructura Type
type Type struct {
	Slot       int `json:"slot"` // Campo Slot que representa la ranura del tipo del Pokémon
	TypeDetail struct {
		Name string `json:"name"` // Campo Name que representa el nombre del tipo del Pokémon
		URL  string `json:"url"`  // Campo URL que almacena la URL relacionada con el tipo del Pokémon
	} `json:"type"`
}
