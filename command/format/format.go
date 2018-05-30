package format

import (
	"github.com/ctripcorp/nephele/command"
)

func init() {
	command.Register("format", func() command.Command {
		return &Command{}
	})
}
