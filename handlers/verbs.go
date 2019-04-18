package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"spanish-api/handlers/utils"
	"spanish-api/lib"
	"spanish-api/repo"
	"strconv"
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
	//TODO: Add pagination meta to response
	params := utils.GetQueryParams(r, []string{"page_number", "page_size"})

	pageNumber, err := strconv.Atoi(params["page_number"])
	if err != nil {
		pageNumber = 1
	}

	pageSize, err := strconv.Atoi(params["page_size"])
	if err != nil {
		pageSize = 10
	}

	json.NewEncoder(w).Encode(repo.GetPaginatedVerbs(pageNumber, pageSize))
}
