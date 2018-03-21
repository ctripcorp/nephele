package output

import (
	"io"
)

type WriteSyncer interface {
	io.Writer
	Sync() error
}

type Output interface {
	Write(p []byte, level string) (n int, err error)
	Sync() error
}

type basicOutput struct {
	internal WriteSyncer
	level    string
}

func (bo *basicOutput) Write(p []byte, level string) (n int, err error) {
	if levelInt(level) <= levelInt(bo.level) {
		if levelInt(level) <= levelInt("error") {
			defer bo.internal.Sync()
		}
		return bo.internal.Write(p)
	}
	return 0, nil
}

func (bo *basicOutput) Sync() error {
	return bo.internal.Sync()
}

func (bo *basicOutput) Level() string {
	return bo.level
}
