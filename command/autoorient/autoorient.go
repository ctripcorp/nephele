package autoorient

import (
	"github.com/ctripcorp/nephele/command"
)

func init() {
	command.Register("autoorient", func() command.Command {
		return &Command{}
	})
}
