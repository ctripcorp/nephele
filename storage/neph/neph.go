package neph

import (
	. "github.com/ctripcorp/nephele/storage"

	"plugin"
)

var New func(config map[string]string) Storage

func init() {
	oss, err := plugin.Open("oss.so")
	if err != nil {
		panic(err)
	}
	symbol, err := oss.Lookup("New")
	if err != nil {
		panic(err)
	}
	New = symbol.(func(config map[string]string) Storage)
}

func Configurate(config map[string]string) {
	return
}
