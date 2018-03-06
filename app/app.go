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

type App struct {
	conf     Config
	server   *Server
	internal *cli.App
}

func New() *App {
	return NewConfigured(new(DemoConfig))
}

func NewConfigured(conf Config) *App {
	return nil
}

func (app *App) Server() *Server {
	return app.server
}

func (app *App) Run() {
	app.internal.Run(os.Args)
}

func newServer(conf Config) (*Server, error) {
	var err error

	if err = store.Init(conf.Store()); err != nil {
		return nil, err
	}

	if err = codec.Init(conf.Codec()); err != nil {
		return nil, err
	}

	ctx := context.New(time.Duration(conf.Service().RequestTimeout))

	return &Server{
		logger:  logger.New(conf.Logger(), nil),
		service: service.New(ctx, conf.Service()),
	}, err
}
