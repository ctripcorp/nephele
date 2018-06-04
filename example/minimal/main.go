package main

import (
	"github.com/ctripcorp/nephele/interpret"
	"github.com/ctripcorp/nephele/process"
	"github.com/ctripcorp/nephele/storage"

	"github.com/ctripcorp/nephele/server"
	"github.com/ctripcorp/nephele/util"

	_ "github.com/ctripcorp/nephele/command/autoorient"
	_ "github.com/ctripcorp/nephele/command/crop"
	_ "github.com/ctripcorp/nephele/command/format"
	_ "github.com/ctripcorp/nephele/command/quality"
	_ "github.com/ctripcorp/nephele/command/resize"
	_ "github.com/ctripcorp/nephele/command/rotate"
	_ "github.com/ctripcorp/nephele/command/sharpen"
	_ "github.com/ctripcorp/nephele/command/strip"
	_ "github.com/ctripcorp/nephele/command/watermark"

	_ "github.com/ctripcorp/nephele/interpret/neph"
	_ "github.com/ctripcorp/nephele/storage/neph"
	_ "github.com/ctripcorp/nephele/server/ping"
)

var Config = struct {
	ServerConfigPath string `toml:"server_config_path"`
	Interpret        *map[string]string
	Process          *map[string]string
	Storage          *map[string]string
}{
	Interpret: &interpret.Config,
	Process:   &process.Config,
	Storage:   &storage.Config,
}

func main() {
	util.FromToml("default.toml", &Config)
	util.FromToml(Config.ServerConfigPath, &server.Config)

	interpret.Init()
	process.Init()
	storage.Init()

	server.Run()
}
