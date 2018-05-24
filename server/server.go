package server

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

var Config struct {
	Port     string `toml:"port"`
	Recovery struct{}
	Entrance struct {
		RequestBuffer            int `toml:"request_buffer"`
		RequestTimeoutMilisecond int `toml:"request_timeout_milisecond"`
	}
	Throttle struct {
		MaxConcurrency int `toml:"max_concurrency"`
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

func Recovery() gin.HandlerFunc {
	return func(httpCtx *gin.Context) {
		defer func() {
			if recover() != nil {
				httpCtx.AbortWithStatus(500)
			}
		}()
		httpCtx.Next()
	}
}

func Entrance() gin.HandlerFunc {
	return func(httpCtx *gin.Context) {
		httpCtx.Next()
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

	r.Use(Recovery())
	r.Use(Entrance())
	r.Use(CORS())

    r.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })

    r.GET("/image/:key", get)

	r.Run(Config.Port)
}
