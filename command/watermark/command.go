package watermark

import (
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/ctripcorp/nephele/command"

	"github.com/ctrip-nephele/gmagick"
	"github.com/ctripcorp/nephele/storage"

	"context"

	"github.com/ctripcorp/nephele/util"
)

//Watermark watermark command
type Command struct {
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

func (c *Command) Support() string {
	return "wand"
}

//verify watermark verify
func (c *Command) Verify(ctx context.Context, params map[string]string) error {
	//log.Debugf(ctx, "watermark verification")
	for k, v := range params {
		if k == watermarkKeyN {
			vByte, e := base64.StdEncoding.DecodeString(v)
			if e != nil {
				return fmt.Errorf(command.ErrorInvalidOptionFormat, k, v)
			}
			if string(vByte) == "" {
				return fmt.Errorf("name of watermark must be provided")
			}
			c.Name = string(vByte)
			return nil
		}
		if k == watermarkKeyD {
			dissolve, e := strconv.Atoi(v)
			if e != nil || dissolve < 0 || dissolve > 100 {
				return fmt.Errorf(command.ErrorInvalidOptionFormat, k, v)
			}
			c.Dissolve = dissolve
			return nil
		}
		if k == watermarkKeyL {
			if !util.InArray(v, watermarkLocations) && v != "" {
				return fmt.Errorf(command.ErrorInvalidOptionFormat, k, v)
			}
			c.Location = v
			return nil
		}
		if k == watermarkKeyX {
			x, e := strconv.Atoi(v)
			if e != nil || x < 0 {
				return fmt.Errorf(command.ErrorInvalidOptionFormat, k, v)
			}
			c.X = x
			return nil
		}
		if k == watermarkKeyY {
			y, e := strconv.Atoi(v)
			if e != nil || y < 0 {
				return fmt.Errorf(command.ErrorInvalidOptionFormat, k, v)
			}
			c.Y = y
			return nil
		}
		if k == watermarkKeyMW {
			mw, e := strconv.Atoi(v)
			if e != nil || mw < 0 {
				return fmt.Errorf(command.ErrorInvalidOptionFormat, k, v)
			}
			c.Minwidth = uint(mw)
			return nil
		}
		if k == watermarkKeyMH {
			mh, e := strconv.Atoi(v)
			if e != nil || mh < 0 {
				return fmt.Errorf(command.ErrorInvalidOptionFormat, k, v)
			}
			c.Minheight = uint(mh)
			return nil
		}
	}
	return fmt.Errorf(command.ErrorInvalidOptionFormat, "watermark", params)
}

func (c *Command) ExecuteOnBlob(ctx context.Context, blob []byte) ([]byte, error) {
	return nil, nil
}

//Exec watermark exec
func (c *Command) ExecuteOnWand(ctx context.Context, wand *gmagick.MagickWand) error {
	var err error
	//log.TraceBegin(ctx, "", "URL.Command", "watermark", "watermarkName", c.Name, "location", c.Location, "dissolve", c.Dissolve, "x", c.X, "y", c.Y)
	//defer log.TraceEnd(ctx, err)
	if wand.GetImageWidth() < c.Minwidth || wand.GetImageHeight() < c.Minheight {
		return nil
	}
	var logoWand *gmagick.MagickWand
	logoWand, err = watermarkGetLogoWand(ctx, c.Name, c.Dissolve)
	if err != nil {
		return err
	}
	var x, y int
	if watermarkDecideLocationType(c.Location, c.X, c.Y) {
		x, y, err = watermarkGetCustomLocation(c.X, c.Y, wand, logoWand)
	} else {
		x, y, err = watermarkGetLocation(c.Location, wand, logoWand)
	}
	if err != nil {
		return err
	}
	err = wand.CompositeImage(logoWand, gmagick.COMPOSITE_OP_OVER, x, y)
	return err
}

//GetCustomLocation: get custom location by coordinate x,y
func watermarkGetCustomLocation(x, y int, wand, logo *gmagick.MagickWand) (int, int, error) {
	width := wand.GetImageWidth()
	height := wand.GetImageHeight()
	logowidth := logo.GetImageWidth()
	logoheight := logo.GetImageHeight()
	if x >= 0 && y >= 0 && uint(x)+logowidth < width && uint(y)+logoheight < height {
		return x, y, nil
	}
	return 0, 0, nil
}

//GetLocation: get location via Sudoku
func watermarkGetLocation(location string, wand, logo *gmagick.MagickWand) (int, int, error) {
	var (
		x uint = 0
		y uint = 0
	)
	width := wand.GetImageWidth()
	height := wand.GetImageHeight()
	logowidth := logo.GetImageWidth()
	logoheight := logo.GetImageHeight()
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
func watermarkGetLogoWand(ctx context.Context, watermarkName string, dissolve int) (*gmagick.MagickWand, error) {
	bt, _, err := storage.Download(ctx, watermarkName)
	if err != nil {
		println("error on watermarkGetLogoWand-storage.Download")
		fmt.Println("watermarkGetLogoWand error", err.Error())
		return nil, err
	}
	logoWand := gmagick.NewMagickWand()
	err = logoWand.ReadImageBlob(bt)
	if err != nil {
		return nil, err
	}
	if dissolve == 0 || dissolve == 100 {
		return logoWand, nil
	}
	logoWand.Dissolve(dissolve)
	if err != nil {
		return nil, err
	}
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
