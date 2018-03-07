package app

import (
	"github.com/nephele/codec"
	"github.com/nephele/context"
	"github.com/nephele/logger"
	"github.com/nephele/service"
	"github.com/nephele/store"
	"github.com/urfave/cli"
	"os"
	"time"
)

// App represents nephele application.
type App struct {
	conf     Config
	server   *Server
	internal *cli.App
}

// Start new nephele application with default configuration.
func New() *App {
	return NewConfigured(new(DemoConfig))
}

// Start new nephele application with provided configuration.
func NewConfigured(conf Config) *App {
	return nil
}

// Return server to make initialization or configure service router.
func (app *App) Server() *Server {
	return app.server
}

// Run nephele application.
func (app *App) Run() {
	app.internal.Run(os.Args)
}

// Build a new server and make initialization with provided configuration.
func newServer(conf Config) (*Server, error) {
	var err error

	// init storage.
	if err = store.Init(conf.Store()); err != nil {
		return nil, err
	}

	// init codec to encode or decode image request URL.
	if err = codec.Init(conf.Codec()); err != nil {
		return nil, err
	}

	// create root context.
	ctx := context.New(time.Duration(conf.Service().RequestTimeout))

	return &Server{
		logger:  logger.New(conf.Logger(), nil),
		service: service.New(ctx, conf.Service()),
	}, err
}
