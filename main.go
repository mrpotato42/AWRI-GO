package main

import (
	"app/handlers"
	"fmt"
	"log" //modulo para revisar errores por consola con la pagina
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()                              //instanciamos el router
	port := "8000"                                         //definimos el puerto
	fmt.Printf("Server listening on port %s...\n", port)   //imprime en consola el puerto que esta conectado
	router.HandleFunc("/{query}", handlers.PokemonHandler) //definimos una ruta a la cual acceder y se ejecuta la funcion PokemonHandler que estamos importando
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":"+port, nil)) //el modulo log encierra al puerto donde se ejecuta la api

}
