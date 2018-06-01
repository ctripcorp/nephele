package resize

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"context"

	"github.com/ctrip-nephele/gmagick"
	"github.com/ctripcorp/nephele/command"
)

//Resize resize command
type Command struct {
	Width      uint
	Height     uint
	Method     string //lfit/fixed
	Limit      bool
	Percentage int
}

func (c *Command) Support() string {
	return "wand"
}

const (
	resizeKeyW     string = "w"
	resizeKeyH     string = "h"
	resizeKeyM     string = "m"
	resizeKeyP     string = "p"
	resizeKeyLimit string = "limit"
)
const (
	resizeKeyMFIXED string = "fixed"
	resizeKeyMLFIT  string = "lfit"
)

//Verify resize Verify
func (c *Command) Verify(ctx context.Context, params map[string]string) error {
	//log.Debugf(ctx, "resize verification")
	for k, v := range params {
		if k == resizeKeyW {
			width, e := strconv.Atoi(v)
			if e != nil {
				return fmt.Errorf(command.ErrorInvalidOptionFormat, k, v)
			}
			c.Width = uint(width)
			return nil
		}
		if k == resizeKeyH {
			height, e := strconv.Atoi(v)
			if e != nil {
				return fmt.Errorf(command.ErrorInvalidOptionFormat, k, v)
			}
			c.Height = uint(height)
			return nil
		}
		if k == resizeKeyM {
			if v != resizeKeyMFIXED && v != resizeKeyMLFIT {
				return fmt.Errorf(command.ErrorInvalidOptionFormat, k, v)
			}
			c.Method = v
			return nil
		}
		if k == resizeKeyLimit {
			if v != "0" && v != "1" {
				return fmt.Errorf(command.ErrorInvalidOptionFormat, k, v)
			}
			c.Limit = v == "1"
			return nil
		}
		if k == resizeKeyP {
			p, e := strconv.Atoi(v)
			if e != nil || p < 0 || p > 10000 {
				return fmt.Errorf(command.ErrorInvalidOptionFormat, k, v)
			}
			c.Percentage = p
			return nil
		}
	}
	if c.Method == resizeKeyMFIXED && c.Width < 1 && c.Height < 1 {
		return errors.New("m, w, h is invalid.")
	}
	if c.Width < 1 && c.Height < 1 && c.Percentage < 1 {
		return errors.New("w,h,p is invalid.")
	}
	return fmt.Errorf(command.ErrorInvalidOptionFormat, "resize", params)
}

//ExecuteOnBlob
func (c *Command) ExecuteOnBlob(ctx context.Context, blob []byte) ([]byte, error) {
	return nil, nil
}

//Exec resize  exec
func (c *Command) ExecuteOnWand(ctx context.Context, wand *gmagick.MagickWand) error {
	var err error
	//log.TraceBegin(ctx, "", "URL.Command", "resize", "method", c.Method, "width", c.Width, "height", c.Height, "percentage", c.Percentage, "limit", c.Limit)
	//defer log.TraceEnd(ctx, err)
	if (c.Width > wand.GetImageWidth() && c.Height > wand.GetImageHeight() && !c.Limit) ||
		(c.Method != resizeKeyMFIXED && c.Percentage > 100 && !c.Limit) {
		return nil
	}
	srcW, srcH := wand.GetImageWidth(), wand.GetImageHeight()
	if srcW == c.Width && srcH == c.Height {
		return nil
	}
	var width, height uint
	if c.Method == resizeKeyMFIXED {
		width, height = resizeFixed(c.Width, c.Height, srcW, srcH)
	} else if c.Percentage != 0 {
		width, height = resizePercentage(c.Percentage, srcW, srcH)
	} else {
		width, height = resizeLfit(c.Width, c.Height, srcW, srcH)
	}
	err = wand.ResizeImage(width, height, gmagick.FILTER_LANCZOS, 1.0)
	return err
}

//Lfit: isotropic scaling with fixed width and height, which tends to disable one of the inputs(width or height) to feed a larger aspect ratio
func resizeLfit(dstW, dstH, srcW, srcH uint) (width, height uint) {
	//auto compute weight or height
	if dstW == 0 {
		width = dstH * srcW / srcH
		height = dstH
		return
	}
	if dstH == 0 {
		width = dstW
		height = dstW * srcH / srcW
		return
	}

	dstP := float64(dstW) / float64(dstH)
	srcP := float64(srcW) / float64(srcH)

	if srcP > dstP {
		width = dstW
		if uint(math.Abs(float64(dstH-srcH))) < 3 {
			height = dstH
		} else {
			height = uint(math.Floor(float64(dstW) / srcP))
		}
	} else {
		height = dstH
		if uint(math.Abs(float64(dstW-srcW))) < 3 {
			width = dstW
		} else {
			width = uint(math.Floor(float64(dstH) * srcP))
		}
	}
	return
}

//Fixed: forced scaling with fixed width and height
func resizeFixed(dstW, dstH, srcW, srcH uint) (width, height uint) {
	width, height = dstW, dstH
	if dstW < 1 {
		width = srcW
	}
	if dstH < 1 {
		height = srcH
	}
	return
}

//Percentage: isotropic scaling by multiplicator(%)
func resizePercentage(p int, srcW, srcH uint) (width, height uint) {
	width = srcW * uint(p) / 100
	height = srcH * uint(p) / 100
	return
}
