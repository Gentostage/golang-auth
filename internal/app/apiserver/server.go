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
		u := s.store.User()
		user := &model.User{
			Email: "example",
		}
		err := u.Get(user)
		if err != nil {
			s.logger.Error(err)
		}

		context.String(http.StatusOK, "Email."+user.Email)
	})
	s.router.GET("/user/create", func(context *gin.Context) {
		u := s.store.User()
		user := &model.User{
			Email: "example",
		}
		err := u.Create(user)
		if err != nil {
			s.logger.Error(err)
		}
		context.String(http.StatusOK, "Email:"+user.Email)
	})
}
