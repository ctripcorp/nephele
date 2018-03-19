package app

import (
	"github.com/ctripcorp/nephele/codec"
	"github.com/ctripcorp/nephele/log"
	"github.com/ctripcorp/nephele/service"
	"github.com/ctripcorp/nephele/store"
)

// Config represents configuration for all components
type Config interface {
	// Return current environment
	Env() string

	// Returns store config.
	Store() store.Config

	// Return codec config.
	Codec() codec.Config

	// Return service config.
	Service() service.Config

	// Return logger config.
	Logger() log.Config

	// Implements how to parse config.
	LoadFrom(env, path string) error

	// Reload
	Reload() error
}
