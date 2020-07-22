package jwt

import (
	"crypto/hmac"
	"crypto/sha512"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/Gentostage/golang-auth/internal/app/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
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

func (t *AccessToken) Decode(token string) (TokenStructData, error) {
	arrayString := strings.Split(token, ".")
	header := arrayString[0]
	payload := arrayString[1]

	tokenHeader := &TokenStructData{}

	headerByte, err := b64.URLEncoding.DecodeString(header)
	if err != nil {
		return TokenStructData{}, err
	}
	err = json.Unmarshal(headerByte, &tokenHeader.Header)
	if err != nil {
		return TokenStructData{}, err
	}
	payloadByte, err := b64.URLEncoding.DecodeString(payload)
	if err != nil {
		return TokenStructData{}, err
	}
	err = json.Unmarshal(payloadByte, &tokenHeader.Payload)
	if err != nil {
		return TokenStructData{}, err
	}
	hash, err := t.generateHash(tokenHeader)
	if hash != token {
		return TokenStructData{}, errors.New("Token not valide")
	}
	return *tokenHeader, nil
}

func (t *AccessToken) generateHash(token *TokenStructData) (string, error) {

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

	return b64Header + "." + b64Payload + "." + b64Signature, nil
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
	hash, err := t.generateHash(token)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (t *AccessToken) Update() {

}
