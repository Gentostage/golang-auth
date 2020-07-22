package store

type Store interface {
	User() UserRepository
	Token() TokenRepository
}
