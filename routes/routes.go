package routes

import (
	"context"
	"fmt"
	"log"
	"time"

	"my-go-rest-api/handlers" // Use your actual module path

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb+srv://hamzamushtaq362:hamza12345@cluster0.tm5ay25.mongodb.net/?retryWrites=true&w=majority&appName=Cluster07")

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

func RegisterRoutes(client *mongo.Client) *mux.Router {
	r := mux.NewRouter()

	// Set MongoDB client in handlers
	handlers.SetClient(client)

	// Define routes
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/api/resource", handlers.CreateResourceHandler).Methods("POST")
	r.HandleFunc("/api/resource/{id}", handlers.GetResourceHandler).Methods("GET")

	return r
}
