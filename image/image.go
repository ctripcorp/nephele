package image

import (
	"github.com/nephele/context"
	"github.com/nephele/transform"
)

type Image struct {
	blob        []byte
	meta        *Meta
	transformer transform.Transformer
}

func New(blob []byte) *Image {
	return nil
}

func (img *Image) Meta() *Meta {
	return img.meta
}

func (img *Image) Blob() []byte {
	return img.blob
}

func (img *Image) Use(transformer transform.Transformer) *Image {
	img.transformer = transformer
	return img
}

func (img *Image) Transform(ctx context.Context) error {
	var err error
	var blob []byte
	if blob, err = img.transformer.Transform(ctx, img.blob); err == nil {
		img.blob = blob
	}
	return err
}
