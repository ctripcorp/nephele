package store

import (
	"github.com/nephele/context"
)

type Disk struct {
}

func (d *Disk) Read(ctx context.Context, path string) ([]byte, error) {
	return nil, nil
}

func (d *Disk) Delete(ctx context.Context, path string) error {
	return nil
}

func (d *Disk) Write(ctx context.Context, path string, blob []byte) error {
	return nil
}

func (d *Disk) WriteOffset(ctx context.Context, path string, blob []byte, offset int64) error {
	return nil
}
