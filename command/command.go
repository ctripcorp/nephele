package command

import (
	"context"
	"github.com/ctrip-nephele/gmagick"
)

var list map[string]func() Command

type Command interface {
	Support() string
	Verify(context.Context, map[string]string) error
	ExecuteOnBlob(context.Context, []byte) ([]byte, error)
	ExecuteOnWand(context.Context, *gmagick.MagickWand) error
}

func List() map[string]func() Command {
	return list
}

func Register(name string, command func() Command) {
	if list == nil {
		list = make(map[string]func() Command)
	}
	if _, ok := list[name]; ok {
		panic("processer name conflict")
	}
	list[name] = command
}

var ErrorInvalidOptionFormat = "invalid %s option: %v"
