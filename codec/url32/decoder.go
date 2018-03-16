package url32

import (
	"strings"

	"github.com/nephele/context"
	"github.com/nephele/index"
	"github.com/nephele/process"
	"github.com/nephele/transform"
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
	return index.SimpleIndex{}
}

//CreateTransformer create transformer
func (e *Decoder) CreateTransformer() transform.Transformer {
	return simple.NewTransformer(e.processes)
}
