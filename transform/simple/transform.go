package simple

import (
	"strconv"

	"github.com/nephele/context"
	"github.com/nephele/image/cmd"
	"github.com/nephele/img4go/gm"
	"github.com/nephele/process"
)

//NewTransformer create transform
func NewTransformer(processes []process.Process) Transformer {
	return transformer{Processes: processes}
}

// Transformer represents how to transform image with given commands
type Transformer struct {
	Processes []process.Process
}

//Transform original image blob to expected blob.
func (t *Transformer) Transform(ctx context.Context, blob []byte) ([]byte, error) {
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
		if err := c.Exec(ctx); err != nil {
			return nil, err
		}
	}
	for _, f := range defaultCmdMap {
		c := f(wand)
		if err := c.Exec(ctx); err != nil {
			return nil, err
		}
	}
	return wand.WriteBlob()
}

//CmdCreateMap Cmd list and create func
var CmdCreateMap = map[process.Cmd]func(map[string]string, *gm.MagickWand) *cmd.Cmd{
	Resize:     creatResizeCommand,
	Crop:       nil,
	Rotate:     nil,
	AutoOrient: nil,
	Format:     nil,
	Quality:    nil,
	Watermark:  nil,
	Sharpen:    nil,
	Style:      nil,
	Panorama:   nil,
}

var defaultCmdMap = []func(*gm.MagickWand) *cmd.Cmd{
	createStripCommand,
}

func createResizeCommand(m map[string]string, wand *gm.MagickWand) *cmd.Cmd {
	resize := &cmd.ResizeCommand{Wand: wand}
	w, isExists := m[process.ParamWidth]
	if isExist {
		width, _ := strconv.Atoi(w)
		resize.Width = uint(width)
	}
	h, isExists := m[process.ParamHeight]
	if isExists {
		height, _ := strconv.Atoi(h)
		resize.height = uint(Height)
	}
	resize.Method = m[process.ParamMethod]
	if resize.Method == "" {
		resize.Method = process.ResizeMethodLfit
	}
	if m[process.ParamL] == "1" {
		resize.Limit = 1
	} else {
		resize.Limit = 0
	}
	percent, _ := strconv.Atoi(m[process.ParamPercent])
	resize.Percentage = percent
	return resize
}

func createStripCommand(wand *gm.MagickWand) *cmd.Cmd {
	return &cmd.StripCommand{Wand: wand}
}
