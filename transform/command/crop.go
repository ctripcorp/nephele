package command

import (
	"fmt"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
	"github.com/ctripcorp/nephele/log"
)

type Crop struct {
	Wand              *gm.MagickWand
	Width             uint
	Height            uint
	Method            string
	X, Y              uint
	Limit, Percentage int
}

// const (
// 	C      string = "c"
// 	T      string = "t"
// 	B      string = "b"
// 	L      string = "l"
// 	R      string = "r"
// 	WC     string = "wc"
// 	HC     string = "hc"
// 	RESIZE string = "resize"
// 	CROP   string = "crop"
// )

//Exec crop exec
func (c *Crop) Exec(ctx *context.Context) error {
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
