package format

import (
    "github.com/ctripcorp/nephele/command"
)

func init() {
    command.SetCommand("format", func() command.Command {
        return &Command{}
    })
}
