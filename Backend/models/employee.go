package models

type Employee struct {
	Id        int64  `json:"id" bson:"_id"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email" bson:"email"`
	CreatedOn string `json:"created_on" bson:"created_on"`
	UpdatedOn string `json:"updated_on" bson:"updated_on"`
}
