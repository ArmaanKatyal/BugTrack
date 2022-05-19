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
	"time"
)

// AllTickets writes the JSON response of all tickets
func AllTickets(w http.ResponseWriter, r *http.Request) {
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

func Ticket(w http.ResponseWriter, r *http.Request) {
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

	var ticket models.Ticket
	// Get the client connection
	client := config.ClientConnection()
	// Get the collection
	coll := client.Database("bugTrack").Collection("tickets")
	// Get the id from the request
	params := mux.Vars(r)
	ticketID, _ := primitive.ObjectIDFromHex(params["id"])
	// Get the project by id
	err = coll.FindOne(context.TODO(), bson.D{{"_id", ticketID}}).Decode(&ticket)
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
	err = json.NewEncoder(w).Encode(ticket)
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

	Author, err := authenticate(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	decoder := json.NewDecoder(r.Body)
	// Create a new project
	var ticket models.CreateTicket
	// Decode the request body into the new project
	err = decoder.Decode(&ticket)
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

	logColl := client.Database("bugTrack").Collection("logs")
	log := models.Log{
		Type:        "Create",
		Author:      Author,
		Date:        primitive.NewDateTimeFromTime(time.Now()),
		Description: Author + " created a ticket with the id: " + result.InsertedID.(primitive.ObjectID).Hex(),
		Table:       "tickets",
	}
	_, err = logColl.InsertOne(context.TODO(), log)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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

	Author, err := authenticate(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
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
	err = decoder.Decode(&ticket)
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

	logColl := client.Database("bugTrack").Collection("logs")
	log := models.Log{
		Type:        "Update",
		Author:      Author,
		Date:        primitive.NewDateTimeFromTime(time.Now()),
		Description: Author + " updated a ticket with the id: " + ticketID.Hex(),
		Table:       "tickets",
	}
	_, err = logColl.InsertOne(context.TODO(), log)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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

	Author, err := authenticate(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
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

	logColl := client.Database("bugTrack").Collection("logs")
	log := models.Log{
		Type:        "Delete",
		Author:      Author,
		Date:        primitive.NewDateTimeFromTime(time.Now()),
		Description: Author + " deleted a ticket with the id: " + ticketID.Hex(),
		Table:       "tickets",
	}
	_, err = logColl.InsertOne(context.TODO(), log)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
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
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			return
		}
	}()
	coll := client.Database("bugTrack").Collection("tickets")
	// filter to delete the project with the id provided
	filter := bson.D{{"project_id", projectID}}
	// Delete the project
	_, err := coll.DeleteMany(context.TODO(), filter)
	if err != nil {
		return false
	}
	return true
}

func ProjectTickets(w http.ResponseWriter, r *http.Request) {
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

	params := mux.Vars(r)
	projectID, _ := primitive.ObjectIDFromHex(params["projectID"])
	// check for the exception where id is not provided
	if projectID == primitive.NilObjectID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the client connection
	client := config.ClientConnection()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			return
		}
	}()

	coll := client.Database("bugTrack").Collection("tickets")
	cursor, err := coll.Find(context.TODO(), bson.D{{"project_id", projectID}})
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

	// set the status to 200 OK
	w.WriteHeader(http.StatusOK)
	// set the header to application/json
	w.Header().Set("Content-Type", "application/json")
	// Write the JSON response
	err = json.NewEncoder(w).Encode(tickets)
	if err != nil {
		return
	}

}

func FilterTicketsByStatus(w http.ResponseWriter, r *http.Request) {
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
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			return
		}
	}()

	coll := client.Database("bugTrack").Collection("tickets")

	// Get the query parameters
	queryParams := r.URL.Query()
	typeVal := queryParams.Get("type")

	// Check if the type is provided
	if typeVal == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if the type is valid
	if typeVal != "open" && typeVal != "closed" && typeVal != "in-progress" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// filter to get the tickets with the status provided
	filter := bson.D{{"status", typeVal}}
	cursor, err := coll.Find(context.TODO(), filter)
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

	// set the status to 200 OK
	w.WriteHeader(http.StatusOK)
	// set the header to application/json
	w.Header().Set("Content-Type", "application/json")
	// Write the JSON response
	err = json.NewEncoder(w).Encode(tickets)
	if err != nil {
		return
	}
}

func FilterTicketsByPriority(w http.ResponseWriter, r *http.Request) {
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
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			return
		}
	}()

	coll := client.Database("bugTrack").Collection("tickets")

	// Get the query parameters
	queryParams := r.URL.Query()
	typeVal := queryParams.Get("type")

	// Check if the type is provided
	if typeVal == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if the type is valid
	if typeVal != "low" && typeVal != "medium" && typeVal != "high" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// filter to get the tickets with the status provided
	filter := bson.D{{"priority", typeVal}}
	cursor, err := coll.Find(context.TODO(), filter)
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

	// set the status to 200 OK
	w.WriteHeader(http.StatusOK)
	// set the header to application/json
	w.Header().Set("Content-Type", "application/json")
	// Write the JSON response
	err = json.NewEncoder(w).Encode(tickets)
	if err != nil {
		return
	}
}
