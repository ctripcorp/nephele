package neph

import "github.com/ctripcorp/nephele/context"

// represents v1 encoder
type Encoder struct {
	ctx *context.Context
}

func (e *Encoder) Encode(seed string) string {
	return ""
}
