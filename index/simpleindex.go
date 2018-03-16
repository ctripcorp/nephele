package index

import (
	"image"

	"github.com/nephele/store"
)

//SimpleIndex simple index
type SimpleIndex struct {
	Path string
}

// FindOriginalImage Find original image for further handling.
func (s *SimpleIndex) FindOriginalImage(key string) (*image.Image, error) {
	blob, err := store.Storage().Read(ctx, s.Path)
	if err != nil {
		return nil, err
	}
	return image.New(blob), nil
}
