package main

import (
	"backend/internal/db"
	"backend/internal/handlers"
	"backend/internal/repositories"
	"backend/internal/services"
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

	//Users
	userRepo := repositories.NewUserRepository(db.DB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Adding router
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	r.Group(func(r chi.Router) {
		r.Route("/api", func(r chi.Router) {
			r.Get("/users", userHandler.GetAllUsers)
			r.Post("/users", userHandler.CreateUser)
		})
	})

	http.ListenAndServe(":8080", r)
}
