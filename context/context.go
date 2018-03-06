package context

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

type Context struct {
	timeout  time.Duration
	internal context.Context
	cancel   context.CancelFunc
	http     *gin.Context
}

func New(timeout time.Duration) *Context {
	return &Context{
		timeout: timeout * time.Millisecond,
	}
}

func (ctx *Context) New(httpCtx *gin.Context) *Context {
	internal, cancel := context.WithTimeout(context.Background(), ctx.timeout)
	return &Context{
		timeout:  ctx.timeout,
		http:     httpCtx,
		internal: internal,
		cancel:   cancel,
	}
}

func (ctx *Context) HTTP() *gin.Context {
	return ctx.http
}

func (ctx *Context) Cancel() {
	ctx.cancel()
}

func (ctx *Context) Deadline() (time.Time, bool) {
	return ctx.internal.Deadline()
}
