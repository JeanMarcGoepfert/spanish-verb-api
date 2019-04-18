package utils

import (
	"log"
	"net/http"
)

type QueryParams map[string]string

func GetQueryParams(r *http.Request, params []string) QueryParams {
	result := make(QueryParams)

	for _, key := range params {
		result[key] = r.URL.Query().Get(key)
	}

	log.Println(result)

	return result
}
