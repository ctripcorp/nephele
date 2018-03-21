package command

import (
	"fmt"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
	"github.com/ctripcorp/nephele/log"
)

type Crop struct {
	Width             uint
	Height            uint
	Method            string
	X, Y              uint
	Limit, Percentage int
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

func methodT(h, srcW, srcH uint, p int) (width, height uint, x, y int) {
	if p > 0 {
		h = srcH * uint(p) / 100
	}

	return
}

func methodC(w, h, srcW, srcH uint, p int) (width, height uint, x, y int) {
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
