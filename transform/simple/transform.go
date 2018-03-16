package simple

import (
	"github.com/nephele/context"
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
		if !isExists{
			continue
		}
		cmd := f(proc.Param)
		
	}
	return wand.WriteBlob()
}

//CmdCreateMap Cmd list and create func
var CmdCreateMap = map[Cmd]func(map[string]string) interface{
	Resize:     creatResizeCmd,
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

func createResizeCmd(m map[string]string)interface{
	return nil
}
