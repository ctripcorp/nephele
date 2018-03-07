package app

import (
	"github.com/nephele/codec"
	"github.com/nephele/logger"
	"github.com/nephele/service"
	"github.com/nephele/store"
)

// Config represents configuration for all components
type Config interface {
	// Implements how to parse config.
	From(path string) error

	// Return store config.
	Store() store.Config

	// Return codec config.
	Codec() codec.Config

	// Return service config.
	Service() service.Config

	// Return logger config.
	Logger() logger.Config
}
