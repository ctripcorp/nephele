package throttle

import (
    "errors"
    "time"
    "fmt"
    "sync"
)

var limiters sync.Map

type limiter struct {
	concurrency *pool
	wait        *pool
}

func (l *limiter) Do(timeout time.Duration) error {
	if l.wait.get(0) != nil {
		return errors.New("fast fail")
	}
	defer l.wait.put()
	return l.concurrency.get(timeout)
}

func (l *limiter) Done() error {
	return l.concurrency.put()
}

func Limiter(name string, maxConcurrency int, maxWait int) *limiter {
	id := fmt.Sprintf("%s-%d-%d", name, maxConcurrency, maxWait)
	l, ok := limiters.Load(id)
	if !ok {
		c := &pool{}
		c.init(maxConcurrency)
		w := &pool{}
		w.init(maxWait)

        limiters.Store(id, &limiter{
			concurrency: c,
			wait:        w,
		})

		l, ok = limiters.Load(id)
    }
    if !ok {
        return nil
    }
    return l.(*limiter)
}
