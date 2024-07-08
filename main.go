package main

import (
	"fmt"
	"log"
	"net/http"

	"my-go-rest-api/routes"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func main() {
	client = routes.ConnectDB() // Connect to MongoDB

	r := routes.RegisterRoutes(client) // Register routes with the router

	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r)) // Start the server
}
