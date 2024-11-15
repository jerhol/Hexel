package main

import (
	"log"
	"net/http"
	"os"

	api "github.com/jerhol/Hexel/api/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env found or unable to load it")
	}

	port := os.Getenv("PORT")

	mux := api.RegisterRoutes()
	log.Printf("Starting server on port %v", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
