package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email     string             `json:"email" bson:"email,omitempty"`
	Password  string             `json:"password" bson:"password,omitempty"`
	LastName  string             `json:"last_name" bson:"last_name,omitempty"`
	FirstName string             `json:"first_name" bson:"first_name,omitempty"`
}
