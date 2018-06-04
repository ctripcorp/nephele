package sharpen

import (
	"github.com/ctripcorp/nephele/command"
)

func init() {
	command.Register("sharpen", func() command.Command {
		return &Command{}
	})
}
