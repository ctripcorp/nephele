package middleware

import (
	"github.com/gin-gonic/gin"
	"testing"
)

type StructTestConfig struct {
	MiddlewareConfig
	Reader *ReaderConfig
}

type PointerTestConfig struct {
	*MiddlewareConfig
	Reader *ReaderConfig
}

type ReaderConfig struct {
	RegisterOrder int
	Path          string
}

func (r *ReaderConfig) Order() int {
	return r.RegisterOrder
}

func (r *ReaderConfig) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func TestDefaultBuild(t *testing.T) {
	conf := &MiddlewareConfig{
		CORS: &CORSConfig{
			RegisterOrder:   1,
			AllowAllOrigins: true,
		},
	}

	if handlers := Build(conf); len(handlers) != 2 {
		t.Errorf("expect handlers :2, but got %d", len(handlers))
	}

	conf.CORS = nil
	if handlers := Build(conf); len(handlers) != 1 {
		t.Errorf("expect handlers :1, but got %d", len(handlers))
	}
}

func TestStructBuild(t *testing.T) {
	conf := &StructTestConfig{
		MiddlewareConfig: MiddlewareConfig{
			CORS: &CORSConfig{
				AllowAllOrigins: true,
			},
		},
		Reader: &ReaderConfig{
			RegisterOrder: 2,
			Path:          "/home",
		},
	}
	if handlers := Build(conf); len(handlers) != 3 {
		t.Errorf("expect handlers :3, but got %d", len(handlers))
	}
	conf.Reader = nil
	conf.MiddlewareConfig.CORS = nil
	if handlers := Build(conf); len(handlers) != 1 {
		t.Errorf("expect handlers :1, but got %d", len(handlers))
	}
}

func TestPointerBuild(t *testing.T) {
	conf := &PointerTestConfig{
		MiddlewareConfig: &MiddlewareConfig{
			CORS: &CORSConfig{
				RegisterOrder:   1,
				AllowAllOrigins: true,
			},
		},
		Reader: &ReaderConfig{
			RegisterOrder: 2,
			Path:          "/home",
		},
	}
	if handlers := Build(conf); len(handlers) != 3 {
		t.Errorf("expect handlers :3, but got %d", len(handlers))
	}
	conf.Reader = nil
	conf.MiddlewareConfig.CORS = nil
	if handlers := Build(conf); len(handlers) != 1 {
		t.Errorf("expect handlers :1, but got %d", len(handlers))
	}
}
