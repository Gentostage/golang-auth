package jwt

import (
	"fmt"
	"github.com/Gentostage/golang-auth/internal/app/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
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

func TestAccessToken_liveTime(t *testing.T) {
	token := &AccessToken{
		SecretKey:  "wqeqwesadsadasxdsadsadsadas",
		TimeToLive: 10,
	}
	oldTime := time.Now().Add(-11 * time.Minute)
	tokenStruct := &TokenStructData{
		Header: struct {
			Alg string `json:"alg"`
			Typ string `json:"typ"`
		}{},
		Payload: struct {
			User primitive.ObjectID `json:"user"`
			Time time.Time          `json:"time"`
		}{Time: oldTime},
	}
	err := token.liveTime(tokenStruct)
	assert.Error(t, err)

	newTime := time.Now().Add(time.Minute * 5)
	tokenStruct.Payload.Time = newTime
	err2 := token.liveTime(tokenStruct)
	assert.NoError(t, err2)
}

func TestRefreshToken_Generate(t *testing.T) {
	u := &model.User{
		ID:       primitive.NewObjectID(),
		Email:    "Example",
		Password: "12312321",
	}
	token := RefreshToken{}
	hash, _ := token.Generate(u)
	fmt.Println(hash)
	assert.NotEmpty(t, hash)
}
