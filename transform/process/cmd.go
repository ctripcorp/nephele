package process

import "github.com/ctripcorp/nephele/context"

type Cmd interface {
	Exec(ctx *context.Context) error
}
