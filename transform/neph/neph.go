package neph

import (
	"errors"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
	"github.com/ctripcorp/nephele/transform/command"
)

// Transformer represents how to transform image with given commands
type Transformer struct {
	commands []command.GMCommand
}

func (t *Transformer) Accept(ctx *context.Context, name string, params map[string]string) error {
	var cmd command.GMCommand
	switch name {
	case command.RESIZE:
		cmd = &command.Resize{}
	case command.CROP:
		cmd = &command.Crop{}
	}

	if cmd != nil {
		if err := cmd.Verfiy(ctx, params); err != nil {
			return err
		}
		t.commands = append(t.commands, cmd)
	}
	return nil
}

//Transform original image blob to expected blob.
func (t *Transformer) Transform(ctx *context.Context, blob []byte) ([]byte, error) {
	wand, err := gm.NewMagickWand(blob)
	if err != nil {
		return nil, err
	}
	if ctx.Canceled() {
		return nil, errors.New("Timeout")
	}
	t.commands = append(t.commands, &command.Strip{})
	for _, command := range t.commands {
		if err := command.Exec(ctx, wand); err != nil {
			return nil, err
		}
		if ctx.Canceled() {
			return nil, errors.New("Timeout")
		}
	}
	return wand.WriteBlob()
}
