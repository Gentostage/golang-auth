package apiserver

import (
	"fmt"
	"github.com/Gentostage/golang-auth/internal/app/jwt"
	"github.com/Gentostage/golang-auth/internal/app/model"
	"github.com/Gentostage/golang-auth/internal/app/store"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type server struct {
	router *gin.Engine
	logger *logrus.Logger
	store  store.Store
	access *jwt.AccessToken
}

func newServer(config Config, store store.Store) *server {
	logger := logrus.New()
	loggerLevel, _ := logrus.ParseLevel(config.LogLevel)
	logger.SetLevel(loggerLevel)
	access := &jwt.AccessToken{
		SecretKey:  "vdgjfesbf tc,jug,jutkufr,jf,juf,f,f,uj f",
		TimeToLive: 20,
	}
	s := &server{
		router: gin.New(),
		logger: logger,
		store:  store,
		access: access,
	}
	s.router.Use(s.Middle())

	s.configureRoute()
	return s
}
func (s *server) Middle() gin.HandlerFunc {
	return func(context *gin.Context) {
		token, _ := context.Cookie("Access_Token")
		tokenHeader := context.Request.Header.Get("X-Auth-Token")

		if tokenHeader != "" {
			token = tokenHeader
		}
		if token != "" {
			tokenStruct, err := s.access.Decode(token)
			if err != nil {
				context.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			context.Set("user_id", tokenStruct.Payload.User)
			fmt.Println(tokenStruct)
		}
	}

}

func (s *server) configureRoute() {
	s.router.GET("/user/", func(context *gin.Context) {
		u := &model.User{}
		userId, exist := context.Get("user_id")

		if !exist {
			context.String(http.StatusUnauthorized, "Ошибка доступа")
		}

		u.ID = userId.(primitive.ObjectID)
		user, err := s.store.User().Get(u)
		if err != nil {
			s.logger.Info(err)
			context.String(http.StatusNotFound, err.Error())
		} else {
			context.JSON(http.StatusOK, user)
		}

	})
	s.router.POST("/user/create", func(context *gin.Context) {
		user := &model.User{}
		if err := context.BindJSON(&user); err != nil {
			context.String(http.StatusBadRequest, err.Error())
		}
		userTemp := &model.User{
			Email: user.Email,
		}
		u, _ := s.store.User().Get(userTemp)
		if u == nil {
			err := user.Validate()
			if err == nil {
				err = user.GeneratePassword()
				if err == nil {
					err = s.store.User().Create(user)
					if err == nil {
						context.JSON(http.StatusOK, user)
						return
					}
				}
			}
			s.logger.Error(err)
			context.String(http.StatusBadRequest, err.Error())

		} else {
			context.String(http.StatusBadRequest, "Пользователь с таким email уже существует")
		}

	})
	s.router.POST("/user/login", func(context *gin.Context) {
		user := &model.User{}
		if err := context.BindJSON(&user); err != nil {
			context.String(http.StatusBadRequest, err.Error())
		}
		tempUser := &model.User{
			Email: user.Email,
		}
		u, err := s.store.User().Get(tempUser)
		if u != nil {
			if u.ComparePassword(user.Password) {
				token, err := s.access.Encode(u)
				if err != nil {
					context.String(http.StatusInternalServerError, err.Error())
					return
				}
				context.SetCookie("Access_Token", token, 3600, "/", "127.0.0.1", false, true)
				context.JSON(http.StatusOK, struct {
					AccessToken string `json:"access_token"`
				}{
					token,
				})
				return
			}
		}
		s.logger.Error(err)
		context.String(http.StatusBadRequest, "Неверно имя пользователя или пароль")
	})

}
