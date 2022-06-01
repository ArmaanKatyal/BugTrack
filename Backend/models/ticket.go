package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Ticket struct holds the ticket data that is stored in the database
type Ticket struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Status      string             `json:"status" bson:"status"`
	Priority    string             `json:"priority" bson:"priority"`
	Tags        []string           `json:"tags" bson:"tags"`
	CreatedBy   string             `json:"created_by" bson:"created_by"`
	AssignedTo  []string           `json:"assigned_to" bson:"assigned_to"`
	CreatedOn   primitive.DateTime `json:"created_on" bson:"created_on"`
	UpdatedOn   primitive.DateTime `json:"updated_on" bson:"updated_on"`
	ProjectId   primitive.ObjectID `json:"project_id" bson:"project_id"`
	ProjectName string             `json:"project_name" bson:"project_name"`
	CompanyCode string             `json:"company_code" bson:"company_code"`
}

// CreateTicket structs holds the ticket data when creating a new ticket
type CreateTicket struct {
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Status      string             `json:"status" bson:"status"`
	Priority    string             `json:"priority" bson:"priority"`
	Tags        []string           `json:"tags" bson:"tags"`
	CreatedBy   string             `json:"created_by" bson:"created_by"`
	AssignedTo  []string           `json:"assigned_to" bson:"assigned_to"`
	CreatedOn   primitive.DateTime `json:"created_on" bson:"created_on, omitempty"`
	UpdatedOn   primitive.DateTime `json:"updated_on" bson:"updated_on"`
	ProjectId   primitive.ObjectID `json:"project_id" bson:"project_id"`
	ProjectName string             `json:"project_name" bson:"project_name"`
	CompanyCode string             `json:"company_code" bson:"company_code"`
}

type CreateTicket2 struct {
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Priority    string             `json:"priority" bson:"priority"`
	ProjectId   primitive.ObjectID `json:"project_id" bson:"project_id"`
	ProjectName string             `json:"project_name" bson:"project_name"`
	AssignedTo  []string           `json:"assigned_to" bson:"assigned_to"`
	Tags        []string           `json:"tags" bson:"tags"`
}

type CreateTicket3 struct {
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Priority    string             `json:"priority" bson:"priority"`
	ProjectId   primitive.ObjectID `json:"project_id" bson:"project_id"`
	ProjectName string             `json:"project_name" bson:"project_name"`
	AssignedTo  []string           `json:"assigned_to" bson:"assigned_to"`
	Tags        []string           `json:"tags" bson:"tags"`
	Status      string             `json:"status" bson:"status"`
}
