package apiserver

import (
	"github.com/Gentostage/golang-auth/internal/app/store/mongostore"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
	}
}

func (s *APIServer) Start() error {
	store, err := mongostore.NewBD()
	if err != nil {
		return err
	}
	srv := newServer(*s.config, store)
	err = srv.router.Run(s.config.BindAddr)

	return err
}
