package app

import (
	"fmt"
	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/service"
	"github.com/urfave/cli"
	"log"
	"os"
	"os/signal"
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
		cli.BoolFlag{
			Name:  "open, o",
			Usage: "open server to serve image",
		},
	}

	app.internal.Action = func(ctx *cli.Context) error {
		if ctx.Bool("open") {
			if err := app.open(ctx, configure); err != nil {
				return cli.NewExitError(err.Error(), 1)
			}
		}
		return nil
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
		log.Fatal(err.Error())
	}
}

// Open server
func (app *App) open(ctx *cli.Context, configure Configurator) error {
	var err error

	env := ctx.String("env")
	path := ctx.String("config")

	app.conf = configure(env)

	if err = app.conf.LoadFrom(env, path); err != nil {
		return err
	}

	s := app.buildServer(app.conf)
	if err = s.init(); err != nil {
		return err
	}

	select {
	case err = <-app.acceptSignals():
	case err = <-s.open():
	}

	return err
}

func (app *App) reload() error {
	return app.conf.Reload()
}

// Close server gracefully.
func (app *App) quit() error {
	return app.server.quit()
}

// Build a new server.
func (app *App) buildServer(conf Config) *Server {
	// create root context.
	ctx := context.New(conf.Env(), time.Duration(conf.Service().RequestTimeout)*time.Millisecond)
	app.server = &Server{
		conf:    conf,
		service: service.New(ctx, conf.Service()),
	}
	return app.server
}

// register signal to c.
// when syscall.SIGINT is received, error log should be written.
// when syscall.SIGHUP is received, app will be reload.
// when syscall.SIGKILL is received, app will forcibly closed.
// when syscall.SIGTERM is received, server will quit gracefully.
func (app *App) acceptSignals() <-chan error {
	c := make(chan error)
	go func() {
		sc := make(chan os.Signal)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM)
		sig := <-sc
		switch sig {
		case syscall.SIGHUP:
			if err := app.reload(); err != nil {
				c <- fmt.Errorf("app closed unexpectedly for reload failed. error:%s", err.Error())
			}
		case syscall.SIGTERM:
			{
				if err := app.quit(); err != nil {
					c <- fmt.Errorf("app recevied signal:%s, but quit failed. error:%s", sig.String(), err.Error())
				} else {
					c <- fmt.Errorf("app closed gracefully for receiving signal:%s", sig.String())
				}
			}
		case syscall.SIGINT:
			c <- fmt.Errorf("app closed unexpectedly for receiving signal:%s", sig.String())
		}
	}()
	return c
}
