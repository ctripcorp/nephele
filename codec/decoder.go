package codec

import (
	"github.com/nephele/index"
	"github.com/nephele/transform"
)

type Decoder interface {
	Decode(uri string) error
	CreateIndex() index.Index
	CreateTransformer() transform.Transformer
}
