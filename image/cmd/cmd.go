package cmd

import "github.com/nephele/context"

type Cmd interface {
	Exec(ctx context.Context) error
}
