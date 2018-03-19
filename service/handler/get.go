package handler

import (
	"strings"

	"github.com/ctripcorp/nephele/codec"
	"github.com/ctripcorp/nephele/context"
)

func getImageHandler() HandlerFunc {
	return func(ctx *context.Context) {
		decoder := codec.GetDecoder(ctx)
		decoder.Decode(strings.TrimPrefix(ctx.HTTP().Request.RequestURI, "/image/"))
		image, err := decoder.CreateIndex().FindOriginalImage()
		if err != nil {
			ctx.HTTP().String(400, err.Error())
			return
		}
		if err := image.Use(decoder.CreateTransformer()).Transform(ctx); err != nil {
			ctx.HTTP().String(400, err.Error())
		}
		ctx.HTTP().Writer.Write(image.Blob())
		//ctx.HTTP().String(200, "hello.world")
	}
}
