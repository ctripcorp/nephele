package image

import (
	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/transform"
)

type Type string

const (
	GIF  Type = "gif"
	PNG  Type = "png"
	JPEG Type = "jpeg"
	WEBP Type = "webp"
)

type Image struct {
	blob        []byte
	meta        *Meta
	transformer transform.Transformer
}

// Return image with body filled.
func New(blob []byte) *Image {
	return &Image{blob: blob}
}

// Return image meta.
func (img *Image) Meta() *Meta {
	return img.meta
}

// Return image blob.
func (img *Image) Blob() []byte {
	return img.blob
}

// Use transformer to transform image.
func (img *Image) Use(transformer transform.Transformer) *Image {
	img.transformer = transformer
	return img
}

// Transform image with given context.
func (img *Image) Transform(ctx *context.Context) error {
	var err error
	var blob []byte
	if blob, err = img.transformer.Transform(ctx, img.blob); err == nil {
		img.blob = blob
	}
	return err
}
