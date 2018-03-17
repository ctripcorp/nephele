package index

import (
	"github.com/nephele/context"
	"github.com/nephele/image"
	"github.com/nephele/store"
)

//Index  index
type Index struct {
	Ctx  *context.Context
	Path string
}

// FindOriginalImage Quicklly fetch image from index. index is built from image request url.
func (i *Index) FindOriginalImage() (*image.Image, error) {
	blob, err := store.Storage().Read(*i.Ctx, i.Path)
	if err != nil {
		return nil, err
	}
	return image.New(blob), nil
}
