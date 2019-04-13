package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"spanish-api/handlers"
	"spanish-api/middleware"
)

func main() {
	router := mux.NewRouter()

	router.Use(middleware.LogRequest)

	router.HandleFunc("/verb/{verb}", handlers.Get).Methods("GET")

	log.Println("Starting server on port 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
