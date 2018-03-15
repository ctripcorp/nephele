package v1

import (
	"github.com/nephele/context"
	"github.com/nephele/index"
	"github.com/nephele/transform"
)

// represents v1 decoder
type Decoder struct {
	ctx *context.Context
}

func (e *Decoder) Decode(uri string) error {
	return nil
}

func (e *Decoder) CreateIndex() index.Index {
	return nil
}

func (e *Decoder) CreateTransformer() transform.Transformer {
	return nil
}
