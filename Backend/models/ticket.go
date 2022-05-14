package models

type Ticket struct {
	Id          int64  `json:"id" bson:"_id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Status      string `json:"status" bson:"status"`
	Priority    string `json:"priority" bson:"priority"`
	CreatedBy   string `json:"created_by" bson:"created_by"`
	AssignedTo  string `json:"assigned_to" bson:"assigned_to"`
	CreatedOn   string `json:"created_on" bson:"created_on"`
	UpdatedOn   string `json:"updated_on" bson:"updated_on"`
}
