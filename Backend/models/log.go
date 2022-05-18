package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Log struct {
	Type   string             `json:"type" bson:"type"`
	Author string             `json:"author" bson:"author"`
	Date   primitive.DateTime `json:"date" bson:"date"`
	Table  string             `json:"table" bson:"table"`
}
