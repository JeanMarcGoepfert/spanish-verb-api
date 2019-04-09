package main

import (
	"encoding/json"
	"github.com/JeanMarcGoepfert/spanish-api/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetVerb(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(data.DB[params["verb"]])
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/verb/{verb}", GetVerb).Methods("GET")

	log.Println("Starting server on port 8080")

	log.Fatal(http.ListenAndServe(":8080", router))

}
