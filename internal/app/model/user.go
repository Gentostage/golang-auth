package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty"bson:"_id,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty"`
	LastName  string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	FirstName string             `json:"first_name,omitempty" bson:"first_name,omitempty"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(requiredIf(u.Password == "")), validation.Length(6, 100)),
	)
}

func (u *User) GeneratePassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	} else {
		u.Password = string(hash)
	}
	return nil
}
