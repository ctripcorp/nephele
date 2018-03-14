package service

import (
	"github.com/gin-gonic/gin"
	"github.com/nephele/context"
	"github.com/nephele/service/handler"
	"runtime"
)

// Service Configuration.
type Config struct {
	BufferSize     int    `toml:"buffer-size"`
	MaxConcurrency int    `toml:"max-concurrency"`
	RequestTimeout int    `toml:"request-timeout"`
	Address        string `toml:"address"`
}

// Represents http service to handle image request.
type Service struct {
	conf    Config
	image   *ImageService
	router  *gin.Engine
	factory *handler.HandlerFactory
}

// Returns service with given context and config.
func New(ctx *context.Context, conf Config) *Service {
	s := &Service{
		conf:    conf,
		router:  gin.New(),
		factory: handler.NewFactory(ctx),
	}
	s.image = &ImageService{internal: s}
	return s
}

// Return an instance of Config with reasonable defaults.
func NewConfig() (Config, error) {
	return Config{
		Address:        ":80",
		BufferSize:     200,
		RequestTimeout: 3000,
		MaxConcurrency: runtime.NumCPU(),
	}, nil
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

// Register http HEAD handler.
func (s *Service) HEAD(relativePath string, handlers ...handler.HandlerFunc) {
	s.router.HEAD(relativePath, s.factory.BuildMany(handlers...)...)
}

// Return image service.
func (s *Service) Image() *ImageService {
	return s.image
}

// Open image http service.
func (s *Service) Open() error {
	s.image.registerAll()
	return s.router.Run(s.conf.Address)
}
