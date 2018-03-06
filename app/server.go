package app

import (
	"github.com/nephele/logger"
	"github.com/nephele/service"
)

type InitializeFunc func(*Server) error

type Server struct {
	init    InitializeFunc
	logger  *logger.Logger
	service *service.Service
}

func (s *Server) Open() error {
	s.init(s)
	return nil
}

func (s *Server) Init(init InitializeFunc) error {
	return nil
}

func (s *Server) Service() *service.Service {
	return s.service
}
