package mongostore

import (
	"fmt"
	"github.com/Gentostage/golang-auth/internal/app/jwt"
	"github.com/Gentostage/golang-auth/internal/app/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"runtime"
	"testing"
	"time"
)

func TestOneBullion(t *testing.T) {
	store, _ := NewBD("mongodb://localhost:27017")
	numbers := [999999]int{}
	runtime.GOMAXPROCS(5)
	for i := range numbers {
		fmt.Println(i)
		go func() {
			u := &model.User{
				ID:        primitive.NewObjectID(),
				Email:     "sasai@gmail.com",
				Password:  "",
				LastName:  "",
				FirstName: "",
			}
			r := jwt.RefreshToken{TimeToLive: 1}
			refresh, _ := r.Generate(u)

			ac := jwt.AccessToken{
				SecretKey:  "sadasdasd",
				TimeToLive: 12,
			}
			access, err2 := ac.Encode(u)

			if err2 != nil {
				log.Println(err2)

			}
			token := &model.Token{
				ID:                 primitive.NewObjectID(),
				AccessRefreshToken: refresh,
				RefreshToken:       "",
				RegisterTime:       time.Time{},
				Alive:              false,
				UserId:             primitive.ObjectID{},
			}

			fmt.Println(refresh)

			err3 := token.GenerateHashToken(access)
			if err3 != nil {
				log.Println(err3)
			}

			err4 := store.Token().Create(token)
			if err4 != nil {
				log.Println(err4)
			}

		}()
	}

}
