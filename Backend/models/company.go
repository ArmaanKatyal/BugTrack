package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Company struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	CompanyCode string             `json:"company_code" bson:"company_code"`
}

type CreateCompany struct {
	Name        string `json:"name" bson:"name"`
	CompanyCode string `json:"company_code" bson:"company_code"`
}
