package index

import (
	"github.com/ctripcorp/nephele/image"
)

// Quicklly fetch image from index. index is built from image request url.
type Index interface {
	// Find original image for further handling.
	FindOriginalImage() (*image.Image, error)
}
