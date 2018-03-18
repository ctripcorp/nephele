package main

import (
	"github.com/nephele/codec"
	"github.com/nephele/codec/v1"
	"github.com/nephele/log"
	"github.com/nephele/service"
	"github.com/nephele/store"
	"github.com/nephele/util"
)

// Define configuration for demo app
type DemoConfig struct {
	env     string
	logger  log.Config        `toml:"log"`
	service service.Config    `toml:"service"`
	store   *store.DiskConfig `toml:"storage"`
	codec   codec.Config
}

// Return current environment
func (conf *DemoConfig) Env() string {
	return conf.env
}

// Return demo service config.
func (conf *DemoConfig) Service() service.Config {
	return conf.service
}

// Return demo logger config.
func (conf *DemoConfig) Logger() log.Config {
	return conf.logger
}

// Return demo storage config.
func (conf *DemoConfig) Store() store.Config {
	return conf.store
}

// Return demo codec config.
func (conf *DemoConfig) Codec() codec.Config {
	return conf.codec
}

// Implementation to parse config.
func (conf *DemoConfig) LoadFrom(env, path string) error {
	var err error

	// give default configuration
	if conf.service, err = service.DefaultConfig(); err != nil {
		return err
	}
	if conf.store, err = store.DefaultConfig(); err != nil {
		return err
	}
	if conf.logger, err = log.DefaultConfig(); err != nil {
		return err
	}
	if conf.codec, err = v1.DefaultConfig(); err != nil {
		return err
	}

	if len(path) != 0 {
		if err = util.FromToml(path, conf); err != nil {
			return err
		}
	}
	conf.env = env

	return err
}
