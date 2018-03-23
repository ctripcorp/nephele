package store

import (
	"io/ioutil"
	"strings"

	"github.com/ctripcorp/nephele/context"
)

type Disk struct {
	Path string
}

func (d *Disk) Read(ctx context.Context, path string) ([]byte, error) {
	if strings.HasPrefix(path, "\\") {
		path = strings.Replace(path, "/", "\\", -1)
	}
	if strings.HasPrefix(path, "/") {
		path = strings.Replace(path, "\\", "/", -1)
	}
	if strings.HasSuffix(d.Path, "/") {
		path = d.Path + path
	} else {
		path = d.Path + "/" + path
	}

	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return buff, nil
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
