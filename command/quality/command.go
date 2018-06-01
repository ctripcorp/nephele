package quality

import (
	"fmt"
	"strconv"

	"context"

	"github.com/ctrip-nephele/gmagick"
	"github.com/ctripcorp/nephele/command"
)

type Command struct {
	Quality uint
}

const (
	qualityKeyV string = "v"
)

func (c *Command) Support() string {
	return "wand"
}

//Verify verify quality
func (c *Command) Verify(ctx context.Context, params map[string]string) error {
	//log.Debugf(ctx, "quality verification")
	for k, v := range params {
		if k == qualityKeyV {
			quality, e := strconv.Atoi(v)
			if e != nil || quality < 0 || quality > 100 {
				return fmt.Errorf(command.ErrorInvalidOptionFormat, k, v)
			}
			c.Quality = uint(quality)
			return nil
		}
	}
	return fmt.Errorf(command.ErrorInvalidOptionFormat, "quality", params)
}

func (c *Command) ExecuteOnBlob(ctx context.Context, blob []byte) ([]byte, error) {
	return nil, nil
}

//Exec exec
func (c *Command) ExecuteOnWand(ctx context.Context, wand *gmagick.MagickWand) error {
	var err error
	//log.TraceBegin(ctx, "", "URL.Command", "quality", "quality", q.Quality)
	//defer log.TraceEnd(ctx, err)
	if c.Quality == 0 {
		err = wand.SetImageOption("jpeg", "preserve-settings", "true")
	} else {
		err = wand.SetCompressionQuality(uint(c.Quality))
	}
	return err
}
