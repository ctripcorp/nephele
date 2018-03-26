package log

import (
	"github.com/ctripcorp/nephele/context"
)

type fakeLogger struct{}

func (l *fakeLogger) Printf(ctx *context.Context,
	level string, format string, values ...interface{}) {
}
func (l *fakeLogger) Printw(ctx *context.Context,
	level string, message string, keysAndValues ...interface{}) {
}
