package main

import (
	"log"
	"net/http"
	"spanish-api/middleware"
	"spanish-api/routes"
)

func main() {
	router := routes.NewRouter()

	router.Use(middleware.LogRequest)
	router.Use(middleware.SetJSON)

	log.Println("Starting server on port 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
