package main

import (
	"context"
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"net/http"

	"github.com/gorilla/mux"
)

// Article ...
type Article struct {
	ID      primitive.ObjectID `json:"id,omitempty"      bson:"_id,omitempty"`
	Title   string             `json:"title,omitempty"   bson:"title,omitempty"`
	Content string             `json:"content,omitempty" bson:"content,omitempty"`
	Author  string             `json:"author,omitempty"  bson:"author,omitempty"`
}

// SBSResponse ...
type SBSResponse struct {
	Status  int
	Message string
	Data    *Article
}

var (
	articleCollection *mongo.Collection
)

func main() {
	println("Starting Simple Blog System...")
	initDb()

	r := mux.NewRouter()

	r.HandleFunc("/articles", createArticle).Methods("POST")
	r.HandleFunc("/articles", getArticles).Methods("GET")
	r.HandleFunc("/articles/{article_id}", getArticle).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func writeResponse(w http.ResponseWriter, rStatus int, rMessage string, rData *Article) {
	res := SBSResponse{
		Status:  rStatus,
		Message: rMessage,
		Data:    rData,
	}
	json.NewEncoder(w).Encode(res)
}

func initDb() {
	println("Connecting to DB...")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	articleCollection = client.Database("sbs").Collection("articles")
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	println("Creating article...")
	w.Header().Set("Content-Type", "application/json")

	println("Extracing article values from json request...")
	var article Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	println("Inserting article to mongodb...")
	result, err := articleCollection.InsertOne(nil, article)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	writeResponse(w, http.StatusOK, "Success", &Article{
		ID: result.InsertedID.(primitive.ObjectID),
	})
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	println("Getting all articles...")
	// json.NewEncoder(w).Encode(article.get())
	w.Write([]byte("Getting all articles..."))
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	println("Getting single article...")
	w.Write([]byte("Getting single article..."))
}
