package concurrency

import (
	"time"
)

var limiters map[string]*limiter

type limiter struct {
	p *pool
}

func (l *limiter) Do() error {
	_, err := l.p.get()
	return err
}

func (l *limiter) Done() error {
	return l.p.put(&poolItem{})
}

func Limiter(name string, maxConcurrency int, maxWait int, timeout time.Time) *limiter {
	return nil
}
