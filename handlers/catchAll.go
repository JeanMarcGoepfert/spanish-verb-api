package handlers

import (
	"net/http"
	"spanish-api/handlers/utils"
)

func CatchAll(w http.ResponseWriter, r *http.Request) {
	utils.SetError(w, http.StatusNotFound, "Route not found")

	return
}
