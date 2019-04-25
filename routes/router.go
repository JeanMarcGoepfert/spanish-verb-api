package routes

import (
	"github.com/gorilla/mux"
	"spanish-api/handlers"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	router.PathPrefix("/").HandlerFunc(handlers.CatchAll)

	return router
}
