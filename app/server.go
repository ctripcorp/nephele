package app

import (
	"github.com/nephele/service"
)

// Define server initialization function type.
type ServerInitializeFunc func(*Server) error

// Server represents holder for service and
// is the entry for all components to initialize, open or quit.
type Server struct {
	service *service.Service
	init    ServerInitializeFunc
}

// Call to make external initialization.
func (s *Server) Init(init ServerInitializeFunc) {
	s.init = init
}

// Open service and other components.
func (s *Server) Open() <-chan error {
	c := make(chan error)

	if s.init != nil {
		if err := s.init(s); err != nil {
			c <- err
			return c
		}
	}

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
