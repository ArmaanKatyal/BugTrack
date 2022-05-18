package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string             `json:"first_name" bson:"first_name"`
	LastName  string             `json:"last_name" bson:"last_name"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	Email     string             `json:"email" bson:"email"`
	CreatedOn primitive.DateTime `json:"created_on" bson:"created_on"`
}

type CreateUser struct {
	FirstName string             `json:"first_name" bson:"first_name"`
	LastName  string             `json:"last_name" bson:"last_name"`
	Username  string             `json:"username" bson:"username"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	CreatedOn primitive.DateTime `json:"created_on" bson:"created_on"`
}

type Profile struct {
	FirstName       string   `json:"first_name" bson:"first_name"`
	LastName        string   `json:"last_name" bson:"last_name"`
	Username        string   `json:"username" bson:"username"`
	Email           string   `json:"email" bson:"email"`
	TicketsCreated  []Ticket `json:"tickets_created" bson:"tickets_created"`
	TicketsAssigned []Ticket `json:"tickets_assigned" bson:"tickets_assigned"`
}

type Credentials struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
