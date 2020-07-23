package mongostore

import (
	"context"
	"github.com/Gentostage/golang-auth/internal/app/model"
)

type TokenRepository struct {
	store *Store
}

func (t *TokenRepository) Create(tokenRefresh *model.Token) error {
	tokenStorage := t.store.db.Database("auth-go").Collection("tokens")
	_, err := tokenStorage.InsertOne(context.TODO(), &tokenRefresh)
	if err != nil {
		return err
	}
	return nil
}

func (t *TokenRepository) Get(token *model.Token) (*model.Token, error) {
	tokenBase := &model.Token{}
	tokenStorage := t.store.db.Database("auth-go").Collection("tokens")
	err := tokenStorage.FindOne(context.TODO(), token).Decode(tokenBase)
	if err != nil {
		return nil, err
	}
	return tokenBase, nil
}

func (t *TokenRepository) Close() error {
	return nil
}

func (t *TokenRepository) DeleteAll() error {
	return nil
}

func (t *TokenRepository) DeleteOne() error {
	return nil
}
