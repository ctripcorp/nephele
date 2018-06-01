package main

import (
	"github.com/ctripcorp/nephele/interpret"

	"github.com/ctripcorp/nephele/storage"
	_ "github.com/ctripcorp/nephele/storage/neph"

	"github.com/ctripcorp/nephele/server"
	"github.com/ctripcorp/nephele/util"

	"fmt"
)

var Config = struct {
	ServerConfigPath string `toml:"server_config_path"`
	Interpret        *map[string]string
	Storage          *map[string]string
}{
	Interpret: &interpret.Config,
	Storage:   &storage.Config,
}

func main() {

	util.FromToml("default.toml", &Config)
	util.FromToml(Config.ServerConfigPath, &server.Config)

	fmt.Println(server.Config)

	storage.Init()
	interpret.Init()
	server.Run()
}
