package jwt

import (
	"fmt"
	"github.com/Gentostage/golang-auth/internal/app/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestAccessToken_Decode(t *testing.T) {
	u := &model.User{
		ID:       primitive.NewObjectID(),
		Email:    "Example",
		Password: "12312321",
	}
	token := &AccessToken{
		SecretKey:  "wqeqwesadsadasxdsadsadsadas",
		TimeToLive: 30,
	}
	hash, err := token.Encode(u)
	fmt.Println(hash)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
}

func TestAccessToken_Encode(t *testing.T) {
	u := &model.User{
		ID:       primitive.NewObjectID(),
		Email:    "Example",
		Password: "12312321",
	}
	token := &AccessToken{
		SecretKey:  "wqeqwesadsadasxdsadsadsadas",
		TimeToLive: 30,
	}
	hash, err := token.Encode(u)
	assert.NoError(t, err)
	tokenData, err := token.Decode(hash)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenData)
	fmt.Println(tokenData)

}
