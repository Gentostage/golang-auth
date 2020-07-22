package mongostore

type TokenRepository struct {
	store *Store
}

func (t *TokenRepository) Create() error {
	return nil
}

func (t *TokenRepository) Close() error {
	return nil
}

func (t *TokenRepository) DeleteAll() error {
	return nil
}
