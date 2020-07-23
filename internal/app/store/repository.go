package store

import "github.com/Gentostage/golang-auth/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	Get(*model.User) (*model.User, error)
}

type TokenRepository interface {
	Create(tokenRefresh *model.Token) error
	Get(token *model.Token) (*model.Token, error)
	Close() error
	DeleteAll() error
	DeleteOne() error
}
