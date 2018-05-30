package server

import (
	"github.com/gin-gonic/gin"

	"github.com/ctripcorp/nephele/interpret"
	"github.com/ctripcorp/nephele/process"
	"github.com/ctripcorp/nephele/storage"

	"context"
)

func get(c *gin.Context) {

	var (
		ctx = c.MustGet("context").(context.Context)
	)

	key, proc, err := interpret.Do(c)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	commands, err := process.Parse(ctx, proc)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	blob, _, err := storage.Download(ctx, key)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	blob, err = process.Do(ctx, blob, commands)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.Writer.Write(blob)
}
