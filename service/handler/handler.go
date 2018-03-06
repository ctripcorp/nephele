package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nephele/context"
)

type Func func(*context.Context)

type Factory struct {
	ctx *context.Context
}

func NewFactory(ctx *context.Context) *Factory {
	return &Factory{ctx: ctx}
}

func (f *Factory) Build(handlerFunc Func) gin.HandlerFunc {
	return func(httpCtx *gin.Context) {
		handlerFunc(f.ctx.New(httpCtx))
		httpCtx.Next()
	}
}

func (f *Factory) BuildMany(handlers ...Func) []gin.HandlerFunc {
	g := make([]gin.HandlerFunc, len(handlers))
	for i, h := range handlers {
		g[i] = f.Build(h)
	}
	return g
}

func (f *Factory) CreateGetImageHandler() gin.HandlerFunc {
	h := GetImageHandler{}
	return f.Build(h.Handler())
}

func (f *Factory) CreateUploadImageHandler() gin.HandlerFunc {
	h := UploadImageHandler{}
	return f.Build(h.Handler())
}

func (f *Factory) CreateDeleteImageHandler() gin.HandlerFunc {
	h := DeleteImageHandler{}
	return f.Build(h.Handler())
}
