package main

import (
    storage "github.com/ctripcorp/nephele/storage/neph"
    "github.com/ctripcorp/nephele/util"
)

var Config = struct {
    Storage *map[string]string
} {
    Storage: &storage.Config,
}

func main() {
    util.FromToml("default.toml", &Config)
    storage.Init()
}
