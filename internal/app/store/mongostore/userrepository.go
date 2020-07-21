package mongostore

import (
	"context"
	"github.com/Gentostage/golang-auth/internal/app/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	userStorage := r.store.db.Database("auth-go").Collection("users")
	_, err := userStorage.InsertOne(context.TODO(), u)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Get(u *model.User) (*model.User, error) {
	user := &model.User{}
	userStorage := r.store.db.Database("auth-go").Collection("users")
	err := userStorage.FindOne(context.TODO(), u).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil

}
