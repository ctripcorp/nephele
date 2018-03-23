package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// See "github.com/gin-contrib/cors".
type CORSConfig struct {
	MaxAge          int      `toml:"max-age"`
	RegisterOrder   int      `toml:"order"`
	AllowAllOrigins bool     `toml:"allow-all-origins"`
	AllowOrigins    []string `toml:"allow-origins"`
	AllowMethods    []string `toml:"allow-methods"`
	AllowHeaders    []string `toml:"allow-headers"`
	ExposeHeaders   []string `toml:"expose-headers"`
}

func (conf *CORSConfig) Order() int {
	return conf.RegisterOrder
}

// Return cors handler.
func (conf *CORSConfig) Handler() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins: conf.AllowAllOrigins,
		AllowOrigins:    conf.AllowOrigins,
		AllowMethods:    conf.AllowMethods,
		AllowHeaders:    conf.AllowHeaders,
		ExposeHeaders:   conf.ExposeHeaders,
		MaxAge:          time.Duration(conf.MaxAge) * time.Second,
	})
}
