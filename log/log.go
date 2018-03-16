package log

import (
	"github.com/ctripcorp/nephele/context"
)

type Config interface {
	BuildLogger() (Logger, error)
}

var instance Logger

func Init(conf Config) (err error) {
	instance, err = conf.BuildLogger()
	return
}

func DefaultConfig() (Config, error) {
	return nil, nil
}

func Debugf(ctx context.Context, format string, values ...interface{}) {
	instance.Debugf(ctx, format, values...)
}

func Infof(ctx context.Context, format string, values ...interface{}) {
	instance.Infof(ctx, format, values...)
}

func Warnf(ctx context.Context, format string, values ...interface{}) {
	instance.Warnf(ctx, format, values...)
}

func Errorf(ctx context.Context, format string, values ...interface{}) {
	instance.Errorf(ctx, format, values...)
}

func Fatalf(ctx context.Context, format string, values ...interface{}) {
	instance.Fatalf(ctx, format, values...)
}

func Debugw(ctx context.Context, message string, keysAndValues ...interface{}) {
	instance.Debugw(ctx, message, keysAndValues...)
}

func Infow(ctx context.Context, message string, keysAndValues ...interface{}) {
	instance.Infow(ctx, message, keysAndValues...)
}

func Warnw(ctx context.Context, message string, keysAndValues ...interface{}) {
	instance.Warnw(ctx, message, keysAndValues...)
}

func Errorw(ctx context.Context, message string, keysAndValues ...interface{}) {
	instance.Errorw(ctx, message, keysAndValues...)
}

func Fatalw(ctx context.Context, message string, keysAndValues ...interface{}) {
	instance.Fatalw(ctx, message, keysAndValues...)
}

func TraceBegin(ctx context.Context, keysAndValues ...interface{}) {
	instance.TraceBegin(ctx, keysAndValues...)
}

func TraceEnd(ctx context.Context, state interface{}, message ...string) {
	instance.TraceEnd(ctx, state, message...)
}

func TraceEndRoot(ctx context.Context, state interface{}, message ...string) {
	instance.TraceEndRoot(ctx, state, message...)
}
