package command

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
	"github.com/ctripcorp/nephele/log"
)

type Crop struct {
	Width      uint
	Height     uint
	Method     string
	X, Y       int
	Limit      bool
	Percentage int
}

const (
	cropKeyW     string = "w"
	cropKeyH     string = "h"
	cropKeyM     string = "m"
	cropKeyP     string = "p"
	cropKeyX     string = "x"
	cropKeyY     string = "y"
	cropKeyLimit string = "limit"
)

const (
	cropKeyMT      string = "t"
	cropKeyMB      string = "b"
	cropKeyML      string = "l"
	cropKeyMR      string = "r"
	cropKeyMWC     string = "wc"
	cropKeyMHC     string = "hc"
	cropKeyMC      string = "c"
	cropKeyMRESIZE string = "resize"
	cropKeyMCROP   string = "crop"
)

//Verify crop Verify
func (c *Crop) Verify(ctx *context.Context, params map[string]string) error {
	log.Debugf(ctx, "crop Verify")
	for k, v := range params {
		if k == cropKeyW {
			width, e := strconv.Atoi(v)
			if e != nil {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			c.Width = uint(width)
		}
		if k == cropKeyH {
			height, e := strconv.Atoi(v)
			if e != nil {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			c.Height = uint(height)
		}
		if k == cropKeyM {
			if v != cropKeyMT && v != cropKeyMB && v != cropKeyMC && v != cropKeyMCROP &&
				v != cropKeyMHC && v != cropKeyML && v != cropKeyMR && v != cropKeyMRESIZE && v != cropKeyMWC {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			c.Method = v
		}
		if k == cropKeyLimit {
			if v != "0" && v != "1" {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			c.Limit = v == "1"
		}
		if k == cropKeyP {
			p, e := strconv.Atoi(v)
			if e != nil || p < 0 || p > 10000 {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			c.Percentage = p
		}
		if k == cropKeyX {
			x, e := strconv.Atoi(v)
			if e != nil {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			c.X = x
		}
		if k == cropKeyY {
			y, e := strconv.Atoi(v)
			if e != nil {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			c.Y = y
		}
	}
	if c.Percentage < 0 || c.Percentage >= 100 {
		return fmt.Errorf(invalidInfoFormat, c.Percentage, cropKeyP)
	}
	if (c.Method == cropKeyMT || c.Method == cropKeyMB || c.Method == cropKeyMHC) &&
		c.Height < 1 && c.Percentage < 1 {
		return errors.New("m,h,p is invalid.")
	}
	if (c.Method == cropKeyML || c.Method == cropKeyMR || c.Method == cropKeyMWC) &&
		c.Width < 1 && c.Percentage < 1 {
		return errors.New("m,w,p is invalid.")
	}
	if c.Method == cropKeyMC && c.Percentage < 1 && c.Width < 1 && c.Height < 1 {
		return errors.New("m,w,h,p is invalid.")
	}
	if (c.Method == cropKeyMRESIZE || c.Method == cropKeyMCROP) &&
		c.Percentage < 1 && (c.Width < 1 || c.Height < 1) {
		return errors.New("m,w,h,p is invalid.")
	}
	return nil
}

//Exec crop exec
func (c *Crop) Exec(ctx *context.Context, wand *gm.MagickWand) error {
	log.TraceBegin(ctx, fmt.Sprintf("crop, method:%s, width:%d,height:%d,x:%d,y:%d,p:%d, limit:%t", c.Method, c.Width, c.Height, c.X, c.Y, c.Percentage, c.Limit),
		"URL.Command", "crop")
	defer log.TraceEnd(ctx, nil)

	srcW, srcH := wand.Width(), wand.Height()
	var width, height uint
	var x, y int
	if c.Method == cropKeyMRESIZE {
		var isResize bool
		width, height, x, y, isResize = cropMRESIZE(c.Width, c.Height, srcW, srcH, c.Limit)
		if width == 0 && height == 0 && x == 0 && y == 0 {
			return nil
		}
		fmt.Println(width, height, x, y, isResize)
		if !(width == srcW && height == srcH && x == 0 && y == 0) {
			if err := wand.Crop(width, height, x, y); err != nil {
				return err
			}
			if !isResize {
				return nil
			}
			if err := wand.LanczosResize(c.Width, c.Height); err != nil {
				return err
			}
		}
		return nil
	}

	switch c.Method {
	case cropKeyMB:
		width, height, x, y = cropMB(c.Height, srcW, srcH, c.Percentage)
	case cropKeyMC:
		width, height, x, y = cropMC(c.Width, c.Height, srcW, srcH, c.Percentage)
	case cropKeyMCROP:
		width, height = cropMCROP(c.Width, c.Height, srcW, srcH, c.Percentage)
		x, y = c.X, c.Y
	case cropKeyMHC:
		width, height, x, y = cropMHC(c.Height, srcW, srcH, c.Percentage)
	case cropKeyML:
		width, height, x, y = cropML(c.Width, srcW, srcH, c.Percentage)
	case cropKeyMR:
		width, height, x, y = cropMR(c.Width, srcW, srcH, c.Percentage)
	case cropKeyMT:
		width, height, x, y = cropMT(c.Height, srcW, srcH, c.Percentage)
	case cropKeyMWC:
		width, height, x, y = cropMWC(c.Width, srcW, srcH, c.Percentage)
	}
	if width < 1 || height < 1 || x >= int(srcW) || y >= int(srcH) {
		return errors.New("param is invalid.")
	}
	return wand.Crop(width, height, x, y)
}

func cropMRESIZE(w, h, srcW, srcH uint, limit bool) (width, height uint, x, y int, resize bool) {
	resize = false
	if !limit && w > srcW && h > srcH {
		return
	}

	width = srcW
	height = srcH
	if !limit && !(w < srcW && h < srcH) {
		if srcW > w {
			x = int((srcW - w) / 2)
			width = w
			return
		}
		if srcH > h {
			y = int((srcH - h) / 2)
			height = h
			return
		}
	}
	resize = true
	dstP := float64(w) / float64(h)
	srcP := float64(srcW) / float64(srcH)
	if math.Abs(dstP-srcP) > 0.0001 {
		if srcP > dstP { //以高缩小
			height = srcH
			width = uint(math.Floor(float64(height) * dstP))
			x = int((srcW - width) / 2)
		}
		if srcP < dstP { //以宽缩小
			width = srcW
			height = uint(math.Floor(float64(width) / dstP))
			y = int((srcH - height) / 2)
		}
	}
	return
}

func cropMT(h, srcW, srcH uint, p int) (width, height uint, x, y int) {
	if p > 0 {
		h = srcH * uint(p) / 100
	}
	width = srcW
	height = srcH - h
	y = int(h)
	return
}

func cropMB(h, srcW, srcH uint, p int) (width, height uint, x, y int) {
	if p > 0 {
		h = srcH * uint(p) / 100
	}
	width = srcW
	height = srcH - h
	return
}

func cropMHC(h, srcW, srcH uint, p int) (width, height uint, x, y int) {
	if p > 0 {
		h = srcH * uint(p) / 100
	}
	width = srcW
	height = srcH - h
	y = int(h) / 2
	return
}

func cropML(w, srcW, srcH uint, p int) (width, height uint, x, y int) {
	if p > 0 {
		w = srcW * uint(p) / 100
	}
	width = srcW - w
	height = srcH
	x = int(w)
	return
}

func cropMR(w, srcW, srcH uint, p int) (width, height uint, x, y int) {
	if p > 0 {
		w = srcW * uint(p) / 100
	}
	width = srcW - w
	height = srcH
	return
}

func cropMWC(w, srcW, srcH uint, p int) (width, height uint, x, y int) {
	if p > 0 {
		w = srcW * uint(p) / 100
	}
	width = srcW - w
	height = srcH
	x = int(w) / 2
	return
}

func cropMC(w, h, srcW, srcH uint, p int) (width, height uint, x, y int) {
	if p > 0 {
		w = srcW * uint(p) / 100
		h = srcH * uint(p) / 100
	}

	width, height = srcW, srcH
	if srcW > w && w != 0 {
		width = w
		x = int((srcW - w) / 2)
	}
	if srcH > h && h != 0 {
		height = h
		y = int((srcH - h) / 2)
	}
	return
}

func cropMCROP(w, h, srcW, srcH uint, p int) (width, height uint) {
	if p > 0 {
		w = srcW * uint(p) / 100
		h = srcH * uint(p) / 100
	}
	if w == 0 {
		w = srcW
	}
	if h == 0 {
		h = srcH
	}
	width = w
	height = h
	return
}
