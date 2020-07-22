package jwt

import (
	"crypto/hmac"
	"crypto/sha512"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"github.com/Gentostage/golang-auth/internal/app/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AccessToken struct {
	SecretKey  string
	TimeToLive int
}

type TokenStructData struct {
	Header struct {
		Alg string `json:"alg"`
		Typ string `json:"typ"`
	}
	Payload struct {
		User primitive.ObjectID `json:"user"`
		Time time.Time          `json:"time"`
	}
}

type RefreshToken struct {
}

func (t *RefreshToken) Refresh() {

}

func (t *RefreshToken) Generate() {

}

func (t *AccessToken) Decode() {

}

func (t *AccessToken) Encode(user *model.User) (string, error) {
	token := &TokenStructData{
		Header: struct {
			Alg string `json:"alg"`
			Typ string `json:"typ"`
		}{Alg: "HS512", Typ: "JWT"},
		Payload: struct {
			User primitive.ObjectID `json:"user"`
			Time time.Time          `json:"time"`
		}{User: user.ID, Time: time.Now()},
	}

	jsonHeader, err := json.Marshal(token.Header)
	if err != nil {
		return "", err
	}
	jsonPayload, err := json.Marshal(token.Payload)
	if err != nil {
		return "", err
	}

	b64Header := b64.URLEncoding.EncodeToString(jsonHeader)
	b64Payload := b64.URLEncoding.EncodeToString(jsonPayload)

	h := hmac.New(sha512.New, []byte(t.SecretKey))
	h.Write([]byte(b64Header + "." + b64Payload))
	hexSignature := hex.EncodeToString(h.Sum(nil))

	b64Signature := b64.URLEncoding.EncodeToString([]byte(hexSignature))

	return b64Header + "." + b64Payload + "." + string(b64Signature), nil
}

func (t *AccessToken) Update() {

}
