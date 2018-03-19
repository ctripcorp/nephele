package app

import (
	"github.com/nephele/codec"
	"github.com/nephele/log"
	"github.com/nephele/service"
	"github.com/nephele/store"
)

// Define server initialization function type.
type ServerInitializeFunc func(*Server) error

// Server represents holder for service and
// is the entry for all components to initialize, open or quit.
type Server struct {
	conf     Config
	service  *service.Service
	initFunc ServerInitializeFunc
}

// Call to make external initialization.
func (s *Server) Init(init ServerInitializeFunc) {
	s.initFunc = init
}

// Open image service
func (s *Server) Open() <-chan error {
	c := make(chan error)
	go func() {
		c <- s.service.Open()
	}()

	return c
}

// Quit service and other components gracefully.
func (s *Server) Quit() error {
	return nil
}

// Return service to register router or modify image handler path.
func (s *Server) Service() *service.Service {
	return s.service
}

// Initialize components and call external init func.
func (s *Server) init() error {
	var err error

	// init storage.
	if err = store.Init(s.conf.Store()); err != nil {
		return err
	}

	// init codec to encode or decode image request URL.
	if err = codec.Init(s.conf.Codec()); err != nil {
		return err
	}

	// init logger
	if err = log.Init(s.conf.Logger()); err != nil {
		return err
	}

	if s.initFunc != nil {
		err = s.initFunc(s)
	}

	return err
}
