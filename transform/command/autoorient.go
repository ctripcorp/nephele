package command

import (
	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
	"github.com/ctripcorp/nephele/log"
)

type AutoOrient struct {
	AutoOrient string
}

const (
	autoOrientKeyA string = "autoOrient"
)

//Verify autoorient verify
func (a *AutoOrient) Verify(ctx *context.Context, params map[string]string) error {
	log.Debugf(ctx, "autoOrient verification")
	return nil
}

//Exec autoorient exec
func (a *AutoOrient) Exec(ctx *context.Context, wand *gm.MagickWand) error {
	log.TraceBegin(ctx, "", "URL.Command", "autoOrient", "")
	defer log.TraceEnd(ctx, nil)
	return wand.AutoOrient()
}
