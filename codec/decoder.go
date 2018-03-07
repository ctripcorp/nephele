package codec

import (
	"github.com/nephele/index"
	"github.com/nephele/transform"
)

// Decoder represnets how to build commands and index from request image url.
type Decoder interface {
	// Decode from request image url.
	// generally there will be multi versions for image file name encoding.
	// and also we will have different image handle commands.
	Decode(uri string) error

	// Create indexer from request image url.
	CreateIndex() index.Index

	// Create transformer from request image url.
	CreateTransformer() transform.Transformer
}
