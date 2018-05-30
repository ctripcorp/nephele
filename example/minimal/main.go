package main

import (
	"github.com/ctripcorp/nephele/interpret"
	"github.com/ctripcorp/nephele/server"
	"github.com/ctripcorp/nephele/storage"
	"github.com/ctripcorp/nephele/util"

	_ "github.com/ctripcorp/nephele/command/format"
	_ "github.com/ctripcorp/nephele/interpret/neph"
	_ "github.com/ctripcorp/nephele/storage/neph"
)

var Config = struct {
	ServerConfigPath string `toml:"server_config_path"`
	Interpret        *map[string]string
	Storage          *map[string]string
	Process          *map[string]string
}{
	Interpret: &interpret.Config,
	Storage:   &storage.Config,
}

func main() {
	util.FromToml("default.toml", &Config)
	util.FromToml(Config.ServerConfigPath, &server.Config)
	storage.Init()
	server.Run()
}
