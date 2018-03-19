package transform

import "github.com/ctripcorp/nephele/context"

// Transformer represents how to transform image with given commands
type Transformer interface {
	//Transform original image blob to expected blob.
	Transform(ctx *context.Context, blob []byte) ([]byte, error)
}
