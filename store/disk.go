package store

import (
	"github.com/nephele/context"
)

type Disk struct {
}

type DiskConfig struct {
	Dir string `toml:"dir"`
}

func (d *Disk) Read(ctx context.Context, path string) ([]byte, error) {
	return nil, nil
}

func (d *Disk) Write(ctx context.Context, blob []byte, path string) error {
	return nil
}

func (d *Disk) Delete(ctx context.Context, path string) error {
	return nil
}
