package neph

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	stor "github.com/ctripcorp/nephele/storage"
)

type file struct {
	root string
	key  string
}

func (f *file) Key() string {
	return f.key
}

func (f *file) Exist() (bool, string, error) {
	_, err := os.Stat(filepath.Join(f.root, f.key))
	if err == nil {
		return true, "", nil
	}
	if os.IsNotExist(err) {
		return false, "", nil
	}
	return false, "", err
}

func (f *file) Meta() (stor.Fetcher, error) {
	return nil, errors.New("Meta not supported")
}

func (f *file) Append(blob []byte, index int64, kvs ...stor.KV) (int64, string, error) {
	return index, "", errors.New("Append not supported")
}

func (f *file) Delete() (string, error) {
	return "", os.Remove(f.Key())
}

func (f *file) Bytes() ([]byte, string, error) {
	blob, err := ioutil.ReadFile(filepath.Join(f.root, f.key))
	fmt.Println("file.go-Bytes", err)
	return blob, "", err
}

func (f *file) SetMeta(...stor.KV) error {
	return errors.New("SetMeta not supported")
}
