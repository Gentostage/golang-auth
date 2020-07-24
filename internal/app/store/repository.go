package store

import "github.com/Gentostage/golang-auth/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	Get(*model.User) (*model.User, error)
}

type TokenRepository interface {
	Create(tokenRefresh *model.Token) error
	Get(token *model.Token) (*model.Token, error)
	GetAllAliveTokensByUser(token *model.Token) ([]*model.Token, error)
	Close(token *model.Token) error
	DeleteAll(token *model.Token) error
	DeleteOne(token *model.Token) error
}

type InvalidTokenRepository interface {
	Find(token string) error
	Insert(token string) error
}
