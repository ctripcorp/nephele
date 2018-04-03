package throttle

import (
	"errors"
	"time"
)

type workflow struct {
	concurrency *pool
	wait        *pool
}

func (l *workflow) Do(timeout time.Duration) error {
	if l.wait.get(0) != nil {
		return errors.New("fast fail")
	}
	defer l.wait.put()
	return l.concurrency.get(timeout)
}

func (l *workflow) Done() error {
	return l.concurrency.put()
}
