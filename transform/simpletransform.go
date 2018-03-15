package transform

import (
	"context"

	"github.com/nephele/process"
)

// SimpleTransform represents how to transform image with given commands
type SimpleTransform struct {
	Processes []process.Process
}

//Transform original image blob to expected blob.
func (s *SimpleTransform) Transform(ctx context.Context, blob []byte) ([]byte, error) {
	return nil, nil
}
