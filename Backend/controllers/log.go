package controllers

import (
	"Backend/config"
	"Backend/models"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

// AllLogs returns all the logs
func AllLogs(w http.ResponseWriter, r *http.Request) {
	// if the request is not GET method, return 405
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// check if the user is logged in
	_, CompanyCode, _, err := authenticate(r)
	// if the user is not logged in, return 401
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// create the client connection
	client := config.ClientConnection()
	// get the collection
	coll := client.Database(config.ViperEnvVariable("dbName")).Collection("logs")
	// slice to store the logs
	var logs []models.Log
	// find all the logs
	cursor, err := coll.Find(context.TODO(), bson.D{{"company_code", CompanyCode}})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// iterate through the logs
	if err = cursor.All(context.TODO(), &logs); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer func() {
		// close the cursor
		if err := cursor.Close(context.TODO()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// close the client connection
		if err := client.Disconnect(context.TODO()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()
	// reverse the slice to get the latest logs first
	reverseLogs := make([]models.Log, len(logs))
	for i, log := range logs {
		reverseLogs[len(logs)-i-1] = log
	}

	// set the content type to json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)                 // 200
	err = json.NewEncoder(w).Encode(reverseLogs) // encode the logs
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
