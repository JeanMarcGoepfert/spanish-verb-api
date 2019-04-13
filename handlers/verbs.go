package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"spanish-api/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	verb := data.GetVerbs()[params["verb"]]
	json.NewEncoder(w).Encode(verb)
}
