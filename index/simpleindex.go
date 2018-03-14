package index

import "image"

type SimpleIndex struct {
}

// FindBy Find image quicklly with given key.
// usually we can use local disk cache, memory cache or distributed storage system
// to cache handled image. that's a performance concern. also varnish or squid is a
// better choice.
func (s *SimpleIndex) FindBy(key string) (*image.Image, error) {
	return nil, nil
}

// FindOriginalBy Find original image for further handling.
func (s *SimpleIndex) FindOriginalBy(key string) (*image.Image, error) {
	return nil, nil
}
