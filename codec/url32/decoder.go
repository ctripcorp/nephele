package url32

import (
	"strings"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/index"
	simpleIndex "github.com/ctripcorp/nephele/index/simple"
	"github.com/ctripcorp/nephele/process"
	"github.com/ctripcorp/nephele/transform"
	"github.com/ctripcorp/nephele/transform/simple"
)

//Decoder represents decoder
type Decoder struct {
	ctx       *context.Context
	processes []process.Process
	path      string
}

//Decode uri
func (e *Decoder) Decode(uri string) error {
	e.path = strings.Split(uri, "?")[0]
	procs, err := process.BuildProcesses(uri)
	if err != nil {
		return err
	}
	e.processes = procs
	return nil
}

//CreateIndex create index
func (e *Decoder) CreateIndex() index.Index {
	return &simpleIndex.Index{Ctx: e.ctx, Path: e.path}
}

//CreateTransformer create transformer
func (e *Decoder) CreateTransformer() transform.Transformer {
	return &simple.Transformer{Processes: e.processes}
}
