package main

import (
	"github.com/ctripcorp/nephele/app"
)

// Demo app starts a powerful image server with minimal components
func main() {
	//run a simple server. enjoy!
	app.New(func(env string) app.Config {
		return new(DemoConfig)
	}).Run()
}
