package rotate

import (
	"fmt"
	"strconv"

	"context"

	"github.com/ctrip-nephele/gmagick"
	"github.com/ctripcorp/nephele/command"
)

//Rotate rotate command
type Command struct {
	Degree int
}

const (
	rotateV string = "v"
)

func (c *Command) Support() string {
	return "wand"
}

//Verify rotate verify params
func (c *Command) Verify(ctx context.Context, params map[string]string) error {
	//log.Debugf(ctx, "rotate verification")
	for k, v := range params {
		if k == rotateV {
			degree, e := strconv.Atoi(v)
			if e != nil || degree < 0 || degree > 360 {
				return fmt.Errorf(command.ErrorInvalidOptionFormat, k, v)
			}
			c.Degree = degree
			return nil
		}
	}
	return fmt.Errorf(command.ErrorInvalidOptionFormat, "rotate", params)
}

func (c *Command) ExecuteOnBlob(ctx context.Context, blob []byte) ([]byte, error) {
	return nil, nil
}

//Exec rotate ExecOnWand
func (c *Command) ExecuteOnWand(ctx context.Context, wand *gmagick.MagickWand) error {
	var err error
	//log.TraceBegin(ctx, "", "URL.Command", "rotate", "degree", c.Degree)
	//defer log.TraceEnd(ctx, err)
	background := gmagick.NewPixelWand()
	background.SetColor("#000000")
	err = wand.RotateImage(background, float64(c.Degree))
	return err
}
