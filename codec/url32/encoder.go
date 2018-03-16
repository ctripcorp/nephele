package v1

import "github.com/nephele/context"

// represents v1 encoder
type Encoder struct {
	ctx *context.Context
}

func (e *Encoder) Encode(seed string) string {
	return ""
}
