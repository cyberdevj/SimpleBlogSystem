package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	println("Starting new server")
	r := mux.NewRouter()

	r.HandleFunc("/articles", createArticle).Methods("POST")
	r.HandleFunc("/articles", getArticles).Methods("GET")
	r.HandleFunc("/articles/{article_id}", getArticle).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	println("Creating article...")
	w.Write([]byte("Creating article..."))
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	println("Getting all articles...")
	w.Write([]byte("Getting all articles..."))
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	println("Getting single article...")
	w.Write([]byte("Getting single article..."))
}
