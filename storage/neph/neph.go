package neph

import (
	stor "github.com/ctripcorp/nephele/storage"

	"plugin"
)

var Config map[string]string

var instance stor.Storage

func Init() {
	if instance != nil {
		return
	}
	if Config["type"] == "inline" {
		instance = &storage{
			root: Config["root"],
		}
	}
	if Config["type"] == "plugin" {
        plugin.Open(Config["path"])
		p, err := plugin.Open(Config["path"])
		if err != nil {
			panic(err)
		}
		s, err := p.Lookup("New")
		if err != nil {
			panic(err)
		}
		instance = s.(func(config map[string]string) stor.Storage)(Config)
	}
}

func File(key string) stor.File {
	return instance.File(key)
}

func StoreFile(key string, blob []byte, options ...stor.KV) (string, error) {
	return instance.StoreFile(key, blob, options...)
}
