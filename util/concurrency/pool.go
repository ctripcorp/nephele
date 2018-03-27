package concurrency

import (
	"errors"
	"time"
)

type poolItem struct{}

type pool struct {
	timeout time.Duration

	items chan *poolItem
}

func (p *pool) init(timeout time.Duration, max int) {
	p.timeout = timeout
	p.items = make(chan *poolItem, max)
	for i := 0; i < max; i++ {
		p.items <- &poolItem{}
	}
}

func (p *pool) get() (*poolItem, error) {
	items := p.items
	if items == nil {
		return nil, errors.New("pool is closed")
	}
	select {
	case item := <-items:
		return item, nil
	case <-time.After(p.timeout):
		return nil, errors.New("timeout")
	}
}

func (p *pool) put(item *poolItem) error {
	items := p.items
	if items == nil {
		return errors.New("pool is closed")
	}
	items <- item
	return nil
}
