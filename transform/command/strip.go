package command

import (
	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
	"github.com/ctripcorp/nephele/log"
)

//Strip strip command
type Strip struct {
}

//Verify strip Verify params
func (s *Strip) Verify(ctx *context.Context, params map[string]string) error {
	log.Debugf(ctx, "strip verification")
	return nil
}

// Exec strip exec
func (s *Strip) Exec(ctx *context.Context, wand *gm.MagickWand) error {
	var err error
	log.TraceBegin(ctx, "", "URL.Command", "strip")
	defer log.TraceEnd(ctx, err)
	err = wand.Strip()
	return err
}
