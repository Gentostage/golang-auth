package mongostore

import (
	"context"
	"github.com/Gentostage/golang-auth/internal/app/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	db              *mongo.Client
	userRepository  *UserRepository
	tokenRepository *TokenRepository
}

func NewBD(dataBaseUrl string) (*Store, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(dataBaseUrl))
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: client,
	}, nil
}

func (s *Store) CloseDB() {
	_ = s.db.Disconnect(context.TODO())
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store:      s,
		collection: s.db.Database("auth-go").Collection("users"),
	}
	return s.userRepository
}

func (s *Store) Token() store.TokenRepository {
	if s.tokenRepository != nil {
		return s.tokenRepository
	}
	s.tokenRepository = &TokenRepository{
		store:      s,
		collection: s.db.Database("auth-go").Collection("tokens"),
	}
	return s.tokenRepository
}
