package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"spanish-api/handlers/utils"
	"spanish-api/models"
	"spanish-api/repo"
	"strconv"
)

func GetVerb(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	verb := repo.GetVerb(params["verb"])

	if verb == nil {
		utils.SetError(w, http.StatusNotFound, "Verb not found")
		return
	}

	json.NewEncoder(w).Encode(verb)
}

func GetVerbs(w http.ResponseWriter, r *http.Request) {
	params := utils.GetQueryParams(r, []string{"page_number", "page_size"})

	pageNumber, err := strconv.Atoi(params["page_number"])
	if err != nil {
		pageNumber = 1
	}

	pageSize, err := strconv.Atoi(params["page_size"])
	if err != nil {
		pageSize = 10
	}

	type Response struct {
		Meta   models.PaginationMeta `json:"meta"`
		Result []models.Meta         `json:"result"`
	}

	response := Response{
		Meta: models.PaginationMeta{
			Total:      repo.GetVerbCount(),
			PageNumber: pageNumber,
			PageSize:   pageSize,
		},
		Result: repo.GetPaginatedVerbs(pageNumber, pageSize),
	}

	json.NewEncoder(w).Encode(response)
}
