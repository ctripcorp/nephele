package command

import (
	"fmt"
	"strconv"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
	"github.com/ctripcorp/nephele/log"
)

type Crop struct {
	Width      uint
	Height     uint
	Method     string
	X, Y       uint
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

//Verfiy crop verfiy
func (r *Crop) Verfiy(ctx *context.Context, params map[string]string) error {
	log.Debugw(ctx, "begin crop verfiy")
	for k, v := range params {
		if k == cropKeyW {
			width, e := strconv.Atoi(v)
			if e != nil {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			r.Width = uint(width)
		}
		if k == cropKeyH {
			height, e := strconv.Atoi(v)
			if e != nil {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			r.Height = uint(height)
		}
		if k == cropKeyM {
			if v != cropKeyMT && v != cropKeyMB && v != cropKeyMC && v != cropKeyMCROP &&
				v != cropKeyMHC && v != cropKeyML && v != cropKeyMR {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			r.Method = v
		}
		if k == cropKeyLimit {
			if v != "0" && v != "1" {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			r.Limit = v == "1"
		}
		if k == cropKeyP {
			p, e := strconv.Atoi(v)
			if e != nil || p < 0 || p > 10000 {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			r.Percentage = p
		}
	}
	return nil
}

//Exec crop exec
func (c *Crop) Exec(ctx *context.Context, wand *gm.MagickWand) error {
	log.TraceBegin(ctx, fmt.Sprintf("crop, method:%s, width:%d,height:%d,x:%d,y:%d,p:%d, limit:%d", c.Method, c.Width, c.Height, c.X, c.Y, c.Percentage, c.Limit),
		"URL.Command", "crop")
	defer log.TraceEnd(ctx, nil)
	// if c.Method == C {

	// }
	return nil
}

func cropMT(h, srcW, srcH uint, p int) (width, height uint, x, y int) {
	if p > 0 {
		h = srcH * uint(p) / 100
	}

	return
}

func cropMC(w, h, srcW, srcH uint, p int) (width, height uint, x, y int) {
	if p > 0 {
		w = srcW * uint(p) / 100
		h = srcH * uint(p) / 100
	}

	width, height = srcW, srcH
	if srcW > w {
		width = w
		x = int((srcW - w) / 2)
	}
	if srcH > h {
		height = h
		y = int((srcH - h) / 2)
	}
	return
}
