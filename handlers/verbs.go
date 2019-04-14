package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"spanish-api/repo"
)

func GetVerb(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	verb := repo.GetVerb(params["verb"])
	json.NewEncoder(w).Encode(verb)
}
