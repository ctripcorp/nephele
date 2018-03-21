package command

import (
	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
)

//Strip strip command
type Strip struct {
}

//Verfiy strip verfiy params
func (s *Strip) Verfiy(ctx *context.Context, params map[string]string) error {
	return nil
}

// Exec strip exec
func (s *Strip) Exec(ctx *context.Context, wand *gm.MagickWand) error {
	//log here
	return wand.Strip()
}
