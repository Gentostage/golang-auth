package model

import (
	"crypto/sha256"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Token struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	AccessRefreshToken string             `bson:"access_refresh_token,omitempty"`
	RefreshToken       string             `bson:"refresh_token,omitempty"`
	RegisterTime       time.Time          `bson:"time_to_live,omitempty"`
	Alive              bool               `bson:"alive,omitempty"`
	UserId             primitive.ObjectID `bson:"user_id"`
}

func (t *Token) GenerateHashToken(accessToken string) error {
	token := accessToken + t.AccessRefreshToken
	tokenSha := sha256.New()
	tokenSha.Write([]byte(token))
	hashToken, err := bcrypt.GenerateFromPassword(tokenSha.Sum(nil), bcrypt.DefaultCost)
	hashRefreshToken, err := bcrypt.GenerateFromPassword([]byte(t.AccessRefreshToken), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	t.RefreshToken = string(hashRefreshToken)
	t.AccessRefreshToken = string(hashToken)
	return nil
}

func (t *Token) CompareTokens(refreshToken string, accessToken string) error {
	token := accessToken + refreshToken
	tokenSha := sha256.New()
	tokenSha.Write([]byte(token))

	err := bcrypt.CompareHashAndPassword([]byte(t.AccessRefreshToken), tokenSha.Sum(nil))
	if err != nil {
		return err
	}
	return nil
}

func (t *Token) CompareRefreshToken(refreshToken string) error {
	err := bcrypt.CompareHashAndPassword([]byte(t.RefreshToken), []byte(refreshToken))
	if err != nil {
		return err
	}
	return nil
}
