package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	Id          primitive.ObjectID `json:"id" bson:"_id, omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description, omitempty"`
	CreatedBy   []string           `json:"created_by" bson:"created_by"`
	CreatedOn   primitive.DateTime `json:"created_on" bson:"created_on"`
}

type CreateProject struct {
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	CreatedBy   []string           `json:"created_by" bson:"created_by"`
	CreatedOn   primitive.DateTime `json:"created_on" bson:"created_on"`
}
