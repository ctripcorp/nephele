package handler

import (
	_ "github.com/nephele/codec"
	"github.com/nephele/context"
)

type GetImageHandler struct {
}

func (h GetImageHandler) Handler() HandlerFunc {
	return func(ctx *context.Context) {
		//todo:
		/*
			decoder := codec.GetDecoder(ctx)
			decoder.Decode(ctx.HTTP().Request.RequestURI)
			image, _ := decoder.CreateIndex().FindOriginalImage()
			if err := image.Use(decoder.CreateTransformer()).Transform(ctx); err != nil {

			}
		*/
	}
}
