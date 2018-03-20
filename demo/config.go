package main

import (
	"fmt"
	"github.com/ctripcorp/nephele/codec"
	"github.com/ctripcorp/nephele/codec/url32"
	"github.com/ctripcorp/nephele/log"
	"github.com/ctripcorp/nephele/service"
	"github.com/ctripcorp/nephele/store"
	"github.com/ctripcorp/nephele/util"
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

	if conf.codec, err = url32.DefaultConfig(); err != nil {
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

// Reload configuration
func (conf *DemoConfig) Reload() error {
	fmt.Println("reload")
	return nil
}
