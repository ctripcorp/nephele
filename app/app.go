package app

import (
	"fmt"
	_ "fmt"
	"github.com/nephele/codec"
	"github.com/nephele/context"
	"github.com/nephele/log"
	"github.com/nephele/service"
	"github.com/nephele/store"
	"github.com/urfave/cli"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// App represents nephele application.
type App struct {
	conf     Config
	server   *Server
	internal *cli.App
}

// Define configurator by runtime environment
type Configurator func(string) Config

// Return new nephele application with given configuration.
func New(configure Configurator) *App {
	app := new(App)
	app.internal = cli.NewApp()
	app.internal.Name = "Nephele"
	app.internal.Usage = "Powerful image service!"
	app.internal.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "load configuration from given path.",
		},
		cli.StringFlag{
			Name:   "env, e",
			Value:  "dev",
			Usage:  "set app environment.",
			EnvVar: "NEPHELE_ENV",
		},
		cli.StringFlag{
			Name:  "signal, s",
			Usage: "send signal to: open, stop, quit, reopen",
		},
	}

	app.internal.Action = func(ctx *cli.Context) error {
		switch strings.ToLower(ctx.String("signal")) {
		case "open":
			return app.open(ctx, configure)
		case "reopen":
			return app.reopen()
		case "quit":
			return app.quit()
		case "stop":
			return app.stop()
		default:
			return app.open(ctx, configure)
		}
	}
	return app
}

// Return server to make initialization or configure service router.
func (app *App) Server() *Server {
	return app.server
}

// Run nephele application.
func (app *App) Run() {
	if err := app.internal.Run(os.Args); err != nil {
		//log.Fatal(err.Error())
	}
}

// Open server
func (app *App) open(ctx *cli.Context, configure Configurator) error {
	var err error

	env := ctx.String("env")
	path := ctx.String("config")

	conf := configure(env)

	if err = conf.LoadFrom(env, path); err != nil {
		return err
	}

	if err = app.initComponents(conf); err != nil {
		return err
	}

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM)
	select {
	case sig := <-c:
		fmt.Println(sig)
	case err = <-app.buildServer(conf).Open():
	}

	return err
}

// Reopen server
func (app *App) reopen() error {
	return nil
}

// Quit server gracefully
func (app *App) quit() error {
	return nil
}

// Stop server immediately
func (app *App) stop() error {
	return nil
}

func (app *App) initComponents(conf Config) error {
	var err error

	// init storage.
	if err = store.Init(conf.Store()); err != nil {
		return err
	}

	// init codec to encode or decode image request URL.
	if err = codec.Init(conf.Codec()); err != nil {
		return err
	}

	// init logger
	if err = log.Init(conf.Logger()); err != nil {
		return err
	}

	return err
}

// Build a new server and make initialization with given configuration.
func (app *App) buildServer(conf Config) *Server {
	// create root context.
	ctx := context.New(conf.Env(), time.Duration(conf.Service().RequestTimeout))

	return &Server{
		service: service.New(ctx, conf.Service()),
	}
}
