package main

import (
	"fmt"

	"github.com/ctripcorp/nephele/codec"
	"github.com/ctripcorp/nephele/codec/neph"
	"github.com/ctripcorp/nephele/log"
	"github.com/ctripcorp/nephele/service"
	"github.com/ctripcorp/nephele/service/middleware"
	"github.com/ctripcorp/nephele/store"
	"github.com/ctripcorp/nephele/util"
)

// Define configuration for demo app
type config struct {
	env           string
	codec         codec.Config
	LogConfig     *log.LoggerConfig `toml:"log"`
	ServiceConfig service.Config    `toml:"service"`
	StoreConfig   *store.DiskConfig `toml:"storage"`
}

// Return current environment
func (conf *config) Env() string {
	return conf.env
}

// Return demo logger config.
func (conf *config) Log() log.Config {
	return conf.LogConfig
}

// Return demo server config.
func (conf *config) Service() service.Config {
	return conf.ServiceConfig
}

// Return demo storage config.
func (conf *config) Store() store.Config {
	return conf.StoreConfig
}

// Return demo codec config.
func (conf *config) Codec() codec.Config {
	return conf.codec
}

// Implementation to parse config.
func (conf *config) LoadFrom(env, path string) error {
	var err error

	// give default configuration
	if conf.ServiceConfig, err = service.DefaultConfig(); err != nil {
		return err
	}

	if conf.StoreConfig, err = store.DefaultConfig(); err != nil {
		return err
	}

	if conf.LogConfig, err = log.DefaultConfig(); err != nil {
		return err
	}

	if conf.codec, err = neph.DefaultConfig(); err != nil {
		return err
	}

	conf.env = env
	if len(path) != 0 {
		if err = util.FromToml(path, conf); err != nil {
			return err
		}
	}
	if len(conf.Service().MiddlewarePath) != 0 {
		middlewareConf := &middleware.MiddlewareConfig{}
		if err = util.FromToml(conf.ServiceConfig.MiddlewarePath, middlewareConf); err != nil {
			return err
		}
		conf.ServiceConfig.Middleware = middlewareConf
	}
	if len(conf.LogConfig.ConfigPath) != 0 {
		loggerConf := &log.LoggerConfig{}
		if err = util.FromToml(conf.LogConfig.ConfigPath, loggerConf); err != nil {
			return err
		}
		conf.LogConfig = loggerConf
	}
	return err
}

// Reload configuration
func (conf *config) Reload() error {
	fmt.Println("reload")
	return nil
}
