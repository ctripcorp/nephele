package command

import (
	"fmt"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
	"github.com/ctripcorp/nephele/log"
	"github.com/ctripcorp/nephele/util"
)

//Strip strip command
type Format struct {
	format string
}

const (
	formatV string = "v"
)

var formats = []string{"jpg", "png", "webp", "gif"}

//Verify format Verify params
func (f *Format) Verify(ctx *context.Context, params map[string]string) error {
	log.Debugf(ctx, "format verification")
	for k, v := range params {
		if k == formatV {
			if !util.InArray(v, formats) {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			f.format = v
		}
	}
	return nil
}

// Exec format exec
func (f *Format) Exec(ctx *context.Context, wand *gm.MagickWand) error {
	log.TraceBegin(ctx, "", "URL.Command", "format", "format", f.format)
	defer log.TraceEnd(ctx, nil)
	return wand.SetFormat(f.format)
}
