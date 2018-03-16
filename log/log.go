package log

import (
	"fmt"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/util"
	"path/filepath"
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
	var err error
	var homeDir string

	if homeDir, err = util.HomeDir(); err != nil {
		return nil, err
	}
	println(homeDir)

	return &diskConfig{
		path: filepath.Join(homeDir, "log/today.log"),
	}, nil
}

func Debugf(ctx context.Context, format string, values ...interface{}) {
	instance.Printf(ctx, "debug", format, values...)
}

func Infof(ctx context.Context, format string, values ...interface{}) {
	instance.Printf(ctx, "info", format, values...)
}

func Warnf(ctx context.Context, format string, values ...interface{}) {
	instance.Printf(ctx, "warn", format, values...)
}

func Errorf(ctx context.Context, format string, values ...interface{}) {
	instance.Printf(ctx, "error", format, values...)
}

func Fatalf(ctx context.Context, format string, values ...interface{}) {
	instance.Printf(ctx, "fatal", format, values...)
}

func Debugw(ctx context.Context, message string, keysAndValues ...interface{}) {
	instance.Printw(ctx, "debug", message, keysAndValues...)
}

func Infow(ctx context.Context, message string, keysAndValues ...interface{}) {
	instance.Printw(ctx, "info", message, keysAndValues...)
}

func Warnw(ctx context.Context, message string, keysAndValues ...interface{}) {
	instance.Printw(ctx, "warn", message, keysAndValues...)
}

func Errorw(ctx context.Context, message string, keysAndValues ...interface{}) {
	instance.Printw(ctx, "error", message, keysAndValues...)
}

func Fatalw(ctx context.Context, message string, keysAndValues ...interface{}) {
	instance.Printw(ctx, "fatal", message, keysAndValues...)
}

func TraceBegin(ctx context.Context, message string, keysAndValues ...interface{}) {
	instance.Printw(ctx, "trace/begin", message, keysAndValues...)
}

func TraceEnd(ctx context.Context, state interface{}) {
	instance.Printw(ctx, "trace/end", fmt.Sprintf("%v", state))
}

func TraceEndRoot(ctx context.Context, state interface{}) {
	instance.Printw(ctx, "trace/endroot", fmt.Sprintf("%v", state))
}
