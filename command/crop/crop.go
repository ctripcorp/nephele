package crop

import (
	"github.com/ctripcorp/nephele/command"
)

func init() {
	command.Register("crop", func() command.Command {
		return &Command{}
	})
}
