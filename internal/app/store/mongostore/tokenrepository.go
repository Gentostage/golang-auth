package mongostore

import (
	"context"
	"encoding/json"
	"github.com/Gentostage/golang-auth/internal/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenRepository struct {
	store      *Store
	collection *mongo.Collection
}

func (t *TokenRepository) Create(tokenRefresh *model.Token) error {
	tokenStorage := t.collection
	_, err := tokenStorage.InsertOne(context.Background(), &tokenRefresh)
	if err != nil {
		return err
	}
	return nil
}

func (t *TokenRepository) Get(token *model.Token) (*model.Token, error) {
	tokenBase := &model.Token{}
	tokenStorage := t.collection
	err := tokenStorage.FindOne(context.TODO(), token).Decode(tokenBase)
	if err != nil {
		return nil, err
	}
	return tokenBase, nil
}

func (t *TokenRepository) GetAllAliveTokensByUser(token *model.Token) ([]*model.Token, error) {
	var doc []*model.Token
	tokenStorage := t.collection
	tokenResult, err := tokenStorage.Find(context.TODO(), token)
	if err != nil {
		return nil, err
	}
	if tokenResult != nil {
		for tokenResult.Next(context.TODO()) {
			var el model.Token
			err = tokenResult.Decode(&el)
			if err != nil {
				return nil, err
			}
			doc = append(doc, &el)
		}
	}

	return doc, nil
}

func (t *TokenRepository) GetAllTokens() ([]*model.Token, error) {
	var doc []*model.Token
	tokenStorage := t.collection
	tokenResult, err := tokenStorage.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	if tokenResult != nil {
		for tokenResult.Next(context.TODO()) {
			var el model.Token
			err = tokenResult.Decode(&el)
			if err != nil {
				return nil, err
			}
			doc = append(doc, &el)
		}
	}

	return doc, nil
}

func (t *TokenRepository) Close(token *model.Token) error {
	update := bson.M{}
	_ = json.Unmarshal([]byte(`{ "$set": {"alive": false}}`), &update)
	if _, err := t.collection.UpdateOne(context.TODO(), bson.M{"_id": token.ID}, update); err != nil {
		return err
	}
	return nil
}

func (t *TokenRepository) DeleteAll(token *model.Token) error {
	if _, err := t.collection.DeleteMany(context.TODO(), bson.M{"user_id": token.UserId}); err != nil {
		return err
	}
	return nil
}

func (t *TokenRepository) DeleteOne(token *model.Token) error {
	if _, err := t.collection.DeleteOne(context.TODO(), token); err != nil {
		return err
	}
	return nil
}
