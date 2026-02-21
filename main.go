package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"todo-api/database"
	"todo-api/handler"
)

func main() {

	db := database.Connect()

	r := chi.NewRouter()

	// ===== Global Middlewares =====
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// ===== Public Route =====
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Todo API"))
	})

	// ===== Auth Routes (Public) =====
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", handler.Register(db))
		r.Post("/login", handler.Login(db))
	})

	// ===== Protected Routes =====
	r.Route("/todos", func(r chi.Router) {

		// 🔐 Auth middleware apply করলাম
		r.Use(handler.AuthMiddleware)

		r.Get("/", handler.GetTodos(db))
		r.Post("/", handler.CreateTodo(db))
		r.Put("/{id}", handler.UpdateTodo(db))
		r.Delete("/{id}", handler.DeleteTodo(db))
	})

	http.ListenAndServe(":8080", r)
}
