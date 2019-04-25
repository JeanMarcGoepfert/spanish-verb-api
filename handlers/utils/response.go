package utils

import (
	"encoding/json"
	"net/http"
	"spanish-api/models"
)

func SetError(w http.ResponseWriter, status int, message string) http.ResponseWriter {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(models.Error{Message: message, Status: status})
	return w
}
