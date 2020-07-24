package mongostore

import (
	"context"
	"errors"
	"github.com/Gentostage/golang-auth/internal/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type InvalidTokenRepository struct {
	store      *Store
	collection *mongo.Collection
}

func (i *InvalidTokenRepository) Insert(token string) error {
	iToken := &model.InvalidToken{
		ID:    primitive.NewObjectID(),
		Token: token,
	}
	storage := i.collection
	_, err := storage.InsertOne(context.TODO(), iToken)
	if err != nil {
		return err
	}
	return nil
}

func (i *InvalidTokenRepository) Find(token string) error {
	iToken := &model.InvalidToken{}
	storage := i.collection
	find := bson.M{"token": token}
	_ = storage.FindOne(context.TODO(), find).Decode(&iToken)
	if iToken.Token == token {
		return errors.New("Token is exist ")
	}
	return nil
}
