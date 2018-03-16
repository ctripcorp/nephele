package codec

import (
	"github.com/nephele/index"
	"github.com/nephele/process"
	"github.com/nephele/transform"
)

// SimpleDecoder represnets how to build commands and index from request image url.
type SimpleDecoder struct {
	
}

// Decode from request image url.
// generally there will be multi versions for image file name encoding.
// and also we will have different image handle commands.
func (s *SimpleDecoder) Decode(uri string) error {
	
}

// CreateIndex  from request image url.
func (s *SimpleDecoder) CreateIndex() index.Index {
	return index.SimpleIndex{}
}

// CreateTransformer  from request image url.
func (s *SimpleDecoder) CreateTransformer() transform.Transformer {
	return SimpleTransform{}
}
