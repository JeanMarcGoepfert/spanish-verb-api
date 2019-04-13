package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"spanish-api/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var DB, _ = data.GetVerbs()

	json.NewEncoder(w).Encode(DB[params["verb"]])
}
