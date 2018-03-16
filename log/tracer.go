package log

import (
	"github.com/ctripcorp/nephele/context"
)

type Tracer interface {
	TraceBegin(ctx context.Context, keysAndValues ...interface{})
	TraceEnd(ctx context.Context, state interface{}, message ...string)
	TraceEndRoot(ctx context.Context, state interface{}, message ...string)
}
