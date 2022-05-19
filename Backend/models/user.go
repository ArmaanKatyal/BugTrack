package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User struct holds the user data that is stored in the database
type User struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string             `json:"first_name" bson:"first_name"`
	LastName  string             `json:"last_name" bson:"last_name"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	Email     string             `json:"email" bson:"email"`
	Role      string             `json:"role" bson:"role"`
	CreatedOn primitive.DateTime `json:"created_on" bson:"created_on"`
}

// CreateUser structs holds the user data when creating a new user
type CreateUser struct {
	FirstName string             `json:"first_name" bson:"first_name"`
	LastName  string             `json:"last_name" bson:"last_name"`
	Username  string             `json:"username" bson:"username"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	Role      string             `json:"role" bson:"role"`
	CreatedOn primitive.DateTime `json:"created_on" bson:"created_on"`
}

// Profile struct holds the user profile data that is stored in the database
type Profile struct {
	FirstName       string   `json:"first_name" bson:"first_name"`
	LastName        string   `json:"last_name" bson:"last_name"`
	Username        string   `json:"username" bson:"username"`
	Email           string   `json:"email" bson:"email"`
	Role            string   `json:"role" bson:"role"`
	TicketsCreated  []Ticket `json:"tickets_created" bson:"tickets_created"`
	TicketsAssigned []Ticket `json:"tickets_assigned" bson:"tickets_assigned"`
}

// Credentials holds the user credentials
type Credentials struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
