package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"my-go-rest-api/models" // Use your actual module path

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func SetClient(mongoClient *mongo.Client) {
	client = mongoClient
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to my API"))
}

func CreateResourceHandler(w http.ResponseWriter, r *http.Request) {
	var resource models.Resource
	_ = json.NewDecoder(r.Body).Decode(&resource)
	resource.ID = primitive.NewObjectID() // Generate a new ObjectID

	collection := client.Database("mydatabase").Collection("resources")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, resource)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resource)
}

func GetResourceHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	collection := client.Database("mydatabase").Collection("resources")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var resource models.Resource
	err := collection.FindOne(ctx, bson.M{"_id": params["id"]}).Decode(&resource)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resource)
}
