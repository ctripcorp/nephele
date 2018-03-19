package simple

import (
	"strconv"

	"github.com/nephele/context"
	"github.com/nephele/image/cmd"
	"github.com/nephele/img4go/gm"
	"github.com/nephele/process"
)

// Transformer represents how to transform image with given commands
type Transformer struct {
	Processes []process.Process
}

//Transform original image blob to expected blob.
func (t *Transformer) Transform(ctx *context.Context, blob []byte) ([]byte, error) {
	wand, err := gm.NewMagickWand(blob)
	if err != nil {
		return nil, err
	}
	for _, proc := range t.Processes {
		f, isExists := CmdCreateMap[proc.Name]
		if !isExists || f == nil {
			continue
		}
		c := f(proc.Param, wand)
		if err := c.Exec(*ctx); err != nil {
			return nil, err
		}
	}
	for _, f := range defaultCmdMap {
		c := f(wand)
		if err := c.Exec(*ctx); err != nil {
			return nil, err
		}
	}
	return wand.WriteBlob()
}

//CmdCreateMap Cmd list and create func
var CmdCreateMap = map[process.Cmd]func(map[process.Key]string, *gm.MagickWand) cmd.Cmd{
	process.Resize:     createResizeCommand,
	process.Crop:       nil,
	process.Rotate:     nil,
	process.AutoOrient: nil,
	process.Format:     nil,
	process.Quality:    nil,
	process.Watermark:  nil,
	process.Sharpen:    nil,
	process.Style:      nil,
	process.Panorama:   nil,
}

var defaultCmdMap = []func(*gm.MagickWand) cmd.Cmd{
	createStripCommand,
}

func createResizeCommand(m map[process.Key]string, wand *gm.MagickWand) cmd.Cmd {
	resize := &cmd.ResizeCommand{Wand: wand}
	w, isExists := m[process.KeyW]
	if isExists {
		width, _ := strconv.Atoi(w)
		resize.Width = uint(width)
	}
	h, isExists := m[process.KeyH]
	if isExists {
		height, _ := strconv.Atoi(h)
		resize.Height = uint(height)
	}
	resize.Method = m[process.KeyM]
	if resize.Method == "" {
		resize.Method = process.ResizeMethodLfit
	}
	if m[process.KeyL] == "1" {
		resize.Limit = 1
	} else {
		resize.Limit = 0
	}
	percent, _ := strconv.Atoi(m[process.KeyP])
	resize.Percentage = percent
	return resize
}

func createStripCommand(wand *gm.MagickWand) cmd.Cmd {
	return &cmd.StripCommand{Wand: wand}
}
