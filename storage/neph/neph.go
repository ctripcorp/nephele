package neph

import (
	"github.com/ctripcorp/nephele/storage"

	"plugin"
)

var Config map[string]string

var instance storage.Storage

func Init() {
    if instance != nil {
        return
    }
    if Config["type"] == "inline" {
        panic("inline storage not supported")
    } 
    if Config["type"] == "plugin" {
        p, err := plugin.Open(Config["name"]+".so")
        if err != nil {
            panic(err)
        }
        s, err := p.Lookup("New")
        if err != nil {
        	panic(err)
        }
        instance = s.(func(config map[string]string) storage.Storage)(Config)
    }
}

func File(key string) storage.File {
    return instance.File(key)
}

func StorageFile(key string, blob []byte, options ...storage.KV) (string, error) {
    return instance.StoreFile(key, blob, options...)
}
