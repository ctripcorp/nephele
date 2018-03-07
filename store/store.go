package store

import "github.com/nephele/context"

// Storage represents where to get or write image.
type Store interface {
	Read(ctx context.Context, path string) ([]byte, error)
	Write(ctx context.Context, blob []byte, path string) error
	Delete(ctx context.Context, path string) error
}

var storage Store

func Init(conf Config) error {
	var err error
	storage, err = conf.BuildStore()
	return err
}

func Read(ctx context.Context, path string) ([]byte, error) {
	return storage.Read(ctx, path)
}

func Write(ctx context.Context, blob []byte, path string) error {
	return storage.Write(ctx, blob, path)
}

func Delete(ctx context.Context, path string) error {
	return storage.Delete(ctx, path)
}

func NewDisk(conf DiskConfig) *Disk {
	return nil
}
