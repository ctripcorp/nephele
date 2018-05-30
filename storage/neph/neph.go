package neph

import (
	stor "github.com/ctripcorp/nephele/storage"
)

func init() {
	stor.Register(New)
}

func New(config map[string]string) stor.Storage {
	return &storage{
		root: config["root"],
	}
}
