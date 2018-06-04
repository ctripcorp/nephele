package strip

import (
	"github.com/ctripcorp/nephele/command"
)

func init() {
	command.Register("strip", func() command.Command {
		return &Command{}
	})
}
