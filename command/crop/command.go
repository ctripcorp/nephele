package crop

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"context"

	"github.com/ctrip-nephele/gmagick"
	"github.com/ctripcorp/nephele/command"
)

type Command struct {
	Width      uint
	Height     uint
	Method     string
	X, Y       int
	Limit      bool
	Percentage int
}

const (
	commandKeyW     string = "w"
	commandKeyH     string = "h"
	commandKeyM     string = "m"
	commandKeyP     string = "p"
	commandKeyX     string = "x"
	commandKeyY     string = "y"
	commandKeyLimit string = "limit"
)

const (
	commandKeyMT      string = "t"
	commandKeyMB      string = "b"
	commandKeyML      string = "l"
	commandKeyMR      string = "r"
	commandKeyMWC     string = "wc"
	commandKeyMHC     string = "hc"
	commandKeyMC      string = "c"
	commandKeyMRESIZE string = "resize"
	commandKeyMCROP   string = "crop"
)

func (c *Command) Support() string {
	return "wand"
}

//Verify crop Verify
func (c *Command) Verify(ctx context.Context, params map[string]string) error {
	//log.Debugf(ctx, "crop verification")
	for k, v := range params {
		if k == commandKeyW {
			width, e := strconv.Atoi(v)
			if e != nil {
				return fmt.Errorf(command.ErrorInvalidOptionFormat, k, v)
			}
			c.Width = uint(width)
		}
		if k == commandKeyH {
			height, e := strconv.Atoi(v)
			if e != nil {
				return fmt.Errorf(command.ErrorInvalidOptionFormat, k, v)
			}
			c.Height = uint(height)
		}
		if k == commandKeyM {
			if v != commandKeyMT && v != commandKeyMB && v != commandKeyMC && v != commandKeyMCROP &&
				v != commandKeyMHC && v != commandKeyML && v != commandKeyMR && v != commandKeyMRESIZE && v != commandKeyMWC {
				return fmt.Errorf(command.ErrorInvalidOptionFormat, k, v)
			}
			c.Method = v
		}
		if k == commandKeyLimit {
			if v != "0" && v != "1" {
				return fmt.Errorf(command.ErrorInvalidOptionFormat, k, v)
			}
			c.Limit = v == "1"
		}
		if k == commandKeyP {
			p, e := strconv.Atoi(v)
			if e != nil || p < 0 || p > 10000 {
				return fmt.Errorf(command.ErrorInvalidOptionFormat, k, v)
			}
			c.Percentage = p
		}
		if k == commandKeyX {
			x, e := strconv.Atoi(v)
			if e != nil {
				return fmt.Errorf(command.ErrorInvalidOptionFormat, k, v)
			}
			c.X = x
		}
		if k == commandKeyY {
			y, e := strconv.Atoi(v)
			if e != nil {
				return fmt.Errorf(command.ErrorInvalidOptionFormat, k, v)
			}
			c.Y = y
		}
	}
	if c.Percentage < 0 || c.Percentage >= 100 {
		return fmt.Errorf(command.ErrorInvalidOptionFormat, commandKeyP, c.Percentage)
	}
	if (c.Method == commandKeyMT || c.Method == commandKeyMB || c.Method == commandKeyMHC) &&
		c.Height < 1 && c.Percentage < 1 {
		return errors.New("m,h,p is invalid.")
	}
	if (c.Method == commandKeyML || c.Method == commandKeyMR || c.Method == commandKeyMWC) &&
		c.Width < 1 && c.Percentage < 1 {
		return errors.New("m,w,p is invalid.")
	}
	if c.Method == commandKeyMC && c.Percentage < 1 && c.Width < 1 && c.Height < 1 {
		return errors.New("m,w,h,p is invalid.")
	}
	if (c.Method == commandKeyMRESIZE || c.Method == commandKeyMCROP) &&
		c.Percentage < 1 && (c.Width < 1 || c.Height < 1) {
		return errors.New("m,w,h,p is invalid.")
	}
	return nil
}

func (c *Command) ExecuteOnBlob(ctx context.Context, blob []byte) ([]byte, error) {
	return nil, nil
}

