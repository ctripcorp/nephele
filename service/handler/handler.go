package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nephele/context"
)

//Image handle func
type Func func(*context.Context)

//Create or build image handler to match gin web framework.
type Factory struct {
	ctx *context.Context
}

//Return handler factory.
func NewFactory(ctx *context.Context) *Factory {
	return &Factory{ctx: ctx}
}

//Return gin http handler with image handler.
func (f *Factory) Build(handlerFunc Func) gin.HandlerFunc {
	return func(httpCtx *gin.Context) {
		handlerFunc(f.ctx.New(httpCtx))
		httpCtx.Next()
	}
}

// Return multi gin http handlers.
func (f *Factory) BuildMany(handlers ...Func) []gin.HandlerFunc {
	g := make([]gin.HandlerFunc, len(handlers))
	for i, h := range handlers {
		g[i] = f.Build(h)
	}
	return g
}

// Create image get handler.
func (f *Factory) CreateGetImageHandler() gin.HandlerFunc {
	h := GetImageHandler{}
	return f.Build(h.Handler())
}

// Create image upload handler.
func (f *Factory) CreateUploadImageHandler() gin.HandlerFunc {
	h := UploadImageHandler{}
	return f.Build(h.Handler())
}

// Create image delete handler.
func (f *Factory) CreateDeleteImageHandler() gin.HandlerFunc {
	h := DeleteImageHandler{}
	return f.Build(h.Handler())
}
