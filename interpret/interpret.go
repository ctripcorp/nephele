package interpret

import (
	"github.com/gin-gonic/gin"
)

var Config map[string]string

var interpreter func(*gin.Context) (string, string, error)

func Register(inter func(*gin.Context) (string, string, error)) {
	interpreter = inter
}

func Do(c *gin.Context) (string, string, error) {
	return interpreter(c)
}
