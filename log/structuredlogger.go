package log

import (
	"github.com/ctripcorp/nephele/context"
)

type StructuredLogger interface {
	Debugw(ctx context.Context, message string, keysAndValues ...interface{})
	Infow(ctx context.Context, message string, keysAndValues ...interface{})
	Warnw(ctx context.Context, message string, keysAndValues ...interface{})
	Errorw(ctx context.Context, message string, keysAndValues ...interface{})
	Fatalw(ctx context.Context, message string, keysAndValues ...interface{})
}
