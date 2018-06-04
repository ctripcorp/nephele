package quality

import (
	"github.com/ctripcorp/nephele/command"
)

func init() {
	command.Register("quality", func() command.Command {
		return &Command{}
	})
}
