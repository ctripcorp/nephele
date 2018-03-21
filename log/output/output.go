package output

import (
    "io"
)

type Output interface {
    Write(p []byte, level string) (n int, err error)
}

type WriteSyncer interface {
    io.Writer
    Sync() error
}

type basicOutput struct {
    internal WriteSyncer
    level string
}

func (bo *basicOutput) Write(p []byte, level string) (n int, err error) {
    if level == bo.level {
        if level == "error" {
            defer bo.internal.Sync()
        }
        return bo.internal.Write(p)
    }
    return 0, nil
}

func (bo *basicOutput) Sync() error {
    return bo.internal.Sync()
}
