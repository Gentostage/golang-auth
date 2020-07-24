package model

import (
	"crypto/sha256"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Token struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	RefreshToken string             `bson:"refresh_token,omitempty"`
	RegisterTime time.Time          `bson:"time_to_live,omitempty"`
	Alive        bool               `bson:"alive,omitempty"`
	UserId       primitive.ObjectID `bson:"user_id"`
}

func (t *Token) GenerateHashToken(accsessToken string) error {
	token := accsessToken + t.RefreshToken
	tokenSha := sha256.New()
	tokenSha.Write([]byte(token))
	hashToken, err := bcrypt.GenerateFromPassword(tokenSha.Sum(nil), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	t.RefreshToken = string(hashToken)
	return nil
}

func (t *Token) CompareTokens(refreshToken string, accsessToken string) error {
	token := accsessToken + refreshToken
	tokenSha := sha256.New()
	tokenSha.Write([]byte(token))

	err := bcrypt.CompareHashAndPassword([]byte(t.RefreshToken), tokenSha.Sum(nil))
	if err != nil {
		return err
	}
	return nil
}
