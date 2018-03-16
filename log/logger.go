package log

import (
	"github.com/ctripcorp/nephele/context"
)

type Logger interface {
	Printf(ctx context.Context, level string, format string, values ...interface{})
	Printw(ctx context.Context, level string, message string, keysAndValues ...interface{})
}
