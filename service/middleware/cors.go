package middleware

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

// See "github.com/gin-contrib/cors".
type CORSConfig struct {
	MaxAge            int      `toml:"max-age"`
	RegistOrder       int      `toml:"order"`
	OptionPassthrough bool     `toml:"option-pass-through"`
	AllowAllOrigins   bool     `toml:"allow-all-origins"`
	AllowOrigins      []string `toml:"allow-origins"`
	AllowMethods      []string `toml:"allow-methods"`
	AllowHeaders      []string `toml:"allow-headers"`
	ExposeHeaders     []string `toml:"expose-headers"`
}

func (conf *CORSConfig) Order() int {
	return conf.RegistOrder
}

// Return cors handler.
func (conf *CORSConfig) Handler() gin.HandlerFunc {
	var middleware gin.HandlerFunc

	if conf.AllowAllOrigins {
		middleware = cors.AllowAll()
	} else {
		middleware = cors.New(cors.Options{
			MaxAge:             conf.MaxAge,
			AllowedOrigins:     conf.AllowOrigins,
			AllowedMethods:     conf.AllowMethods,
			AllowedHeaders:     conf.AllowHeaders,
			ExposedHeaders:     conf.ExposeHeaders,
			OptionsPassthrough: conf.OptionPassthrough,
		})
	}

	return middleware
}
