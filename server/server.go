package server

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/satori/go.uuid"

	"context"
	"time"
)

var Config struct {
	Port     string `toml:"port"`
	Recovery struct{}
	Entrance struct {
		RequestBuffer             int `toml:"request_buffer"`
		RequestTimeoutMillisecond int `toml:"request_timeout_millisecond"`
		MaxConcurrency            int `toml:"max_concurrency"`
	}
	CORS struct {
		AllowAllOrigins   bool     `toml:"allow_all_origins"`
		MaxAge            int      `toml:"max_age"`
		OptionPassThrough bool     `toml:"option_pass_through"`
		AllowedOrigins    []string `toml:"allowed_origins"`
		AllowedMethods    []string `toml:"allowed_methods"`
		AllowedHeaders    []string `toml:"allowed_headers"`
		ExposedHeaders    []string `toml:"exposed_headers"`
	}
}

func Entrance() gin.HandlerFunc {
	var conf = Config.Entrance
	return func(c *gin.Context) {
		var (
			ctx      context.Context
			cancel   context.CancelFunc
			uuidUUID uuid.UUID
			err      error
		)
		uuidUUID, err = uuid.NewV1()
		if err != nil {
			c.String(500, err.Error())
			return
		}
		ctx = context.WithValue(context.Background(), "contextID", uuidUUID.String())
		ctx = context.WithValue(ctx, "request", c.Request)
		ctx, cancel = context.WithTimeout(ctx,
			time.Duration(conf.RequestTimeoutMillisecond)*time.Millisecond)
		c.Set("context", ctx)
		c.Next()
		cancel()
	}
}

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatus(500)
			}
		}()
		c.Next()
	}
}

func CORS() gin.HandlerFunc {
	var conf = Config.CORS

	if conf.AllowAllOrigins {
		return cors.AllowAll()
	}
	return cors.New(cors.Options{
		MaxAge:             conf.MaxAge,
		OptionsPassthrough: conf.OptionPassThrough,
		AllowedOrigins:     conf.AllowedOrigins,
		AllowedMethods:     conf.AllowedMethods,
		AllowedHeaders:     conf.AllowedHeaders,
		ExposedHeaders:     conf.ExposedHeaders,
	})
}

func Run() {
	r := gin.New()

	r.Use(Entrance())
	r.Use(Recovery())
	r.Use(CORS())

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/image/:key", get)

	r.Run(Config.Port)
}
