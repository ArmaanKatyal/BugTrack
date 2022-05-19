package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Log struct holds the log data that is stored in the database
type Log struct {
	Type        string             `json:"type" bson:"type"`
	Author      string             `json:"author" bson:"author"`
	Date        primitive.DateTime `json:"date" bson:"date"`
	Description string             `json:"description" bson:"description"`
	Table       string             `json:"table" bson:"table"`
}
