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
	internal        WriteSyncer
	internalCreater func() (WriteSyncer, error)
	level           string
}

func createBasicOutput(creater func() (WriteSyncer, error), level string) (bo *basicOutput, err error) {
	var internal WriteSyncer
	internal, err = creater()
	if err != nil {
		return nil, err
	}
	return &basicOutput{
		internal,
		creater,
		level,
	}, nil
}

func (bo *basicOutput) Reset() (err error) {
	if bo.internalCreater == nil {
		return nil
	}
	bo.internal, err = bo.internalCreater()
	return err
}

func (bo *basicOutput) Write(p []byte, level string) (n int, err error) {
	if levelInt(level) <= levelInt(bo.level) {
		var internal WriteSyncer
		internal = bo.internal
		n, err = internal.Write(p)
		if err != nil {
			internal = bo.internal
			n, err = internal.Write(p)
		}
		if levelInt(level) <= levelInt("error") {
			defer internal.Sync()
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
