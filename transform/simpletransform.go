package transform

import "context"

// SimpleTransform represents how to transform image with given commands
type SimpleTransform struct {
}

//Transform original image blob to expected blob.
func (s *SimpleTransform) Transform(ctx context.Context, blob []byte) ([]byte, error) {

}
