package neph

import (
	stor "github.com/ctripcorp/nephele/storage"
	"io/ioutil"
	"path/filepath"
)

type storage struct {
	root string
}

func (s *storage) File(key string) stor.File {
	return &file{s.root, key}
}

func (s *storage) Iterator(prefix string, lastKey string) stor.Iterator {
	return nil
}

func (s *storage) StoreFile(key string, blob []byte, kvs ...stor.KV) (string, error) {
	return "", ioutil.WriteFile(filepath.Join(s.root, key), []byte(blob), 0666)
}
