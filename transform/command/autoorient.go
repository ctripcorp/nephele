package command

import (
	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
	"github.com/ctripcorp/nephele/log"
)

type Autoorient struct {
	Autoorient string
}

const (
	autoorientKeyA string = "autoorient"
)

//Verify autoorient verify
func (a *Autoorient) Verify(ctx *context.Context, params map[string]string) error {
	log.Debugf(ctx, "autoorient verification")
	return nil
}

//Exec autoorient exec
func (a *Autoorient) Exec(ctx *context.Context, wand *gm.MagickWand) error {
	log.TraceBegin(ctx, "", "URL.Command", "autoorient", "")
	defer log.TraceEnd(ctx, nil)
	return wand.AutoOrient()
}
