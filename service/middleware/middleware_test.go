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
	Writer *WriterConfig
}

type ReaderConfig struct {
	RegistOrder int
	Path        string
}

func (r *ReaderConfig) Order() int {
	return r.RegistOrder
}

func (r *ReaderConfig) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		configOrder = append(configOrder, "reader")
	}
}

type WriterConfig struct {
	RegistOrder int
}

func (w *WriterConfig) Order() int {
	return w.RegistOrder
}

func (w *WriterConfig) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		configOrder = append(configOrder, "writer")
	}
}

var configOrder = make([]string, 0)

func TestDefaultBuild(t *testing.T) {
	conf := &MiddlewareConfig{
		CORS: &CORSConfig{
			RegistOrder:     1,
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
			RegistOrder: 2,
			Path:        "/home/",
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
				RegistOrder:     1,
				AllowAllOrigins: true,
			},
		},
		Reader: &ReaderConfig{
			RegistOrder: 2,
			Path:        "/home/",
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

func TestOrderedConfigBuild(t *testing.T) {
	conf := &PointerTestConfig{
		MiddlewareConfig: &MiddlewareConfig{
			CORS: &CORSConfig{
				RegistOrder:     2,
				AllowAllOrigins: true,
			},
		},
		Reader: &ReaderConfig{
			RegistOrder: 4,
			Path:        "/home/",
		},
		Writer: &WriterConfig{
			RegistOrder: 3,
		},
	}

	handlers := Build(conf)

	if len(handlers) != 4 {
		t.Errorf("expect handlers :4, but got %d", len(handlers))
	}

	handlers[2](nil)
	handlers[3](nil)

	if configOrder[0] != "writer" && configOrder[1] != "reader" {
		t.Errorf("expect order:[`writer`, `reader`] but got %v", configOrder)
	}
}
