package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Project struct holds the project data that is stored in the database
type Project struct {
	Id          primitive.ObjectID `json:"id" bson:"_id, omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description, omitempty"`
	CreatedBy   string             `json:"created_by" bson:"created_by"`
	CreatedOn   primitive.DateTime `json:"created_on" bson:"created_on"`
	AssignedTo  []string           `json:"assigned_to" bson:"assigned_to"`
	CompanyCode string             `json:"company_code" bson:"company_code"`
}

// CreateProject structs holds the project data when creating a new project
type CreateProject struct {
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	CreatedBy   string             `json:"created_by" bson:"created_by"`
	CreatedOn   primitive.DateTime `json:"created_on" bson:"created_on"`
	AssignedTo  []string           `json:"assigned_to" bson:"assigned_to"`
	CompanyCode string             `json:"company_code" bson:"company_code"`
}

type CreateProject2 struct {
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	//CreatedBy   string `json:"created_by" bson:"created_by"`
	//CreatedOn   primitive.DateTime `json:"created_on" bson:"created_on"`
	//AssignedTo  []string `json:"assigned_to" bson:"assigned_to"`
	//CompanyCode string   `json:"company_code" bson:"company_code"`
}
