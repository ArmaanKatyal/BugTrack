package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User struct holds the user data that is stored in the database
type User struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	FirstName   string             `json:"first_name" bson:"first_name"`
	LastName    string             `json:"last_name" bson:"last_name"`
	Username    string             `json:"username" bson:"username"`
	Email       string             `json:"email" bson:"email"`
	Role        string             `json:"role" bson:"role"`
	CreatedOn   primitive.DateTime `json:"created_on" bson:"created_on"`
	CompanyCode string             `json:"company_code" bson:"company_code"`
	Locked      bool               `json:"locked" bson:"locked"`
}

// CreateUser structs holds the user data when creating a new user
type CreateUser struct {
	FirstName   string             `json:"first_name" bson:"first_name"`
	LastName    string             `json:"last_name" bson:"last_name"`
	Username    string             `json:"username" bson:"username"`
	Email       string             `json:"email" bson:"email"`
	Role        string             `json:"role" bson:"role"`
	CreatedOn   primitive.DateTime `json:"created_on" bson:"created_on"`
	CompanyCode string             `json:"company_code" bson:"company_code"`
	Locked      bool               `json:"locked" bson:"locked"`
}

type UpdateUser struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email" bson:"email"`
	Role      string `json:"role" bson:"role"`
}

type UserDatafromAdmin struct {
	FirstName   string             `json:"first_name" bson:"first_name"`
	LastName    string             `json:"last_name" bson:"last_name"`
	Username    string             `json:"username" bson:"username"`
	Password    string             `json:"password" bson:"password"`
	Email       string             `json:"email" bson:"email"`
	Role        string             `json:"role" bson:"role"`
	CreatedOn   primitive.DateTime `json:"created_on" bson:"created_on"`
	CompanyCode string             `json:"company_code" bson:"company_code"`
	Locked      bool               `json:"locked" bson:"locked"`
}

type UserDatafromAdmin2 struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Username  string `json:"username" bson:"username"`
	Password  string `json:"password" bson:"password"`
	Email     string `json:"email" bson:"email"`
	Role      string `json:"role" bson:"role"`
	Locked    bool   `json:"locked" bson:"locked"`
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
	CompanyCode     string   `json:"company_code" bson:"company_code"`
}

// Credentials holds the user credentials
type Credentials struct {
	Username    string `json:"username" bson:"username"`
	Password    string `json:"password" bson:"password"`
	CompanyCode string `json:"company_code" bson:"company_code"`
}

type DatabaseCreds struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Username    string             `json:"username" bson:"username"`
	Password    string             `json:"password" bson:"password"`
	CompanyCode string             `json:"company_code" bson:"company_code"`
}
