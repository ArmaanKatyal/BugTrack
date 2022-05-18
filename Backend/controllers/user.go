package controllers

import (
	"Backend/config"
	"Backend/models"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET or not
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_, err := authenticate(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the client connection
	client := config.ClientConnection()
	// Get the collection
	coll := client.Database("bugTrack").Collection("users")
	// Get the cursor
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Create a slice of tickets
	var Users []models.User
	// Iterate through the cursor
	if err = cursor.All(context.TODO(), &Users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer func() {
		// Close the cursor
		if err := cursor.Close(context.TODO()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Disconnect the client
		if err := client.Disconnect(context.TODO()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()

	// Marshal the tickets into JSON
	w.Header().Set("Content-Type", "application/json")
	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(Users)
	if err != nil {
		return
	}

}

func User(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET or not
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_, err := authenticate(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var user models.User
	// Get the client connection
	client := config.ClientConnection()
	// Get the collection
	coll := client.Database("bugTrack").Collection("users")
	// Get the id from the request
	params := mux.Vars(r)
	// Get the project by id
	err = coll.FindOne(context.TODO(), bson.D{{"username", params["username"]}}).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer func() {
		// Disconnect the client
		if err := client.Disconnect(context.TODO()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()

	// Marshal the project into JSON
	w.Header().Set("Content-Type", "application/json")
	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST or not
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_, err := authenticate(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	decoder := json.NewDecoder(r.Body)
	// Create a new user
	var user models.CreateUser
	// Decode the request body into the new user
	err = decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the client connection
	client := config.ClientConnection()
	// Get the collection
	coll := client.Database("bugTrack").Collection("users")

	// TODO : Check if the user already exists before creating it

	// Insert the user into the database
	result, err := coll.InsertOne(context.TODO(), user)

	if err != nil {
		output := struct {
			Status string `json:"status"`
		}{
			Status: "error",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(output)
		if err != nil {
			return
		}
		return
	}

	// set the status to 201 Created
	w.WriteHeader(http.StatusCreated)
	// set the header to application/json
	w.Header().Set("Content-Type", "application/json")

	// struct to hold the id and status of the new project created
	output := struct {
		Status string `json:"status"`
		ID     string `json:"id"`
	}{
		Status: "success",
		ID:     result.InsertedID.(primitive.ObjectID).Hex(),
	}

	// Write the JSON response
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		return
	}

	defer func() {
		// Disconnect the client
		if err := client.Disconnect(context.TODO()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()

}

// UpdateUser updates the user with the given username
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is PUT or not
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_, err := authenticate(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// get the username from the url
	params := mux.Vars(r)
	username := params["username"]
	// check for the exception where username is not provided
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	// Create a new user
	var user models.CreateUser
	// Decode the request body into the new user
	err = decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the client connection
	client := config.ClientConnection()
	// Get the collection
	coll := client.Database("bugTrack").Collection("users")
	// filter to update the user with the username provided
	filter := bson.D{{"username", username}}
	update := bson.D{{"$set", user}}
	// Update the user
	result, err := coll.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		output := struct {
			Status string `json:"status"`
		}{
			Status: "error",
		}
		// set the status to 500 Internal Server Error
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		// Write the JSON response
		err = json.NewEncoder(w).Encode(output)
		if err != nil {
			return
		}
		return
	}

	// if the values are not updated
	if result.ModifiedCount == 0 {
		output := struct {
			Status string `json:"status"`
		}{
			Status: "error",
		}
		// set the status to 404 Not Found
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		// Write the JSON response
		err = json.NewEncoder(w).Encode(output)
		if err != nil {
			return
		}
		return
	}

	// set the status to 200 OK
	w.WriteHeader(http.StatusOK)
	// set the header to application/json
	w.Header().Set("Content-Type", "application/json")
	output := struct {
		Status string `json:"status"`
	}{
		Status: "success",
	}
	// Write the JSON response
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		return
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()

}

// DeleteUser deletes the user with the given username
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is DELETE or not
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_, err := authenticate(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// get the username from the url
	params := mux.Vars(r)
	username := params["username"]
	// check for the exception where username is not provided
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the client connection
	client := config.ClientConnection()
	coll := client.Database("bugTrack").Collection("users")
	// filter to delete the user with the username provided
	filter := bson.D{{"username", username}}
	// Delete the user
	result, err := coll.DeleteOne(context.TODO(), filter)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// if the user is not deleted
	if result.DeletedCount == 0 {
		output := struct {
			Status string `json:"status"`
		}{
			Status: "error",
		}
		// set the status to 404 Not Found
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		// Write the JSON response
		err = json.NewEncoder(w).Encode(output)
		if err != nil {
			return
		}
		return
	}

	output := struct {
		Status string `json:"status"`
	}{
		Status: "success",
	}

	// set the status to 200 OK
	w.WriteHeader(http.StatusOK)
	// set the header to application/json
	w.Header().Set("Content-Type", "application/json")
	// Write the JSON response
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		return
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()

}

// CheckUsernameExists checks if the username exists in the database or not
func CheckUsernameExists(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET or not
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	client := config.ClientConnection()
	defer func() {
		// Disconnect the client
		if err := client.Disconnect(context.TODO()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()
	// Get the collection
	coll := client.Database("bugTrack").Collection("users")
	// Get the username from the request
	params := mux.Vars(r)
	// find the username in the database
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{"username", params["username"]}}).Decode(&result)
	// If the username is not found
	if err == mongo.ErrNoDocuments {
		// Write the JSON response
		w.WriteHeader(http.StatusOK)
		output := struct {
			Exists bool `json:"exists"`
		}{
			Exists: false,
		}
		err = json.NewEncoder(w).Encode(output)
		if err != nil {
			return
		}
		return
	} else if err != nil {
		// if the server encounters an error
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		// if the username already exists
		w.WriteHeader(http.StatusOK)
		output := struct {
			Exists bool `json:"exists"`
		}{
			Exists: true,
		}
		// Write the JSON response
		err = json.NewEncoder(w).Encode(output)
		if err != nil {
			return
		}
		return
	}
}

func UserProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_, err := authenticate(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	username := params["username"]
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	client := config.ClientConnection()
	defer func() {
		// Disconnect the client
		if err := client.Disconnect(context.TODO()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()

	coll := client.Database("bugTrack").Collection("users")
	var user models.User
	err = coll.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	createdTicketColl := client.Database("bugTrack").Collection("tickets")
	// Get the cursor
	cursor, err := createdTicketColl.Find(context.TODO(), bson.D{{"created_by", username}})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Create a slice of tickets
	var createdTicket []models.Ticket
	// Iterate through the cursor
	if err = cursor.All(context.TODO(), &createdTicket); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	assignedTicketColl := client.Database("bugTrack").Collection("tickets")
	// Get the cursor
	cursor, err = assignedTicketColl.Find(context.TODO(), bson.D{{"assigned_to", username}})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Create a slice of tickets
	var assignedTicket []models.Ticket
	// Iterate through the cursor
	if err = cursor.All(context.TODO(), &assignedTicket); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userProfile := models.Profile{
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Username:        user.Username,
		Email:           user.Email,
		TicketsCreated:  createdTicket,
		TicketsAssigned: assignedTicket,
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(userProfile)
	if err != nil {
		return
	}

}
