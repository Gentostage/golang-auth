package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Token struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	RefreshToken string             `bson:"refresh_token,omitempty"`
	TimeToLive   time.Time          `bson:"time_to_live,omitempty"`
	Alive        bool               `bson:"alive,omitempty"`
}

func (t *Token) GenerateHashToken(accsessToken string) error {
	token := accsessToken + t.RefreshToken
	hashToken, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	t.RefreshToken = string(hashToken)
	return nil
}

func (t *Token) CompareTokens(refreshToken string, accsessToken string) error {
	token := accsessToken + refreshToken
	err := bcrypt.CompareHashAndPassword([]byte(t.RefreshToken), []byte(token))
	if err != nil {
		return err
	}
	return nil
}
