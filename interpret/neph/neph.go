package neph

import (
	"github.com/ctripcorp/nephele/interpret"
	"github.com/gin-gonic/gin"
)

func init() {
	interpret.Register(inter)
}

func inter(c *gin.Context) (string, string, error) {
	return c.Param("key"), c.Query("x-nephele-process"), nil
}
