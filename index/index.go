package index

import "github.com/nephele/image"

// Quicklly fetch image from index. index is built from image request url.
type Index interface {
	// Find image quicklly with given key.
	// usually we can use local disk cache, memory cache or distributed storage system
	// to cache handled image. that's a performance concern. also varnish or squid is a
	// better choice.
	FindBy(key string) (*image.Image, error)

	// Find original image for further handling.
	FindOriginalBy(key string) (*image.Image, error)
}
