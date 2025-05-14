package main

import (
	"backend/internal/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	}

	// Connect to DB
	if err := db.Connect(); err != nil {
		panic(err)
	}

	// Adding router
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	r.Route("/api", func(r chi.Router) {
	})

	http.ListenAndServe(":8080", r)
}
