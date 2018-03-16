package app

import (
	"fmt"
	"github.com/nephele/context"
	"github.com/nephele/service"
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

	s := app.buildServer(conf)
	if err = s.init(); err != nil {
		return err
	}

	// register signal to c.
	// when syscall.SIGINT is received, error log should be written.
	// when syscall.SIGHUP is received, app should be reopened.
	// when syscall.SIGTERM is received, server should quit gracefully.
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM)
	select {
	case sig := <-c:
		switch sig {
		case syscall.SIGHUP:
		case syscall.SIGTERM:
		case syscall.SIGINT:
			err = fmt.Errorf("app closed unexpectedly for receving signal:%s", sig.String())
		}
	case err = <-s.Open():
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

// Build a new server.
func (app *App) buildServer(conf Config) *Server {
	// create root context.
	ctx := context.New(conf.Env(), time.Duration(conf.Service().RequestTimeout))

	return &Server{
		conf:    conf,
		service: service.New(ctx, conf.Service()),
	}
}
