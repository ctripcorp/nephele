package command

import (
	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
)

//Strip strip command
type Strip struct {
}

//Verify strip Verify params
func (s *Strip) Verify(ctx *context.Context, params map[string]string) error {
	//log.Debugw(ctx, "begin strip verfiy")
	return nil
}

// Exec strip exec
func (s *Strip) Exec(ctx *context.Context, wand *gm.MagickWand) error {
	//log.TraceBegin(ctx, "strip exec", "URL.Command", "strip")
	//defer log.TraceEnd(ctx, nil)
	return wand.Strip()
}
