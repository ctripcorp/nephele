package transform

import "github.com/nephele/context"

type Transformer interface {
	Transform(ctx context.Context, blob []byte) ([]byte, error)
}
