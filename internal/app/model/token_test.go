package model

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestToken_CompareTokens(t *testing.T) {
	token := &struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}{
		AccessToken:  "abv",
		RefreshToken: "123",
	}
	tokenM := &Token{
		ID:           primitive.ObjectID{},
		RefreshToken: "",
		RegisterTime: time.Time{},
		Alive:        false,
		UserId:       primitive.ObjectID{},
	}
	tokenM.RefreshToken = token.RefreshToken
	err := tokenM.GenerateHashToken(token.AccessToken)
	assert.NoError(t, err)

	err = tokenM.CompareTokens(token.RefreshToken, token.AccessToken)
	assert.NoError(t, err)

	token.RefreshToken = "321h5jk 4ih5fy3445$&^%*&%*&%(#B Y$*( #&(*"
	token.AccessToken = "PidarasPid^B&(36578936986809609q6w3098PidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidarasPidaras"
	err = tokenM.CompareTokens(token.RefreshToken, token.AccessToken)
	assert.Error(t, err)

	tokenM.RefreshToken = "$2a$10$7lI/WxCH4uceXqcXZ2oPQ..3pRR1xmUz3VdyrLocQToK09XHS2fdy"
	err = tokenM.CompareTokens(token.RefreshToken, token.AccessToken)
	assert.Error(t, err)

	token.RefreshToken = "zBtmIDa3ZxD8cR4FoCGrCH98tLw="
	token.AccessToken = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiNWYxN2U1YWE3OTc5MjJjOTJlMjBjM2RlIiwidGltZSI6IjIwMjAtMDctMjRUMTI6MzM6NTEuMzM0ODA4MSswOTowMCJ9.Mzg1OThjYjVjY2UyYWJmY2MwY2I3ZTkzM2ViYjY1MGZlOGM4NDU3MTEzNTRiZGNhOGQwNDQzYmIxNmQ2OTYzMmU3NzFmODcwYmExNjBlNGEzNGE4NzllNmFhYjQ1Y2U3MzgzZTI3OTQ1MGI4ZWUwYzFjODZlYzU4ODllNmM4ZTc="
	err = tokenM.CompareTokens(token.RefreshToken, token.AccessToken)
	assert.Error(t, err)

}
