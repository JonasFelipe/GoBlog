package main

import(
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
	"go-blog/internal/routes"
	"go-blog/internal/database"
)

func main() {
	database.Init()

	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	log.Println("🚀 Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}