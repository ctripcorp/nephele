package context

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

// Context walks through all the functions.
type Context struct {
	env      string
	timeout  time.Duration
	internal context.Context
	cancel   context.CancelFunc
	http     *gin.Context
}

// Return root context and share its variables.
func New(env string, timeout time.Duration) *Context {
	return &Context{
		env:     env,
		timeout: timeout * time.Millisecond,
	}
}

// Return sub context with *gin.Context.
func (ctx *Context) New(httpCtx *gin.Context) *Context {
	internal, cancel := context.WithTimeout(context.Background(), ctx.timeout)
	return &Context{
		env:      ctx.env,
		timeout:  ctx.timeout,
		http:     httpCtx,
		internal: internal,
		cancel:   cancel,
	}
}

// Return current environment
func (ctx *Context) Env() string {
	return ctx.env
}

// Return *gin.Context.
func (ctx *Context) HTTP() *gin.Context {
	return ctx.http
}

// Cancel image handling.
func (ctx *Context) Cancel() {
	ctx.cancel()
}

// Return context deadline.
func (ctx *Context) Deadline() (time.Time, bool) {
	return ctx.internal.Deadline()
}
