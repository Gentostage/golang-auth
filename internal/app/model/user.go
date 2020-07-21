package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty"bson:"_id,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty"`
	LastName  string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	FirstName string             `json:"first_name,omitempty" bson:"first_name,omitempty"`
}
