package apiserver

import (
	"github.com/Gentostage/golang-auth/internal/app/model"
	"github.com/Gentostage/golang-auth/internal/app/store"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type server struct {
	router *gin.Engine
	logger *logrus.Logger
	store  store.Store
}

func newServer(config Config, store store.Store) *server {
	logger := logrus.New()
	loggerLevel, _ := logrus.ParseLevel(config.LogLevel)
	logger.SetLevel(loggerLevel)
	s := &server{
		router: gin.New(),
		logger: logger,
		store:  store,
	}
	s.configureRoute()
	return s
}

func (s *server) configureRoute() {
	s.router.GET("/user/", func(context *gin.Context) {
		u := &model.User{}
		if err := context.BindJSON(&u); err != nil {
			context.String(http.StatusBadRequest, err.Error())
		}
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
		user_temp := &model.User{
			Email: user.Email,
		}
		u, _ := s.store.User().Get(user_temp)
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
}
