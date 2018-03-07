package service

import (
	"github.com/gin-gonic/gin"
	"github.com/nephele/context"
	"github.com/nephele/service/handler"
)

// Service Configuration.
type Config struct {
	BufferSize     int `toml:"buffer-size"`
	MaxConcurrency int `toml:"max-concurrency"`
	RequestTimeout int `toml:"request-timeout"`
}

// Represents http service to handle image request.
type Service struct {
	conf    Config
	image   *ImageService
	router  *gin.Engine
	factory *handler.HandlerFactory
}

// Return service with given context and config.
func New(ctx *context.Context, conf Config) *Service {
	return &Service{
		conf:    conf,
		router:  gin.New(),
		image:   new(ImageService),
		factory: handler.NewFactory(ctx),
	}
}

// Register http GET handler.
func (s *Service) GET(relativePath string, handlers ...handler.HandlerFunc) {
	s.router.GET(relativePath, s.factory.BuildMany(handlers...)...)
}

// Register http POST handler.
func (s *Service) POST(relativePath string, handlers ...handler.HandlerFunc) {
	s.router.POST(relativePath, s.factory.BuildMany(handlers...)...)
}

// Register http DELETE handler.
func (s *Service) DELETE(relativePath string, handlers ...handler.HandlerFunc) {
	s.router.DELETE(relativePath, s.factory.BuildMany(handlers...)...)
}

// Register http PUT handler.
func (s *Service) PUT(relativePath string, handlers ...handler.HandlerFunc) {
	s.router.PUT(relativePath, s.factory.BuildMany(handlers...)...)
}

// Register http OPTIONS handler.
func (s *Service) OPTIONS(relativePath string, handlers ...handler.HandlerFunc) {
	s.router.OPTIONS(relativePath, s.factory.BuildMany(handlers...)...)
}

// Register htt HEAD handler.
func (s *Service) HEAD(relativePath string, handlers ...handler.HandlerFunc) {
	s.router.HEAD(relativePath, s.factory.BuildMany(handlers...)...)
}

// Return image service.
func (s *Service) Image() *ImageService {
	return s.image
}

// run image http service.
func (s *Service) Run() error {
	return nil
}
