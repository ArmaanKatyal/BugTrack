package controllers

import (
	"Backend/config"
	"Backend/models"
	"container/list"
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

	// check if the user is authenticated and logged in
	Author, CompanyCode, AuthorRole, err := authenticate(r)
	// if not, return unauthorized
	if err != nil {
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

	//userColl := client.Database(config.ViperEnvVariable("dbName")).Collection("users") // get the users collection
	//filter := bson.D{{"username", Author}, {"company_code", CompanyCode}}              // filter to get the user by id
	//var user models.User                                                               // create a new user
	//err = userColl.FindOne(context.TODO(), filter).Decode(&user)
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}

	// Create a slice of tickets
	var tickets []models.Ticket

	if AuthorRole == "admin" { // if the user is an admin, get all the tickets
		// Get the collection
		coll := client.Database(config.ViperEnvVariable("dbName")).Collection("tickets")
		// Get the cursor
		cursor, err := coll.Find(context.TODO(), bson.D{{"company_code", CompanyCode}})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Iterate through the cursor
		if err = cursor.All(context.TODO(), &tickets); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := cursor.Close(context.TODO()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if AuthorRole == "developer" { // if the user is a developer, get all the tickets assigned to him
		coll := client.Database(config.ViperEnvVariable("dbName")).Collection("tickets")
		cursor, err := coll.Find(context.TODO(), bson.D{{"assigned_to", Author}, {"company_code", CompanyCode}})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err = cursor.All(context.TODO(), &tickets); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := cursor.Close(context.TODO()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if AuthorRole == "submitter" { // if the user is a submitter, get all the tickets submitted by him
		coll := client.Database(config.ViperEnvVariable("dbName")).Collection("tickets")
		cursor, err := coll.Find(context.TODO(), bson.D{{"created_by", Author}, {"company_code", CompanyCode}})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err = cursor.All(context.TODO(), &tickets); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := cursor.Close(context.TODO()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if AuthorRole == "project-manager" { // if the user is a project manager, get all the tickets assigned to the project he is in
		//projectColl := client.Database(config.ViperEnvVariable("dbName")).Collection("projects")
		//filter := bson.D{{"project_manager", Author}, {"company_code", CompanyCode}}
		//var project models.Project
		//err = projectColl.FindOne(context.TODO(), filter).Decode(&project)
		//if err != nil {
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		//coll := client.Database(config.ViperEnvVariable("dbName")).Collection("tickets")
		//cursor, err := coll.Find(context.TODO(), bson.D{{"project_id", project.Id}, {"company_code", CompanyCode}})
		//if err != nil {
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		//if err = cursor.All(context.TODO(), &tickets); err != nil {
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		//if err := cursor.Close(context.TODO()); err != nil {
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		projectColl := client.Database(config.ViperEnvVariable("dbName")).Collection("projects")
		filter := bson.D{{"project_manager", Author}, {"company_code", CompanyCode}}
		var projects []models.Project
		cursor, err := projectColl.Find(context.TODO(), filter)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err = cursor.All(context.TODO(), &projects); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := cursor.Close(context.TODO()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		ticketColl := client.Database(config.ViperEnvVariable("dbName")).Collection("tickets")
		ticketList := list.New()
		for _, project := range projects {
			cursor, err := ticketColl.Find(context.TODO(), bson.D{{"project_id", project.Id}, {"company_code", CompanyCode}})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			for cursor.Next(context.TODO()) {
				var ticket models.Ticket
				err = cursor.Decode(&ticket)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				ticketList.PushBack(ticket)
			}
			if err := cursor.Close(context.TODO()); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		for e := ticketList.Front(); e != nil; e = e.Next() {
			tickets = append(tickets, e.Value.(models.Ticket))
		}
	}

	// Marshal the tickets into JSON
	w.Header().Set("Content-Type", "application/json")
	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(tickets) // encode the tickets into JSON
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

	// check if the user is authenticated and logged in
	_, CompanyCode, _, err := authenticate(r)
	// if not, return unauthorized
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var ticket models.Ticket //	create a new ticket
	// Get the client connection
	client := config.ClientConnection()
	// Get the collection
	coll := client.Database(config.ViperEnvVariable("dbName")).Collection("tickets")
	// Get the id from the request
	params := mux.Vars(r)
	ticketID, _ := primitive.ObjectIDFromHex(params["id"])
	// Get the project by id
	filter := bson.D{{"_id", ticketID}, {"company_code", CompanyCode}}
	err = coll.FindOne(context.TODO(), filter).Decode(&ticket)
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

	// check if the user is authenticated and logged in
	Author, CompanyCode, _, err := authenticate(r)
	// if not, return unauthorized
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	decoder := json.NewDecoder(r.Body)
	// Create a new project
	var ticket models.CreateTicket2
	// Decode the request body into the new project
	err = decoder.Decode(&ticket)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var createTicket models.CreateTicket
	createTicket.Title = ticket.Title
	createTicket.Description = ticket.Description
	createTicket.Status = "open"
	createTicket.Priority = ticket.Priority
	createTicket.Tags = ticket.Tags
	createTicket.CreatedBy = Author
	createTicket.AssignedTo = ticket.AssignedTo
	createTicket.CreatedOn = primitive.NewDateTimeFromTime(time.Now())
	createTicket.UpdatedOn = primitive.NewDateTimeFromTime(time.Now())
	createTicket.ProjectId = ticket.ProjectId
	createTicket.ProjectName = ticket.ProjectName
	createTicket.CompanyCode = CompanyCode

	// Get the client connection
	client := config.ClientConnection()
	// Get the collection
	coll := client.Database(config.ViperEnvVariable("dbName")).Collection("tickets")
	// Insert the project
	result, err := coll.InsertOne(context.TODO(), createTicket)

	if err != nil {
		// if there is an error
		output := struct {
			Status string `json:"status"`
		}{
			Status: "error",
		}
		w.Header().Set("Content-Type", "application/json") // set the content type
		w.WriteHeader(http.StatusInternalServerError)      // set the status code
		err = json.NewEncoder(w).Encode(output)            // encode the output
		if err != nil {
			return
		}
		return
	}

	logColl := client.Database(config.ViperEnvVariable("dbName")).Collection("logs") // get the logs collection
	log := models.Log{                                                               // create a new log
		Type:        "Create",
		Author:      Author,
		Date:        primitive.NewDateTimeFromTime(time.Now()),
		Description: Author + " created a ticket with the id: " + result.InsertedID.(primitive.ObjectID).Hex(),
		Table:       "tickets",
		CompanyCode: CompanyCode,
	}
	_, err = logColl.InsertOne(context.TODO(), log) // insert the log
	if err != nil {                                 // if there is an error
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

	// check if the user is authenticated and logged in
	Author, CompanyCode, _, err := authenticate(r)
	if err != nil { // if not, return unauthorized
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
	var ticket models.CreateTicket3
	// Decode the request body into the new project
	err = decoder.Decode(&ticket)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var updateTicket models.CreateTicket
	updateTicket.Title = ticket.Title
	updateTicket.Description = ticket.Description
	updateTicket.Status = ticket.Status
	updateTicket.Priority = ticket.Priority
	updateTicket.Tags = ticket.Tags
	updateTicket.AssignedTo = ticket.AssignedTo
	updateTicket.UpdatedOn = primitive.NewDateTimeFromTime(time.Now())
	updateTicket.CompanyCode = CompanyCode
	updateTicket.CreatedBy = Author
	updateTicket.ProjectName = ticket.ProjectName
	updateTicket.ProjectId = ticket.ProjectId

	// Get the client connection
	client := config.ClientConnection()
	// Get the collection
	coll := client.Database(config.ViperEnvVariable("dbName")).Collection("tickets")
	// filter to update the project with the id provided
	filter := bson.D{{"_id", ticketID}}
	update := bson.D{{"$set", updateTicket}}
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

	if result.ModifiedCount == 0 { // if the project is not updated
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

	logColl := client.Database(config.ViperEnvVariable("dbName")).Collection("logs") // get the logs collection
	log := models.Log{                                                               // create a new log
		Type:        "Update",
		Author:      Author,
		Date:        primitive.NewDateTimeFromTime(time.Now()),
		Description: Author + " updated a ticket with the id: " + ticketID.Hex(),
		Table:       "tickets",
		CompanyCode: CompanyCode,
	}
	_, err = logColl.InsertOne(context.TODO(), log) // insert the log
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// set the status to 200 OK
	w.WriteHeader(http.StatusOK)
	// set the header to application/json
	w.Header().Set("Content-Type", "application/json")
	output := struct { // struct to hold the status of the project updated
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

	// check if the user is authenticated and logged in
	Author, CompanyCode, _, err := authenticate(r)
	if err != nil { // if not, return unauthorized
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
	coll := client.Database(config.ViperEnvVariable("dbName")).Collection("tickets")
	// filter to delete the project with the id provided
	filter := bson.D{{"_id", ticketID}, {"company_code", CompanyCode}}
	// Delete the project
	result, err := coll.DeleteOne(context.TODO(), filter)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 { // if the project is not deleted
		output := struct { // struct to hold the status of the project deleted
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

	output := struct { // struct to hold the status of the project deleted
		Status string `json:"status"`
	}{
		Status: "success",
	}

	logColl := client.Database(config.ViperEnvVariable("dbName")).Collection("logs") // get the logs collection
	log := models.Log{                                                               // create a new log
		Type:        "Delete",
		Author:      Author,
		Date:        primitive.NewDateTimeFromTime(time.Now()),
		Description: Author + " deleted a ticket with the id: " + ticketID.Hex(),
		Table:       "tickets",
		CompanyCode: CompanyCode,
	}
	_, err = logColl.InsertOne(context.TODO(), log) // insert the log
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
	coll := client.Database(config.ViperEnvVariable("dbName")).Collection("tickets")
	// filter to delete the project with the id provided
	filter := bson.D{{"project_id", projectID}}
	// Delete the project
	_, err := coll.DeleteMany(context.TODO(), filter)
	if err != nil {
		return false
	}
	return true // return true if the tickets are deleted
}

func ProjectTickets(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET or not
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// check if the user is authenticated and logged in
	_, CompanyCode, _, err := authenticate(r)
	if err != nil { // if not, return unauthorized
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r) // get the project id from the url
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

	coll := client.Database(config.ViperEnvVariable("dbName")).Collection("tickets")                           // get the tickets collection
	cursor, err := coll.Find(context.TODO(), bson.D{{"project_id", projectID}, {"company_code", CompanyCode}}) // find the tickets of the project
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

	// set the header to application/json
	w.Header().Set("Content-Type", "application/json")
	// set the status to 200 OK
	w.WriteHeader(http.StatusOK)
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

	// check if the user is authenticated and logged in
	_, CompanyCode, _, err := authenticate(r)
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

	coll := client.Database(config.ViperEnvVariable("dbName")).Collection("tickets") // get the tickets collection

	// Get the query parameters
	queryParams := r.URL.Query() // get the query parameters
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
	filter := bson.D{{"status", typeVal}, {"company_code", CompanyCode}}
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

	// set the header to application/json
	w.Header().Set("Content-Type", "application/json")
	// set the status to 200 OK
	w.WriteHeader(http.StatusOK)
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

	// check if the user is authenticated and logged in
	_, CompanyCode, _, err := authenticate(r)
	if err != nil { // if not, return unauthorized
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

	coll := client.Database(config.ViperEnvVariable("dbName")).Collection("tickets") // get the tickets collection

	// Get the query parameters
	queryParams := r.URL.Query() // get the query parameters
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
	filter := bson.D{{"priority", typeVal}, {"company_code", CompanyCode}}
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
