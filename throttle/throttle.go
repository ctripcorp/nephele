package throttle

import (
	"fmt"
	"sync"
)

var workflows sync.Map

func Workflow(name string, maxConcurrency int, maxWait int) *workflow {
	id := fmt.Sprintf("%s-%d-%d", name, maxConcurrency, maxWait)
	l, ok := workflows.Load(id)
	if !ok {
		c := &pool{}
		c.init(maxConcurrency)
		w := &pool{}
		w.init(maxWait)

		workflows.Store(id, &workflow{
			concurrency: c,
			wait:        w,
		})

		l, ok = workflows.Load(id)
	}
	if !ok {
		return nil
	}
	return l.(*workflow)
}
