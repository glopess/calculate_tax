package main

import (
	"CalculateTax/api"
	"log"
	"net/http"
)

func main() {
	// Create a new GraphQL handler
	h := api.NewGraphQLHandler()

	// Set up the HTTP server
	http.Handle("/graphql", h)

	log.Println("Server running at http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
