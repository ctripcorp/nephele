package resize

import (
	"github.com/ctripcorp/nephele/command"
)

func init() {
	command.Register("resize", func() command.Command {
		return &Command{}
	})
}
