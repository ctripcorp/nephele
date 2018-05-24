package main

import (
    "github.com/ctripcorp/nephele/server"
    "github.com/ctripcorp/nephele/util"
)

func main() {
    util.FromToml("server.toml", &server.Config)
    server.Run()
}
