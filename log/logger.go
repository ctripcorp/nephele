package log

import (
	"bytes"
	"fmt"
    "io"
	"os"
	"time"
    "path/filepath"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/log/output"
)

type Logger interface {
	Printf(ctx *context.Context, level string, format string, values ...interface{})
	Printw(ctx *context.Context, level string, message string, keysAndValues ...interface{})
}

type logger struct {
    output Output
}

func (l *logger) Printf(ctx *context.Context, level string, format string, values ...interface{}) {
    var rb *bytes.Buffer = new(bytes.Buffer)
    fmt.Fprintf(rb, "[%s] %s [%s]", level, ctx.ID(), time.Now().Format(time.RFC3339))
    fmt.Fprintf(l.output, "\t\"+format+\"\n", values...)
}

func (l *logger) Printw(ctx *context.Context, level string, message string, keysAndValues ...interface{}) {
    var rb *bytes.Buffer = new(bytes.Buffer)
    fmt.Fprintf(rb, "[%s] %s [%s]", level, ctx.ID(), time.Now().Format(time.RFC3339))
	for i := 0; i < len(keysAndValues)/2; i++ {
		fmt.Fprintf(rb, "\t\"%v\"", keysAndValues[i*2])
		fmt.Fprintf(rb, "\t%v", keysAndValues[i*2+1])
	}
    fmt.Fprintf(l.output, "\t\"%s\"\n", message)
}

type LoggerConfig struct {
    output []output.OutputConfig
}

func (dc *LoggerConfig) BuildLogger() (Logger, error) {
    f, err := os.OpenFile(filepath.Join(dc.Path, "now.log"), 
            os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
    l := &dump{
        f,
    } 
	return l, nil
}
