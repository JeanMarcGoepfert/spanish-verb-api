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

func GetVerbs(w http.ResponseWriter, r *http.Request) {
	//TODO: Get these values out of r.URL.Query
	pageNumber := 1
	pageSize := 10
	foo := repo.GetPaginatedVerbs(pageNumber, pageSize)

	json.NewEncoder(w).Encode(foo)
}
