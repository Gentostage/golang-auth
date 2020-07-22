package store

import "github.com/Gentostage/golang-auth/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	Get(*model.User) (*model.User, error)
}

type TokenRepository interface {
	Create() error
	Close() error
	DeleteAll() error
}
