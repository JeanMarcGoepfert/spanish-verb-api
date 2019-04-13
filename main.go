package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"spanish-api/data"
	"spanish-api/middleware"
)

var DB, err = data.GetVerbs()

func GetVerb(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(DB[params["verb"]])
}

func main() {
	if err != nil {
		log.Fatal("Unable to read data. Exiting.")
		os.Exit(1)
	}

	router := mux.NewRouter()

	router.Use(middleware.LogRequest)

	router.HandleFunc("/verb/{verb}", GetVerb).Methods("GET")

	log.Println("Starting server on port 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
