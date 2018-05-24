package server

import (
	"github.com/gin-gonic/gin"

    "strings"
)

struct {
    Name string
    Option map[string]string
}

func get(c *gin.Context) {
    key := c.Param("key")
    process := c.Query("x-nephele-process")
    uri := strings.TrimPrefix(c.Request.RequestURI, "/image/")
    println(key, process, uri)
}
