package log

import (
	"bytes"
	"fmt"
	"sync"
	"time"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/log/output"
)

type Logger interface {
	Printf(ctx *context.Context, level string, format string, values ...interface{})
	Printw(ctx *context.Context, level string, message string, keysAndValues ...interface{})
}

type logger struct {
	outputs    []output.Output
	bufferPool *sync.Pool
}

func (l *logger) Printf(ctx *context.Context, level string, format string, values ...interface{}) {
	var rb *bytes.Buffer = l.bufferPool.Get().(*bytes.Buffer)
	fmt.Fprintf(rb, "[%s] %s [%s]", level, ctx.ID(), time.Now().Format(time.RFC3339))
	fmt.Fprintf(rb, "\t\""+format+"\"\n", values...)

	for _, o := range l.outputs {
		o.Write(rb.Bytes(), level)
	}
}

func (l *logger) Printw(ctx *context.Context, level string, message string, keysAndValues ...interface{}) {
	var rb *bytes.Buffer = l.bufferPool.Get().(*bytes.Buffer)
	fmt.Fprintf(rb, "[%s] %s [%s]", level, ctx.ID(), time.Now().Format(time.RFC3339))
	for i := 0; i < len(keysAndValues)/2; i++ {
		fmt.Fprintf(rb, "\t\"%v\"", keysAndValues[i*2])
		fmt.Fprintf(rb, "\t%v", keysAndValues[i*2+1])
	}
	fmt.Fprintf(rb, "\t\"%s\"\n", message)

	for _, o := range l.outputs {
		o.Write(rb.Bytes(), level)
	}
}

type LoggerConfig struct {
	Stdout *output.StdoutConfig
	Dump   *output.DumpConfig
}

func (lc *LoggerConfig) Build() (Logger, error) {
	var stdout, dump output.Output
	var err error
	stdout, err = lc.Stdout.Build()
	dump, err = lc.Dump.Build()
	return &logger{
		[]output.Output{
			stdout,
			dump,
		},
		&sync.Pool{
			New: func() interface{} { return new(bytes.Buffer) },
		},
	}, err
}
