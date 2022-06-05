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
	_, CompanyCode, AuthorRole, err := authenticate(r)
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

	//if verifyAdmin(Author, CompanyCode, client) == false { // If the user is not an admin
	//	w.WriteHeader(http.StatusUnauthorized)
	//	return
	//}
	var users []models.User           // Create a new slice of users
	role := r.URL.Query().Get("role") // Get the role from the request

	coll := client.Database(config.ViperEnvVariable("dbName")).Collection("users") // Get the users collection
	if AuthorRole == "project-manager" {
		if role == "developer" {
			filter := bson.D{{"role", role}, {"company_code", CompanyCode}} // Filter to get the users with the role provided
			cursor, err := coll.Find(context.TODO(), filter)                // Get the cursor
			if err != nil {                                                 // If there is an error
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
		}
	} else if AuthorRole == "admin" { // If the role is empty then get all the users
		if role == "" {
			// Get the cursor
			cursor, err := coll.Find(context.TODO(), bson.D{{"company_code", CompanyCode}})
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
			popIndex := 0
			for index, user := range users {
				if user.Role == "admin" {
					popIndex = index
				}
			}
			users = append(users[:popIndex], users[popIndex+1:]...)
		} else if role == "developer" || role == "project-manager" || role == "submitter" {
			filter := bson.D{{"role", role}, {"company_code", CompanyCode}} // Filter to get the users with the role provided
			cursor, err := coll.Find(context.TODO(), filter)                // Get the cursor
			if err != nil {                                                 // If there is an error
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
	Author, CompanyCode, _, err := authenticate(r)
	if err != nil { // if not, return unauthorized
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var user models.User // Create a new user
	// Get the client connection
	client := config.ClientConnection()

	if verifyAdmin(Author, CompanyCode, client) == false { // If the user is not an admin
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the collection
	coll := client.Database(config.ViperEnvVariable("dbName")).Collection("users")
	// Get the id from the request
	params := mux.Vars(r)
	// Get the project by id
	err = coll.FindOne(context.TODO(), bson.D{{"username", params["username"]}, {"company_code", CompanyCode}}).Decode(&user)
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
	Author, CompanyCode, _, err := authenticate(r)
	if err != nil { // if not, return unauthorized
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	decoder := json.NewDecoder(r.Body) // Create a new decoder
	// Create a new user
	var userDatafromAdmin models.UserDatafromAdmin2
	// Decode the request body into the new user
	err = decoder.Decode(&userDatafromAdmin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if UserExists(userDatafromAdmin.Username, CompanyCode) { // Check if the user already exists
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the client connection
	client := config.ClientConnection()

	if verifyAdmin(Author, CompanyCode, client) == false { // If the user is not an admin
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the collection
	coll := client.Database(config.ViperEnvVariable("dbName")).Collection("users") // Get the users collection
	InsertUser := models.CreateUser{                                               // Create a new user
		FirstName:   userDatafromAdmin.FirstName,
		LastName:    userDatafromAdmin.LastName,
		Username:    userDatafromAdmin.Username,
		Email:       userDatafromAdmin.Email,
		Role:        userDatafromAdmin.Role,
		CreatedOn:   primitive.NewDateTimeFromTime(time.Now()),
		CompanyCode: CompanyCode,
		Locked:      userDatafromAdmin.Locked,
	}

	// Insert the user into the database
	result, err := coll.InsertOne(context.TODO(), InsertUser)

	if err != nil { // If there is an error
		w.WriteHeader(http.StatusInternalServerError) // Write the status code // Encode the output
		return
	}

	authColl := client.Database(config.ViperEnvVariable("dbName")).Collection("auth") // Get the collection
	hashedPass, err := HashedPassword(userDatafromAdmin.Password)                     // Hash the password
	if err != nil {                                                                   // If there is an error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	InsertAuth := models.Credentials{ // Create a new auth
		Username:    userDatafromAdmin.Username,
		Password:    hashedPass,
		CompanyCode: CompanyCode,
	}
	_, err = authColl.InsertOne(context.TODO(), InsertAuth) // Insert the auth into the database
	if err != nil {                                         // If there is an error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logColl := client.Database(config.ViperEnvVariable("dbName")).Collection("logs") // Get the logs collection
	log := models.Log{                                                               // Create a new log
		Type:        "User Created",
		Author:      Author,
		Date:        primitive.NewDateTimeFromTime(time.Now()),
		Description: "Created user " + userDatafromAdmin.Username,
		Table:       "users",
		CompanyCode: CompanyCode,
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
	Author, CompanyCode, _, err := authenticate(r)
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

	if verifyAdmin(Author, CompanyCode, client) == false { // If the user is not an admin
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the collection
	coll := client.Database(config.ViperEnvVariable("dbName")).Collection("users")
	// filter to update the user with the username provided
	filter := bson.D{{"username", username}, {"company_code", CompanyCode}}
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

	logColl := client.Database(config.ViperEnvVariable("dbName")).Collection("logs") // Get the logs collection
	log := models.Log{                                                               // Create a new log
		Type:        "Update",
		Author:      Author,
		Date:        primitive.NewDateTimeFromTime(time.Now()),
		Description: "Updated user " + username,
		Table:       "users",
		CompanyCode: CompanyCode,
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
	Author, CompanyCode, _, err := authenticate(r)
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

	if verifyAdmin(Author, CompanyCode, client) == false { // If the user is not an admin
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	coll := client.Database(config.ViperEnvVariable("dbName")).Collection("users")
	// filter to delete the user with the username provided
	filter := bson.D{{"username", username}, {"company_code", CompanyCode}}
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

	logColl := client.Database(config.ViperEnvVariable("dbName")).Collection("logs") // Get the logs collection
	log := models.Log{                                                               // Create a new log
		Type:        "Delete",
		Author:      Author,
		Date:        primitive.NewDateTimeFromTime(time.Now()),
		Description: "Deleted user " + username,
		Table:       "users",
		CompanyCode: CompanyCode,
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
	_, CompanyCode, _, err := authenticate(r)
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
	coll := client.Database(config.ViperEnvVariable("dbName")).Collection("users")
	// Get the username from the request
	params := mux.Vars(r)
	// find the username in the database
	var result bson.M // Create a new result
	err = coll.FindOne(context.TODO(), bson.D{{"username", params["username"]}, {"company_code", CompanyCode}}).Decode(&result)
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
		w.WriteHeader(http.StatusConflict)
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
	Author, CompanyCode, _, err := authenticate(r)
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

	coll := client.Database(config.ViperEnvVariable("dbName")).Collection("users")                                  // Get the collection
	var user models.User                                                                                            // Create a new user
	err = coll.FindOne(context.TODO(), bson.D{{"username", username}, {"company_code", CompanyCode}}).Decode(&user) // Find the user in the database
	if err != nil {
		if err == mongo.ErrNoDocuments { // If the user is not found
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError) // If there is an error
		return
	}

	createdTicketColl := client.Database(config.ViperEnvVariable("dbName")).Collection("tickets") // Get the ticket collection
	// Get the cursor
	cursor, err := createdTicketColl.Find(context.TODO(), bson.D{{"created_by", username}, {"company_code", CompanyCode}}) // Find the tickets created by the user
	if err != nil {                                                                                                        // If there is an error
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

	assignedTicketColl := client.Database(config.ViperEnvVariable("dbName")).Collection("tickets") // Get the tickets collection
	// Get the cursor
	cursor, err = assignedTicketColl.Find(context.TODO(), bson.D{{"assigned_to", username}, {"company_code", CompanyCode}}) // Find the tickets assigned to the user
	if err != nil {                                                                                                         // If there is an error
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
		Role:            user.Role,
		TicketsCreated:  createdTicket,
		TicketsAssigned: assignedTicket,
		CompanyCode:     CompanyCode,
	}
	w.Header().Set("Content-Type", "application/json") // Set the content type to JSON
	w.WriteHeader(http.StatusOK)                       // Set the status to 200 OK
	err = json.NewEncoder(w).Encode(userProfile)       // Write the JSON response
	if err != nil {
		return
	}

}

// UserExists checks if the user with the given username exists
func UserExists(username string, CompanyCode string) bool {
	client := config.ClientConnection() // Get the client connection
	defer func() {
		// Disconnect the client
		if err := client.Disconnect(context.TODO()); err != nil {
			return
		}
	}()

	coll := client.Database(config.ViperEnvVariable("dbName")).Collection("users")                                   // Get the collection
	var user models.User                                                                                             // Create a new user
	err := coll.FindOne(context.TODO(), bson.D{{"username", username}, {"company_code", CompanyCode}}).Decode(&user) // Find the user in the database
	if err != nil {                                                                                                  // If there is an error
		if err == mongo.ErrNoDocuments {
			return false
		}
		return false
	}
	return true
}

// verifyAdmin checks if the user is an admin or not
func verifyAdmin(Author string, CompanyCode string, client *mongo.Client) bool {
	userColl := client.Database(config.ViperEnvVariable("dbName")).Collection("users") // Get the users collection
	filter := bson.D{{"username", Author}, {"company_code", CompanyCode}}              // Filter to get the user with the username provided
	var user models.User                                                               // Create a new user
	err := userColl.FindOne(context.TODO(), filter).Decode(&user)                      // Get the user with the username provided

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
	coll := client.Database(config.ViperEnvVariable("dbName")).Collection("projects") // Get the projects collection
	filter := bson.D{{"_id", projectID}}                                              // Filter to get the project with the project id provided
	var project models.Project                                                        // Create a new project
	err := coll.FindOne(context.TODO(), filter).Decode(&project)                      // Get the project with the project id provided
	if err != nil {
		return false
	}

	if project.CreatedBy != Author { // If the user is not the project manager
		return false
	}
	return true
}

func Lock(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" { // If the method is not POST
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_, CompanyCode, AuthorRole, err := authenticate(r) // Authenticate the user
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the username from the request
	params := mux.Vars(r)
	username := params["username"]

	if !UserExists(username, CompanyCode) { // If the user does not exist
		w.WriteHeader(http.StatusNotFound)
		return
	}

	client := config.ClientConnection() // Get the client connection
	defer func() {
		// Disconnect the client
		if err := client.Disconnect(context.TODO()); err != nil {
			return
		}
	}()

	if AuthorRole != "admin" { // If the user is not an admin
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	//if !verifyAdmin(Author, CompanyCode, client) { // If the user is not an admin
	//	w.WriteHeader(http.StatusUnauthorized)
	//	return
	//}

	userColl := client.Database(config.ViperEnvVariable("dbName")).Collection("users") // Get the users collection
	filter := bson.D{{"username", username}, {"company_code", CompanyCode}}            // Filter to get the user with the username provided
	update := bson.D{{"$set", bson.D{{"locked", true}}}}                               // Update the user with the locked field set to true
	_, err = userColl.UpdateOne(context.TODO(), filter, update)                        // Update the user with the locked field set to true
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Set the content type to JSON
	w.WriteHeader(http.StatusOK)

	output := struct {
		Message string `json:"message"`
	}{
		Message: "success",
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		return
	}
}

func UnLock(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_, CompanyCode, AuthorRole, err := authenticate(r) // Authenticate the user
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the username from the request
	params := mux.Vars(r)
	username := params["username"]

	if !UserExists(username, CompanyCode) { // If the user does not exist
		w.WriteHeader(http.StatusNotFound)
		return
	}

	client := config.ClientConnection() // Get the client connection
	defer func() {
		// Disconnect the client
		if err := client.Disconnect(context.TODO()); err != nil {
			return
		}
	}()

	//if !verifyAdmin(Author, CompanyCode, client) { // If the user is not an admin
	//	w.WriteHeader(http.StatusUnauthorized)
	//	return
	//}
	if AuthorRole != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	userColl := client.Database(config.ViperEnvVariable("dbName")).Collection("users") // Get the users collection
	filter := bson.D{{"username", username}, {"company_code", CompanyCode}}            // Filter to get the user with the username provided
	update := bson.D{{"$set", bson.D{{"locked", false}}}}                              // Update the user with the locked field set to false
	_, err = userColl.UpdateOne(context.TODO(), filter, update)                        // Update the user with the locked field set to false
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Set the content type to JSON
	w.WriteHeader(http.StatusOK)

	output := struct {
		Message string `json:"message"`
	}{
		Message: "success",
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		return
	}
}

func GetUserRole(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	Author, CompanyCode, _, err := authenticate(r) // Authenticate the user
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	client := config.ClientConnection() // Get the client connection
	defer func() {
		// Disconnect the client
		if err := client.Disconnect(context.TODO()); err != nil {
			return
		}
	}()

	var user models.User
	coll := client.Database(config.ViperEnvVariable("dbName")).Collection("users") // Get the users collection
	filter := bson.D{{"username", Author}, {"company_code", CompanyCode}}          // Filter to get the user with the username provided
	err = coll.FindOne(context.TODO(), filter).Decode(&user)                       // Get the user with the username provided
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Set the content type to JSON
	w.WriteHeader(http.StatusOK)
	output := struct {
		Role string `json:"role"`
	}{
		Role: user.Role,
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		return
	}
}

func LockedUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	Author, CompanyCode, _, err := authenticate(r) // Authenticate the user
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !verifyAdmin(Author, CompanyCode, config.ClientConnection()) { // If the user is not an admin
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

	var users []models.User
	coll := client.Database(config.ViperEnvVariable("dbName")).Collection("users") // Get the users collection
	filter := bson.D{{"company_code", CompanyCode}, {"locked", true}}              // Filter to get the users with the company code provided and locked set to true
	cur, err := coll.Find(context.TODO(), filter)                                  // Get the users with the company code provided and locked set to true
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err = cur.All(context.TODO(), &users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = cur.Close(context.TODO()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Set the content type to JSON
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		return
	}
}
