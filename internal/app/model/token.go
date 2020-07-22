package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Token struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	RefreshToken byte               `bson:"refresh_token,omitempty"`
	Alive        bool               `bson:"alive, omitempty"`
}
