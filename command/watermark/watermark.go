package watermark

import (
	"github.com/ctripcorp/nephele/command"
)

func init() {
	command.Register("watermark", func() command.Command {
		return &Command{}
	})
}
