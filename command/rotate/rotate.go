package rotate

import (
	"github.com/ctripcorp/nephele/command"
)

func init() {
	command.Register("rotate", func() command.Command {
		return &Command{}
	})
}
