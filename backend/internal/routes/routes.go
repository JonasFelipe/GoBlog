package routes

import (
	"github.com/gorilla/mux"
	"go-blog/internal/handlers"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/posts", handlers.GetPosts).Methods("GET")
	r.HandleFunc("/api/posts", handlers.CreatePost).Methods("POST")
	r.HandleFunc("/api/posts/{id}", handlers.GetPost).Methods("GET")
	r.HandleFunc("/api/posts/{id}", handlers.DeletePost).Methods("DELETE")
}