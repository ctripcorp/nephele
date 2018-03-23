package log

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/log/output"
	"github.com/ctripcorp/nephele/util"
)

var instance Logger = &fakeLogger{}

func Init(loggerConfig *LoggerConfig) (err error) {
	instance, err = loggerConfig.Build()
	return
}

func DefaultConfig() (*LoggerConfig, error) {
	var err error
	var hd string
	var d string

	hd, err = util.HomePath()
	if err != nil {
		return nil, err
	}

	d = filepath.Join(hd, "log/")
	err = os.MkdirAll(d, 0777)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &LoggerConfig{
		&output.StdoutConfig{"debug"},
		&output.DumpConfig{"info", d},
	}, nil
}

func Debugf(ctx *context.Context, format string, values ...interface{}) {
	instance.Printf(ctx, "debug", format, values...)
}

func Infof(ctx *context.Context, format string, values ...interface{}) {
	instance.Printf(ctx, "info", format, values...)
}

func Warnf(ctx *context.Context, format string, values ...interface{}) {
	instance.Printf(ctx, "warn", format, values...)
}

func Errorf(ctx *context.Context, format string, values ...interface{}) {
	instance.Printf(ctx, "error", format, values...)
}

func Fatalf(ctx *context.Context, format string, values ...interface{}) {
	instance.Printf(ctx, "fatal", format, values...)
}

func Debugw(ctx *context.Context, message string, keysAndValues ...interface{}) {
	instance.Printw(ctx, "debug", message, keysAndValues...)
}

func Infow(ctx *context.Context, message string, keysAndValues ...interface{}) {
	instance.Printw(ctx, "info", message, keysAndValues...)
}

func Warnw(ctx *context.Context, message string, keysAndValues ...interface{}) {
	instance.Printw(ctx, "warn", message, keysAndValues...)
}

func Errorw(ctx *context.Context, message string, keysAndValues ...interface{}) {
	instance.Printw(ctx, "error", message, keysAndValues...)
}

func Fatalw(ctx *context.Context, message string, keysAndValues ...interface{}) {
	instance.Printw(ctx, "fatal", message, keysAndValues...)
}

func TraceBegin(ctx *context.Context, message string, keysAndValues ...interface{}) {
	instance.Printw(ctx, "trace/begin", message, keysAndValues...)
}

func TraceEnd(ctx *context.Context, state interface{}) {
	instance.Printw(ctx, "trace/end", fmt.Sprintf("%v", state))
}

func TraceEndRoot(ctx *context.Context, state interface{}) {
	instance.Printw(ctx, "trace/endroot", fmt.Sprintf("%v", state))
}
