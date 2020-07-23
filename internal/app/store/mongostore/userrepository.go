package mongostore

import (
	"context"
	"github.com/Gentostage/golang-auth/internal/app/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	store      *Store
	collection *mongo.Collection
}

func (r *UserRepository) Create(u *model.User) error {
	u.ID = primitive.NewObjectID()
	userStorage := r.collection
	_, err := userStorage.InsertOne(context.TODO(), u)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Get(u *model.User) (*model.User, error) {
	user := &model.User{}
	userStorage := r.collection
	err := userStorage.FindOne(context.TODO(), u).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil

}
