package store

import (
	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/util"
	"path/filepath"
)

// Storage represents where to get or write image.
type Store interface {
	Read(ctx context.Context, path string) ([]byte, error)
	Delete(ctx context.Context, path string) error
	Write(ctx context.Context, path string, blob []byte) error
	WriteOffset(ctx context.Context, path string, blob []byte, offset int64) error
}

// Config represents how to build storage and storage configuration.
type Config interface {
	BuildStore() (Store, error)
}

var storage Store

func Init(conf Config) error {
	var err error
	storage, err = conf.BuildStore()
	return err
}

func DefaultConfig() (*DiskConfig, error) {
	var err error
	var homeDir string

	if homeDir, err = util.HomeDir(); err != nil {
		return nil, err
	}

	return &DiskConfig{
		Dir: filepath.Join(homeDir, "image"),
	}, nil
}

func Storage() Store {
	return storage
}
