package sharpen

import (
	"fmt"
	"strconv"

	"context"

	"github.com/ctrip-nephele/gmagick"
	"github.com/ctripcorp/nephele/command"
)

type Command struct {
	Radius float64
	Sigma  float64
}

const (
	commandKeyR string = "r"
	commandKeyS string = "s"
)

func (c *Command) Support() string {
	return "wand"
}

//Verify sharpen Verify
func (c *Command) Verify(ctx context.Context, params map[string]string) error {
	//log.Debugf(ctx, "sharpen verification")
	for k, v := range params {
		if k == commandKeyR {
			radius, e := strconv.ParseFloat(v, 64)
			if e != nil {
				return fmt.Errorf(command.ErrorInvalidOptionFormat, k, v)
			}
			c.Radius = radius
		}
		if k == commandKeyS {
			sigma, e := strconv.ParseFloat(v, 64)
			if e != nil {
				return fmt.Errorf(command.ErrorInvalidOptionFormat, k, v)
			}
			c.Sigma = sigma
		}
	}
	return nil
}

func (c *Command) ExecuteOnBlob(ctx context.Context, blob []byte) ([]byte, error) {
	return nil, nil
}

//Exec sharpen exec
func (c *Command) ExecuteOnWand(ctx context.Context, wand *gmagick.MagickWand) error {
	//log.TraceBegin(ctx, "", "URL.Command", "sharpen", "radius", c.Radius, "sigma", c.Sigma)
	//defer log.TraceEnd(ctx, err)
	return wand.SharpenImage(c.Radius, c.Sigma)
}
