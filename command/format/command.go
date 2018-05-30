package format

import (
	"context"
	"fmt"
	"github.com/ctrip-nephele/gmagick"
	"github.com/ctripcorp/nephele/command"
	"github.com/ctripcorp/nephele/util"
)

type Command struct {
	format string
}

func (c *Command) Support() string {
	return "wand"
}

func (c *Command) Verify(ctx context.Context, option map[string]string) error {
	//log.Debugf(ctx, "format verification")
	for k, v := range option {
		if k == "v" {
			if !util.InArray(v, []string{"jpg", "png", "webp", "gif"}) {
				return fmt.Errorf(command.ErrorInvalidOptionFormat, "format", option)
			}
			c.format = v
			return nil
		}
	}
	return fmt.Errorf(command.ErrorInvalidOptionFormat, "format", option)
}

func (c *Command) ExecuteOnBlob(ctx context.Context, blob []byte) ([]byte, error) {
	return nil, nil
}

func (c *Command) ExecuteOnWand(ctx context.Context, wand *gmagick.MagickWand) error {
	var err error
	//log.TraceBegin(ctx, "execute on wand", "url.command", "format", "format", c.format)
	//defer log.TraceEnd(ctx, err)
	err = wand.SetImageFormat(c.format)
	return err
}
