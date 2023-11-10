package main

import (
	"app/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Crea un nuevo enrutador Gorilla Mux
	router := mux.NewRouter()

	// Define el puerto en el que el servidor escuchará las solicitudes
	port := "8000"

	// Imprime un mensaje en la consola para indicar que el servidor está escuchando
	fmt.Printf("Server listening on port %s...\n", port)

	// Configura un manejador para archivos estáticos
	fs := http.FileServer(http.Dir("static"))

	/* Configura una ruta para servir archivos estáticos y utiliza http.StripPrefix
	   para eliminar el prefijo "/static/" de las rutas de archivos*/
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Define tus rutas de manejo, incluyendo la ruta raíz ("/") que ejecuta el controlador PokemonHandler
	router.HandleFunc("/", handlers.PokemonHandler)

	// Asocia el enrutador con el servidor HTTP
	http.Handle("/", router)

	// Inicia el servidor y comienza a escuchar en el puerto especificado
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
