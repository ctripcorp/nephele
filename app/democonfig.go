package app

import (
	"github.com/nephele/codec"
	"github.com/nephele/logger"
	"github.com/nephele/service"
	"github.com/nephele/store"
)

// Define default configuration for demo app
type DemoConfig struct {
}

// Implementation to parse config.
func (conf *DemoConfig) From(path string) error {
	return nil
}

// Return demo service config.
func (conf *DemoConfig) Service() service.Config {
	return service.Config{}
}

// Return demo logger config.
func (conf *DemoConfig) Logger() logger.Config {
	return logger.Config{}
}

// Return demo storage config.
func (conf *DemoConfig) Store() store.Config {
	return nil
}

// Return demo codec config.
func (conf *DemoConfig) Codec() codec.Config {
	return nil
}
