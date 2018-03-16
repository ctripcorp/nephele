package transform

import (
	"github.com/nephele/context"
	"github.com/nephele/process"
)

// Transformer represents how to transform image with given commands
type Transformer interface {
	//Transform original image blob to expected blob.
	Transform(ctx *context.Context, blob []byte) ([]byte, error)
}
