package apiserver

import "github.com/sirupsen/logrus"

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
	srv := newServer(s.config)
	err := srv.router.Run(s.config.BindAddr)
	if err != nil {
		s.logger.Error(err)
	}
	return nil
}
