package main

import (
	"main/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":8000", router)
}
