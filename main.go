package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mjande/grocery-list-generator-microservice/handlers"
)

func main() {
	// Create new router
	router := chi.NewRouter()

	// CORS Middleware
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{os.Getenv("CLIENT_URL")},
		AllowedMethods: []string{"POST"},
		AllowedHeaders: []string{"Content-Type"},
	}))

	// Logging middleware
	router.Use(middleware.Logger)

	// Routes
	router.Post("/generate", handlers.PostGroceryList)

	// Start server
	log.Printf("Grocery list generator service listening on port %s", os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), router)
	if err != nil {
		log.Fatal(err)
	}
}
