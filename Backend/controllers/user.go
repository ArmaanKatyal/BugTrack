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
	"time"
)

// AllUsers returns all the users in the database based on the role passed in the url
func AllUsers(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET or not
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// check if the user is authenticated
	Author, err := authenticate(r)
	if err != nil { // if not, return unauthorized
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the client connection
	client := config.ClientConnection()
	defer func() {
		// Disconnect the client
		if err := client.Disconnect(context.TODO()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()

	if verifyAdmin(Author, client) == false { // If the user is not an admin
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var users []models.User           // Create a new slice of users
	role := r.URL.Query().Get("role") // Get the role from the request

	coll := client.Database("bugTrack").Collection("users") // Get the users collection
	if role == "developer" || role == "project-manager" || role == "submitter" || role == "admin" {
		filter := bson.D{{"role", role}}                 // Filter to get the users with the role provided
		cursor, err := coll.Find(context.TODO(), filter) // Get the cursor
		if err != nil {                                  // If there is an error
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Iterate through the cursor
		if err = cursor.All(context.TODO(), &users); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Close the cursor
		if err := cursor.Close(context.TODO()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if role == "" { // If the role is empty then get all the users
		// Get the cursor
		cursor, err := coll.Find(context.TODO(), bson.D{})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Iterate through the cursor
		if err = cursor.All(context.TODO(), &users); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Close the cursor
		if err := cursor.Close(context.TODO()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else { // If the role is not valid
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Set the content type to JSON
	w.WriteHeader(http.StatusOK)                       // Set the status to 200 OK
	err = json.NewEncoder(w).Encode(users)             // Write the JSON response
	if err != nil {
		return
	}

}

// User returns a user from the database
func User(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET or not
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// check if the user is authenticated
	Author, err := authenticate(r)
	if err != nil { // if not, return unauthorized
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var user models.User // Create a new user
	// Get the client connection
	client := config.ClientConnection()

	if verifyAdmin(Author, client) == false { // If the user is not an admin
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

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

// CreateUser creates a new user in the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST or not
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// check if the user is authenticated
	Author, err := authenticate(r)
	if err != nil { // if not, return unauthorized
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	decoder := json.NewDecoder(r.Body) // Create a new decoder
	// Create a new user
	var user models.CreateUser
	// Decode the request body into the new user
	err = decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if UserExists(user.Username) { // Check if the user already exists
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the client connection
	client := config.ClientConnection()

	if verifyAdmin(Author, client) == false { // If the user is not an admin
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the collection
	coll := client.Database("bugTrack").Collection("users")

	// Insert the user into the database
	result, err := coll.InsertOne(context.TODO(), user)

	if err != nil { // If there is an error
		output := struct { // Create a new output
			Status string `json:"status"`
		}{
			Status: "error",
		}
		w.Header().Set("Content-Type", "application/json") // Set the content type
		w.WriteHeader(http.StatusInternalServerError)      // Write the status code
		err = json.NewEncoder(w).Encode(output)            // Encode the output
		if err != nil {
			return
		}
		return
	}

	logColl := client.Database("bugTrack").Collection("logs") // Get the logs collection
	log := models.Log{                                        // Create a new log
		Type:        "Create",
		Author:      Author,
		Date:        primitive.NewDateTimeFromTime(time.Now()),
		Description: "Created user " + user.Username,
		Table:       "users",
	}
	_, err = logColl.InsertOne(context.TODO(), log) // Insert the log into the database
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

// UpdateUser updates the user with the given username
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is PUT or not
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// check if the user is authenticated
	Author, err := authenticate(r)
	if err != nil { // if not, return unauthorized
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

	decoder := json.NewDecoder(r.Body) // Create a new decoder
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

	if verifyAdmin(Author, client) == false { // If the user is not an admin
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

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
		output := struct { // Create a new output
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

	logColl := client.Database("bugTrack").Collection("logs") // Get the logs collection
	log := models.Log{                                        // Create a new log
		Type:        "Update",
		Author:      Author,
		Date:        primitive.NewDateTimeFromTime(time.Now()),
		Description: "Updated user " + username,
		Table:       "users",
	}
	_, err = logColl.InsertOne(context.TODO(), log) // Insert the log into the database
	if err != nil {                                 // If there is an error
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

// DeleteUser deletes the user with the given username
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is DELETE or not
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// check if the user is authenticated
	Author, err := authenticate(r)
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

	if verifyAdmin(Author, client) == false { // If the user is not an admin
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	coll := client.Database("bugTrack").Collection("users")
	// filter to delete the user with the username provided
	filter := bson.D{{"username", username}}
	// Delete the user
	result, err := coll.DeleteOne(context.TODO(), filter)

	if err != nil { // If there is an error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// if the user is not deleted
	if result.DeletedCount == 0 {
		output := struct { // Create a new output
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

	output := struct { // Create a new output
		Status string `json:"status"`
	}{
		Status: "success", // set the status to success
	}

	logColl := client.Database("bugTrack").Collection("logs") // Get the logs collection
	log := models.Log{                                        // Create a new log
		Type:        "Delete",
		Author:      Author,
		Date:        primitive.NewDateTimeFromTime(time.Now()),
		Description: "Deleted user " + username,
		Table:       "users",
	}
	_, err = logColl.InsertOne(context.TODO(), log) // Insert the log into the database
	if err != nil {                                 // If there is an error
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

// CheckUsernameExists checks if the username exists in the database or not
func CheckUsernameExists(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET or not
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// check if the user is authenticated
	_, err := authenticate(r)
	if err != nil { // If there is an error
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	client := config.ClientConnection() // Get the client connection
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
	var result bson.M // Create a new result
	err = coll.FindOne(context.TODO(), bson.D{{"username", params["username"]}}).Decode(&result)
	// If the username is not found
	if err == mongo.ErrNoDocuments {
		// Write the JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		output := struct { // Create a new output
			Exists bool `json:"exists"`
		}{
			Exists: false,
		}
		err = json.NewEncoder(w).Encode(output) // Write the JSON response
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		output := struct { // Create a new output
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

// UserProfile returns the user profile of the user with the given username
func UserProfile(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET or not
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed) // set the status to 405 Method Not Allowed
		return
	}

	// check if the user is authenticated
	Author, err := authenticate(r)
	if err != nil { // If not, return unauthorized
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// The user can only see access his own profile or only the admin can see other users profile

	params := mux.Vars(r) // Get the username from the url
	username := params["username"]
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if Author != username {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	client := config.ClientConnection() // Get the client connection
	defer func() {
		// Disconnect the client
		if err := client.Disconnect(context.TODO()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()

	coll := client.Database("bugTrack").Collection("users")                          // Get the collection
	var user models.User                                                             // Create a new user
	err = coll.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&user) // Find the user in the database
	if err != nil {
		if err == mongo.ErrNoDocuments { // If the user is not found
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError) // If there is an error
		return
	}

	createdTicketColl := client.Database("bugTrack").Collection("tickets") // Get the ticket collection
	// Get the cursor
	cursor, err := createdTicketColl.Find(context.TODO(), bson.D{{"created_by", username}}) // Find the tickets created by the user
	if err != nil {                                                                         // If there is an error
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

	assignedTicketColl := client.Database("bugTrack").Collection("tickets") // Get the tickets collection
	// Get the cursor
	cursor, err = assignedTicketColl.Find(context.TODO(), bson.D{{"assigned_to", username}}) // Find the tickets assigned to the user
	if err != nil {                                                                          // If there is an error
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

	userProfile := models.Profile{ // Create a new user profile
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Username:        user.Username,
		Email:           user.Email,
		TicketsCreated:  createdTicket,
		TicketsAssigned: assignedTicket,
	}
	w.Header().Set("Content-Type", "application/json") // Set the content type to JSON
	w.WriteHeader(http.StatusOK)                       // Set the status to 200 OK
	err = json.NewEncoder(w).Encode(userProfile)       // Write the JSON response
	if err != nil {
		return
	}

}

// UserExists checks if the user with the given username exists
func UserExists(username string) bool {
	client := config.ClientConnection() // Get the client connection
	defer func() {
		// Disconnect the client
		if err := client.Disconnect(context.TODO()); err != nil {
			return
		}
	}()

	coll := client.Database("bugTrack").Collection("users")                           // Get the collection
	var user models.User                                                              // Create a new user
	err := coll.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&user) // Find the user in the database
	if err != nil {                                                                   // If there is an error
		if err == mongo.ErrNoDocuments {
			return false
		}
		return false
	}
	return true
}

// verifyAdmin checks if the user is an admin or not
func verifyAdmin(Author string, client *mongo.Client) bool {
	userColl := client.Database("bugTrack").Collection("users")   // Get the users collection
	filter := bson.D{{"username", Author}}                        // Filter to get the user with the username provided
	var user models.User                                          // Create a new user
	err := userColl.FindOne(context.TODO(), filter).Decode(&user) // Get the user with the username provided

	if err != nil {
		return false
	}

	if user.Role != "admin" { // If the user is not an admin
		return false
	}
	return true
}

// verifyProjectManager checks if the user is a project manager or not
func verifyProjectManager(Author string, projectID primitive.ObjectID, client *mongo.Client) bool {
	coll := client.Database("bugTrack").Collection("projects")   // Get the projects collection
	filter := bson.D{{"_id", projectID}}                         // Filter to get the project with the project id provided
	var project models.Project                                   // Create a new project
	err := coll.FindOne(context.TODO(), filter).Decode(&project) // Get the project with the project id provided
	if err != nil {
		return false
	}

	if project.CreatedBy != Author { // If the user is not the project manager
		return false
	}
	return true
}
