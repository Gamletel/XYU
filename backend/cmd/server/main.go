package main

import (
	"backend/internal/db"
	"backend/internal/handlers"
	"backend/internal/middlewares"
	"backend/internal/repositories"
	"backend/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"net/http"
	"time"
)

func main() {
	// Connect to DB
	if err := db.Connect(); err != nil {
		panic(err)
	}

	// Users
	userRepo := repositories.NewUserRepository(db.DB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	//Todos
	todoRepo := repositories.NewTodoRepository(db.DB)
	todoService := services.NewTodoService(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoService)

	// Adding router
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(httprate.LimitByIP(100, time.Minute))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	r.Group(func(r chi.Router) {
		r.Route("/api", func(r chi.Router) {
			r.Post("/login", handlers.Login)
			r.Post("/register", handlers.Register)

			r.Route("/v1", func(r chi.Router) {
				r.Use(middlewares.Auth)

				r.Route("/users", func(r chi.Router) {
					r.Get("/", userHandler.GetAllUsers)
					r.Get("/by-email/{email}", userHandler.GetUserByEmail)
					r.Post("/", userHandler.CreateUser)
					r.Put("/{id}", userHandler.UpdateUser)
					r.Delete("/{id}", userHandler.DeleteUser)
				})

				r.Route("/todos", func(r chi.Router) {
					r.Get("/by-title", todoHandler.GetTodoByTitle)
					r.Get("/by-user-id", todoHandler.GetTodoByUserId)
					r.Post("/", todoHandler.CreateTodo)
					r.Put("/{id}", todoHandler.UpdateTodo)
					r.Delete("/{id}", todoHandler.DeleteTodo)
				})
			})
		})
	})

	http.ListenAndServe(":8080", r)
}
