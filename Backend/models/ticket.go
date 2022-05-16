package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ticket struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Status      string             `json:"status" bson:"status"`
	Priority    string             `json:"priority" bson:"priority"`
	Tags        []string           `json:"tags" bson:"tags"`
	CreatedBy   string             `json:"created_by" bson:"created_by"`
	AssignedTo  string             `json:"assigned_to" bson:"assigned_to"`
	CreatedOn   primitive.DateTime `json:"created_on" bson:"created_on"`
	UpdatedOn   primitive.DateTime `json:"updated_on" bson:"updated_on"`
	ProjectId   primitive.ObjectID `json:"project_id" bson:"project_id"`
}
