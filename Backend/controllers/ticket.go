package controllers

import (
	"Backend/config"
	"Backend/models"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// AllTickets writes the JSON response of all tickets
func AllTickets(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET or not
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Get the client connection
	client := config.ClientConnection()
	// Get the collection
	coll := client.Database("bugTrack").Collection("tickets")
	// Get the cursor
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Create a slice of tickets
	var tickets []models.Ticket
	// Iterate through the cursor
	if err = cursor.All(context.TODO(), &tickets); err != nil {
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
	err = json.NewEncoder(w).Encode(tickets)
	if err != nil {
		return
	}

}

// CreateTicket accepts a JSON request and creates a new ticket
func CreateTicket(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST or not
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	// Create a new project
	var ticket models.CreateTicket
	// Decode the request body into the new project
	err := decoder.Decode(&ticket)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the client connection
	client := config.ClientConnection()
	// Get the collection
	coll := client.Database("bugTrack").Collection("tickets")
	// Insert the project
	result, err := coll.InsertOne(context.TODO(), ticket)

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

// UpdateTicket accepts a JSON request and updates a ticket
func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is PUT or not
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// get the project id from the url
	params := mux.Vars(r)
	ticketID, _ := primitive.ObjectIDFromHex(params["id"])
	// check for the exception where id is not provided
	if ticketID == primitive.NilObjectID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	// Create a new project
	var ticket models.CreateTicket
	// Decode the request body into the new project
	err := decoder.Decode(&ticket)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the client connection
	client := config.ClientConnection()
	// Get the collection
	coll := client.Database("bugTrack").Collection("tickets")
	// filter to update the project with the id provided
	filter := bson.D{{"_id", ticketID}}
	update := bson.D{{"$set", ticket}}
	// Update the project
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

// DeleteTicket accepts the ticketID and deletes the ticket from the database
func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is DELETE or not
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// get the project id from the url
	params := mux.Vars(r)
	ticketID, _ := primitive.ObjectIDFromHex(params["id"])
	// check for the exception where id is not provided
	if ticketID == primitive.NilObjectID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the client connection
	client := config.ClientConnection()
	coll := client.Database("bugTrack").Collection("tickets")
	// filter to delete the project with the id provided
	filter := bson.D{{"_id", ticketID}}
	// Delete the project
	result, err := coll.DeleteOne(context.TODO(), filter)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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

// DeleteProjectTickets accepts the projectID and deletes all the tickets of that project from the database
func DeleteProjectTickets(projectID primitive.ObjectID) bool {
	// Get the client connection
	client := config.ClientConnection()
	coll := client.Database("bugTrack").Collection("tickets")
	// filter to delete the project with the id provided
	filter := bson.D{{"projectID", projectID}}
	// Delete the project
	_, err := coll.DeleteMany(context.TODO(), filter)
	if err != nil {
		return false
	}
	return true
}
