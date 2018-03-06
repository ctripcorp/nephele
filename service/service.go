package service

import (
	"github.com/gin-gonic/gin"
	"github.com/nephele/context"
	"github.com/nephele/service/handler"
)

type Config struct {
	BufferSize     int `toml:"buffer-size"`
	MaxConcurrency int `toml:"max-concurrency"`
	RequestTimeout int `toml:"request-timeout"`
}

type Service struct {
	conf    Config
	image   *ImageService
	router  *gin.Engine
	factory *handler.Factory
}

func New(ctx *context.Context, conf Config) *Service {
	return &Service{
		conf:    conf,
		router:  gin.New(),
		image:   new(ImageService),
		factory: handler.NewFactory(ctx),
	}
}

func (s *Service) GET(relativePath string, handlers ...handler.Func) {
	s.router.GET(relativePath, s.factory.BuildMany(handlers...)...)
}

func (s *Service) POST(relativePath string, handlers ...handler.Func) {
	s.router.POST(relativePath, s.factory.BuildMany(handlers...)...)
}

func (s *Service) DELETE(relativePath string, handlers ...handler.Func) {
	s.router.DELETE(relativePath, s.factory.BuildMany(handlers...)...)
}

func (s *Service) PUT(relativePath string, handlers ...handler.Func) {
	s.router.PUT(relativePath, s.factory.BuildMany(handlers...)...)
}

func (s *Service) OPTIONS(relativePath string, handlers ...handler.Func) {
	s.router.OPTIONS(relativePath, s.factory.BuildMany(handlers...)...)
}

func (s *Service) HEAD(relativePath string, handlers ...handler.Func) {
	s.router.HEAD(relativePath, s.factory.BuildMany(handlers...)...)
}

func (s *Service) Image() *ImageService {
	return s.image
}

func (s *Service) Run() error {
	return nil
}
