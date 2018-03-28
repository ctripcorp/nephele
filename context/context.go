package context

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"time"
)

const GlobalName = "@@GlobalContextName"

// Context walks through all the functions.
type Context struct {
	id       string
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
	uuid, _ := uuid.NewV1()
	return &Context{
		id:       uuid.String(),
		env:      ctx.env,
		timeout:  ctx.timeout,
		http:     httpCtx,
		internal: internal,
		cancel:   cancel,
	}
}

// Return current env
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

// Context has been canceled for timed out or deadline exceeded
func (ctx *Context) Canceled() bool {
	return ctx.internal.Err() != nil
}

// Return context deadline.
func (ctx *Context) Deadline() (time.Time, bool) {
	return ctx.internal.Deadline()
}

// Wait for cancelation
func (ctx *Context) Done() <-chan struct{} {
	return ctx.internal.Done()
}

// Return context id
func (ctx *Context) ID() string {
	return ctx.id
}
