package interpret

import (
	"github.com/gin-gonic/gin"
)

var Config map[string]string

var registeredInterpreter func(*gin.Context) (string, string, error)

func Init() {}

func Register(interpreter func(*gin.Context) (string, string, error)) {
	registeredInterpreter = interpreter
}

func Do(c *gin.Context) (string, string, error) {
	return registeredInterpreter(c)
}
