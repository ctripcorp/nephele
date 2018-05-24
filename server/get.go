package server

import (
	"github.com/gin-gonic/gin"
    "github.com/satori/go.uuid"

	"github.com/ctripcorp/nephele/interpret"
	"github.com/ctripcorp/nephele/process"
	storage "github.com/ctripcorp/nephele/storage/neph"

    "context"
)

func get(c *gin.Context) {
	key, proc, err := interpret.Do(c)
	if err != nil {
		c.String(400, err.Error())
		return
	}

    uuid, err := uuid.NewV1()
    if err != nil {
		c.String(500, err.Error())
		return
    }
    ctx := context.WithValue(c, "contextID", uuid)

    commands, err := process.Parse(ctx, proc)
    if err != nil {
		c.String(400, err.Error())
		return
    }

	//blob, rid, err := storage.File(key).Bytes()
	blob, _, err := storage.File(key).Bytes()
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