//Exec crop exec
func (c *Command) ExecuteOnWand(ctx context.Context, wand *gmagick.MagickWand) error {
	var err error
	//log.TraceBegin(ctx, "", "URL.Command", "crop", "method", c.Method, "width", c.Width, "height", c.Height, "x", c.X, "y", c.Y)
	//defer log.TraceEnd(ctx, err)

	srcW, srcH := wand.GetImageWidth(), wand.GetImageHeight()
	var width, height uint
	var x, y int
	if c.Method == commandKeyMRESIZE {
		var isResize bool
		width, height, x, y, isResize = cropMRE(c.Width, c.Height, srcW, srcH, c.Limit)
		if width == 0 && height == 0 && x == 0 && y == 0 {
			return nil
		}
		fmt.Println(width, height, x, y, isResize)
		if !(width == srcW && height == srcH && x == 0 && y == 0) {
			if err = wand.CropImage(width, height, x, y); err != nil {
				return err
			}
			if !isResize {
				return nil
			}
			if err = wand.ResizeImage(width, height, gmagick.FILTER_LANCZOS, 1.0); err != nil {
				return err
			}
		}
		return nil
	}

	switch c.Method {
	case commandKeyMB:
		width, height, x, y = cropMB(c.Height, srcW, srcH, c.Percentage)
	case commandKeyMC:
		width, height, x, y = cropMC(c.Width, c.Height, srcW, srcH, c.Percentage)
	case commandKeyMCROP:
		width, height = cropMCR(c.Width, c.Height, srcW, srcH, c.Percentage)
		x, y = c.X, c.Y
	case commandKeyMHC:
		width, height, x, y = cropMHC(c.Height, srcW, srcH, c.Percentage)
	case commandKeyML:
		width, height, x, y = cropML(c.Width, srcW, srcH, c.Percentage)
	case commandKeyMR:
		width, height, x, y = cropMR(c.Width, srcW, srcH, c.Percentage)
	case commandKeyMT:
		width, height, x, y = cropMT(c.Height, srcW, srcH, c.Percentage)
	case commandKeyMWC:
		width, height, x, y = cropMWC(c.Width, srcW, srcH, c.Percentage)
	}
	if width < 1 || height < 1 || x >= int(srcW) || y >= int(srcH) {
		err = errors.New("param is invalid.")
		return err
	}
	err = wand.CropImage(width, height, x, y)
	return err
}

// crop,m_resize,w_100,h_100 , first equal comparison cut image, finally resize
func cropMRE(w, h, srcW, srcH uint, limit bool) (width, height uint, x, y int, resize bool) {
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

//crop,m_t,h_100  crop top height 100
func cropMT(h, srcW, srcH uint, p int) (width, height uint, x, y int) {
	if p > 0 {
		h = srcH * uint(p) / 100
	}
	width = srcW
	height = srcH - h
	y = int(h)
	return
}

//crop,m_b,h_100 , crop bottom
func cropMB(h, srcW, srcH uint, p int) (width, height uint, x, y int) {
	if p > 0 {
		h = srcH * uint(p) / 100
	}
	width = srcW
	height = srcH - h
	return
}

//crop,m_hc,h_100 ,  cut the upper and lower sides
func cropMHC(h, srcW, srcH uint, p int) (width, height uint, x, y int) {
	if p > 0 {
		h = srcH * uint(p) / 100
	}
	width = srcW
	height = srcH - h
	y = int(h) / 2
	return
}

//crop,m_l,w_100,  crop left
func cropML(w, srcW, srcH uint, p int) (width, height uint, x, y int) {
	if p > 0 {
		w = srcW * uint(p) / 100
	}
	width = srcW - w
	height = srcH
	x = int(w)
	return
}

//crop,m_r,w_100,  crop right
func cropMR(w, srcW, srcH uint, p int) (width, height uint, x, y int) {
	if p > 0 {
		w = srcW * uint(p) / 100
	}
	width = srcW - w
	height = srcH
	return
}

//crop,m_wc,w_100, cut the left and right sides
func cropMWC(w, srcW, srcH uint, p int) (width, height uint, x, y int) {
	if p > 0 {
		w = srcW * uint(p) / 100
	}
	width = srcW - w
	height = srcH
	x = int(w) / 2
	return
}

//crop,m_c,w_100 cut image for center
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

// crop,m_crop,w_100, crop cut image
func cropMCR(w, h, srcW, srcH uint, p int) (width, height uint) {
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
