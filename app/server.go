package app

import (
	"github.com/ctripcorp/nephele/codec"
	"github.com/ctripcorp/nephele/log"
	"github.com/ctripcorp/nephele/service"
	"github.com/ctripcorp/nephele/store"
)

// Define server initialization function type.
// It's called before server opened.
type ServerInitializeFunc func(*Server) error

// Define server quit function type.
// It's called after service no longer serving any image request.
type ServerQuitFunc func(*Server) error

// Server represents holder for service and
// is the entry for all components to initialize, open or quit.
type Server struct {
	conf     Config
	service  *service.Service
	quitFunc ServerQuitFunc
	initFunc ServerInitializeFunc
}

// Call to make external components initialization.
func (s *Server) OnInit(init ServerInitializeFunc) {
	s.initFunc = init
}

// call to quit external components
func (s *Server) OnQuit(quit ServerQuitFunc) {
	s.quitFunc = quit
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

// Open image service
func (s *Server) open() <-chan error {
	c := make(chan error)
	go func() {
		c <- s.service.Open()
	}()

	return c
}

// Quit service and other components gracefully.
func (s *Server) quit() error {
	var err error

	if err = s.service.Quit(); err != nil {
		return err
	}

	if s.quitFunc != nil {
		err = s.quitFunc(s)
	}

	return err
}
