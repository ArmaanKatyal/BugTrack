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

// AllProjects writes the JSON response of all projects
func AllProjects(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET or not
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Get the client connection
	client := config.ClientConnection()
	// Get the collection
	coll := client.Database("bugTrack").Collection("projects")
	// Get the cursor
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Create a slice of projects
	var projects []models.Project
	// Iterate through the cursor
	if err = cursor.All(context.TODO(), &projects); err != nil {
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

	// Marshal the projects into JSON
	w.Header().Set("Content-Type", "application/json")
	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(projects)
	if err != nil {
		return
	}
}

// CreateProject accepts a JSON request and creates a new project
func CreateProject(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST or not
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	// Create a new project
	var project models.CreateProject
	// Decode the request body into the new project
	err := decoder.Decode(&project)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the client connection
	client := config.ClientConnection()
	// Get the collection
	coll := client.Database("bugTrack").Collection("projects")
	// Insert the project
	result, err := coll.InsertOne(context.TODO(), project)

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

// UpdateProject accepts a JSON request and updates a project
func UpdateProject(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is PUT or not
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// get the project id from the url
	params := mux.Vars(r)
	projectID, _ := primitive.ObjectIDFromHex(params["id"])
	// check for the exception where id is not provided
	if projectID == primitive.NilObjectID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	// Create a new project
	var project models.CreateProject
	// Decode the request body into the new project
	err := decoder.Decode(&project)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the client connection
	client := config.ClientConnection()
	// Get the collection
	coll := client.Database("bugTrack").Collection("projects")
	// filter to update the project with the id provided
	filter := bson.D{{"_id", projectID}}
	update := bson.D{{"$set", project}}
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

// DeleteProject accepts the projectID and deletes a project and all the tickets associated with it
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is DELETE or not
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// get the project id from the url
	params := mux.Vars(r)
	projectID, _ := primitive.ObjectIDFromHex(params["id"])
	// check for the exception where id is not provided
	if projectID == primitive.NilObjectID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the client connection
	client := config.ClientConnection()
	coll := client.Database("bugTrack").Collection("projects")
	// filter to delete the project with the id provided
	filter := bson.D{{"_id", projectID}}
	// Delete the project
	result, err := coll.DeleteOne(context.TODO(), filter)

	// If the project is deleted then all the existing tickets in the project are also
	// deleted. So we need to delete all the tickets in the project

	// delete all the tickets of the project
	ticketDeleteResult := DeleteProjectTickets(projectID)
	if ticketDeleteResult == false {
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
