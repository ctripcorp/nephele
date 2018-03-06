package index

import "github.com/nephele/image"

type Index interface {
	FindBy(key string) (*image.Image, error)
	FindPrototypeBy(key string) (*image.Image, error)
}
