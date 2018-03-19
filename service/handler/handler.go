package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nephele/context"
)

//Image handle func
type HandlerFunc func(*context.Context)

//Create or build image handler to match gin web framework.
type HandlerFactory struct {
	ctx *context.Context
}

//Return handler factory.
func NewFactory(ctx *context.Context) *HandlerFactory {
	return &HandlerFactory{ctx: ctx}
}

//Return gin http handler with image handler.
func (f *HandlerFactory) Build(handlerFunc HandlerFunc) gin.HandlerFunc {
	return func(httpCtx *gin.Context) {
		handlerFunc(f.ctx.New(httpCtx))
		httpCtx.Next()
	}
}

// Return multi gin http handlers.
func (f *HandlerFactory) BuildMany(handlers ...HandlerFunc) []gin.HandlerFunc {
	g := make([]gin.HandlerFunc, len(handlers))
	for i, h := range handlers {
		g[i] = f.Build(h)
	}
	return g
}

// Create image get handler.
func (f *HandlerFactory) CreateGetImageHandler() HandlerFunc {
	return getImageHandler()
}

// Create image upload handler.
func (f *HandlerFactory) CreateUploadImageHandler() HandlerFunc {
	return uploadImageHandler()
}

// Create image delete handler.
func (f *HandlerFactory) CreateDeleteImageHandler() HandlerFunc {
	return deleteImageHandler()
}

func (f *HandlerFactory) CreateHealthcheckHandler() HandlerFunc {
	return healthcheckHandler()
}
