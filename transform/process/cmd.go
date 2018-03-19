package process

import "github.com/nephele/context"

type Cmd interface {
	Exec(ctx context.Context) error
}
