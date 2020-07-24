package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InvalidToken struct {
	ID    primitive.ObjectID `bson:"_id"`
	Token string             `bson:"token"`
}
