package process

import (
	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
)

type StripCommand struct {
	Wand *gm.MagickWand
}

// strip
func (s *StripCommand) Exec(ctx *context.Context) error {
	//log here
	return s.Wand.Strip()
}
