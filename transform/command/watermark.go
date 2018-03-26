package command

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
	"github.com/ctripcorp/nephele/log"
	"github.com/ctripcorp/nephele/util"
)

//Watermark watermark command
type Watermark struct {
	Name                string
	Dissolve            int //1-100
	Location            string
	Minwidth, Minheight uint
	X, Y                int
}

const (
	watermarkKeyN  string = "n"
	watermarkKeyD  string = "d"
	watermarkKeyL  string = "l"
	watermarkKeyX  string = "x"
	watermarkKeyY  string = "y"
	watermarkKeyMW string = "mw"
	watermarkKeyMH string = "mh"
)

var watermarkLocations = []string{"nw", "north", "ne", "west", "center", "east", "sw", "south", "se"}

//verify watermark verify
func (w *Watermark) Verify(ctx *context.Context, params map[string]string) error {
	log.Debugw(ctx, "begin watermark verify")
	for k, v := range params {
		if k == watermarkKeyN {
			vByte, e := base64.StdEncoding.DecodeString(v)
			if e != nil {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			if string(vByte) == "" {
				return fmt.Errorf("name of watermark must be provided")
			}
			w.Name = string(vByte)
		}
		if k == watermarkKeyD {
			dissolve, e := strconv.Atoi(v)
			if e != nil || dissolve < 0 || dissolve > 100 {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			w.Dissolve = dissolve
		}
		if k == watermarkKeyL {
			if !util.InArray(v, watermarkLocations) && v != "" {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			w.Location = v
		}
		if k == watermarkKeyX {
			x, e := strconv.Atoi(v)
			if e != nil || x < 0 {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			w.X = x
		}
		if k == watermarkKeyY {
			y, e := strconv.Atoi(v)
			if e != nil || y < 0 {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			w.Y = y
		}
		if k == watermarkKeyMW {
			mw, e := strconv.Atoi(v)
			if e != nil || mw < 0 {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			w.Minwidth = uint(mw)
		}
		if k == watermarkKeyMH {
			mh, e := strconv.Atoi(v)
			if e != nil || mh < 0 {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			w.Minheight = uint(mh)
		}
	}
	return nil
}

//Exec watermark exec
func (w *Watermark) Exec(ctx *context.Context, wand *gm.MagickWand) error {
	log.TraceBegin(ctx, "watermark exec", "URL.Command", "watermark")
	defer log.TraceEnd(ctx, nil)
	if wand.Width() < w.Minwidth || wand.Height() < w.Minheight {
		return nil
	}
	logoWand, err := watermarkGetLogoWand(w.Name, w.Dissolve)
	if err != nil {
		return err
	}
	var x, y int
	if watermarkDecideLocationType(w.Location, w.X, w.Y) {
		x, y, err = watermarkGetCustomLocation(w.X, w.Y, wand, logoWand)
	} else {
		x, y, err = watermarkGetLocation(w.Location, wand, logoWand)
	}
	if err != nil {
		return err
	}
	return wand.Composite(logoWand, x, y)
}

//GetCustomLocation: get custom location by coordinate x,y
func watermarkGetCustomLocation(x, y int, wand, logo *gm.MagickWand) (int, int, error) {
	width, height, err := wand.Size()
	if err != nil {
		return 0, 0, err
	}
	logowidth, logoheight, err := logo.Size()
	if err != nil {
		return 0, 0, err
	}
	if x >= 0 && y >= 0 && uint(x)+logowidth < width && uint(y)+logoheight < height {
		return x, y, nil
	}
	return 0, 0, nil
}

//GetLocation: get location via Sudoku
func watermarkGetLocation(location string, wand, logo *gm.MagickWand) (int, int, error) {
	var (
		x uint = 0
		y uint = 0
	)
	width, height, err := wand.Size()
	if err != nil {
		return 0, 0, err
	}

	logowidth, logoheight, err := logo.Size()
	if err != nil {
		return 0, 0, err
	}
	switch location {
	case "nw":
		x, y = 0, 0
	case "north":
		x, y = (width-logowidth)/2, 0
	case "ne":
		x, y = width-logowidth, 0
	case "west":
		x, y = 0, (height-logoheight)/2
	case "center":
		x, y = (width-logowidth)/2, (height-logoheight)/2
	case "east":
		x, y = width-logowidth, (height-logoheight)/2
	case "sw":
		x, y = 0, height-logoheight
	case "south":
		x, y = (width-logowidth)/2, height-logoheight
	default:
		x, y = width-logowidth, height-logoheight
	}
	if x < 0 {
		x = 0
	}
	if y < 0 {
		y = 0
	}
	return int(x), int(y), nil
}

//GetLogoWand: get magickwand of logoImage
func watermarkGetLogoWand(wmName string, dissolve int) (*gm.MagickWand, error) {
	bt, err := ioutil.ReadFile(wmName)
	if err != nil {
		return nil, err
	}
	logoWand, err := gm.NewMagickWand(bt)
	if err != nil {
		return nil, err
	}
	if dissolve == 0 || dissolve == 100 {
		return logoWand, nil
	}
	logoWand.Dissolve(dissolve)
	return logoWand, nil
}

//DecideLocationType: decide function getLocation or function getCustomLocation to use
func watermarkDecideLocationType(location string, dstX, dstY int) bool {
	if location != "" {
		return false
	}
	if dstX != 0 || dstY != 0 {
		return true
	}
	return false
}
