package app

import (
	"github.com/nephele/codec"
	"github.com/nephele/logger"
	"github.com/nephele/service"
	"github.com/nephele/store"
)

type Config interface {
	From(path string) error
	Store() store.Config
	Codec() codec.Config
	Service() service.Config
	Logger() logger.Config
}

// define default configuration for demo app
type DemoConfig struct {
}

//
func (conf *DemoConfig) From(path string) error {
	return nil
}

func (conf *DemoConfig) Service() service.Config {
	return service.Config{}
}

func (conf *DemoConfig) Logger() logger.Config {
	return logger.Config{}
}

func (conf *DemoConfig) Store() store.Config {
	return nil
}

func (conf *DemoConfig) Codec() codec.Config {
	return nil
}
