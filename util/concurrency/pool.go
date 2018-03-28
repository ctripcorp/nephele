package concurrency

import (
	"errors"
	"time"
)

type poolItem struct{}

type pool struct {
	items chan *poolItem
}

func (p *pool) init(max int) {
	p.items = make(chan *poolItem, max)
	for i := 0; i < max; i++ {
		p.items <- &poolItem{}
	}
}

func (p *pool) get(timeout time.Duration) error {
	items := p.items
	if items == nil {
		return errors.New("pool is closed")
	}
	select {
	case <-items:
		return nil
	case <-time.After(timeout):
		return errors.New("timeout")
	}
}

func (p *pool) put() error {
	items := p.items
	if items == nil {
		return errors.New("pool is closed")
	}
	items <- &poolItem{}
	return nil
}
