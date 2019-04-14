package routes

import (
	"net/http"
	"spanish-api/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"verb",
		"GET",
		"/verb/{verb}",
		handlers.GetVerb,
	},
}
