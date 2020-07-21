package apiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type server struct {
	router *gin.Engine
	logger *logrus.Logger
}

func newServer(config *Config) *server {
	logger := logrus.New()
	loggerLevel, _ := logrus.ParseLevel(config.LogLevel)
	logger.SetLevel(loggerLevel)

	s := &server{
		router: gin.New(),
		logger: logger,
	}
	s.configureRoute()
	return s
}

func (s *server) configureRoute() {
	s.router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Ok")
	})
}
