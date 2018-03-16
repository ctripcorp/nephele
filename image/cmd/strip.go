package cmd

import (
	"github.com/nephele/context"
	"github.com/nephele/img4go/gm"
)

type StripCommand struct {
	Wand *gm.MagickWand
}

func (s *StripCommand) Exec(ctx context.Context) error {
	//log here
	println("test strip")
	return s.Wand.Strip()
}
