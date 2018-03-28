package concurrency

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var limiters map[string]*limiter
var limiterMutex sync.RWMutex

func init() {
	limiters = make(map[string]*limiter)
}

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
	limiterMutex.RLock()
	l, ok := limiters[id]
	limiterMutex.RUnlock()
	if !ok {
		limiterMutex.Lock()
		c := &pool{}
		c.init(maxConcurrency)
		w := &pool{}
		w.init(maxWait)
		limiters[id] = &limiter{
			concurrency: c,
			wait:        w,
		}
		limiterMutex.Unlock()

		limiterMutex.RLock()
		l = limiters[id]
		limiterMutex.RUnlock()
	}
	return l
}
