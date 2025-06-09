package handlers 

import (
	"database/sql"
	"enconding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go-blog/internal/database"
	"go-blog/internal/models"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {	
	rows, err := database.DB.Query("SELECT id, tittle, content FROM posts")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var p models.Post
		if err := rows.Scan(&p.ID, &p.Title, &p.Content); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err := database.DB.Exec("INSERT INTO posts (title, content) VALUES ($1, $2)", P.Title, p.Content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var p models.Post
	err = database.DB.QueryRow("SELECT id, title, content FROM posts WHERE id=$1").Scan(&p.ID, &p.Title, &p.Content)
	if err == sql.ErrNoRows {
		http.Error(w, "Post not Found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEnconder(w).Encode(p)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	_, err = database.DB.Exec("DELETE FROM posts WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

