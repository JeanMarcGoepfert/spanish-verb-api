package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"spanish-api/lib"
	"spanish-api/repo"
)

func GetVerb(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	verb := repo.GetVerb(params["verb"])

	if verb == nil {
		lib.SetError(w, http.StatusNotFound, "Verb not found")
		return
	}

	json.NewEncoder(w).Encode(verb)
}
