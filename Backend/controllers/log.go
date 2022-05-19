package controllers

import (
	"Backend/config"
	"Backend/models"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func AllLogs(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_, err := authenticate(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	client := config.ClientConnection()
	coll := client.Database("bugTrack").Collection("logs")
	var logs []models.Log
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = cursor.All(context.TODO(), &logs); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer func() {
		if err := cursor.Close(context.TODO()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := client.Disconnect(context.TODO()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()

	reverseLogs := make([]models.Log, len(logs))
	for i, log := range logs {
		reverseLogs[len(logs)-i-1] = log
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(reverseLogs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
