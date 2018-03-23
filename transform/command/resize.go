package command

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
)

//Resize resize command
type Resize struct {
	Width      uint
	Height     uint
	Method     string //lfit/fixed
	Limit      bool
	Percentage int
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
func (r *Resize) Verify(ctx *context.Context, params map[string]string) error {
	//log.Debugw(ctx, "begin resize Verify")
	for k, v := range params {
		if k == resizeKeyW {
			width, e := strconv.Atoi(v)
			if e != nil {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			r.Width = uint(width)
		}
		if k == resizeKeyH {
			height, e := strconv.Atoi(v)
			if e != nil {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			r.Height = uint(height)
		}
		if k == resizeKeyM {
			if v != resizeKeyMFIXED && v != resizeKeyMLFIT {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			r.Method = v
		}
		if k == resizeKeyLimit {
			if v != "0" && v != "1" {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			r.Limit = v == "1"
		}
		if k == resizeKeyP {
			p, e := strconv.Atoi(v)
			if e != nil || p < 0 || p > 10000 {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			r.Percentage = p
		}
	}
	if r.Method == resizeKeyMFIXED && r.Width < 1 && r.Height < 1 {
		return errors.New("m, w, h is invalid.")
	}
	if r.Width < 1 && r.Height < 1 && r.Percentage < 1 {
		return errors.New("w,h,p is invalid.")
	}
	return nil
}

//Exec resize  exec
func (r *Resize) Exec(ctx *context.Context, wand *gm.MagickWand) error {
	//log.TraceBegin(ctx, "resize exec", "URL.Command", "resize")
	//defer log.TraceEnd(ctx, nil)

	if (r.Width > wand.Width() && r.Height > wand.Height() && !r.Limit) ||
		(r.Method != resizeKeyMFIXED && r.Percentage > 100 && !r.Limit) {
		return nil
	}
	srcW, srcH := wand.Width(), wand.Height()
	if srcW == r.Width && srcH == r.Height {
		return nil
	}
	var width, height uint
	if r.Method == resizeKeyMFIXED {
		width, height = resizeFixed(r.Width, r.Height, srcW, srcH)
	} else if r.Percentage != 0 {
		width, height = resizePercentage(r.Percentage, srcW, srcH)
	} else {
		width, height = resizeLfit(r.Width, r.Height, srcW, srcH)
	}
	return wand.LanczosResize(width, height)
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
