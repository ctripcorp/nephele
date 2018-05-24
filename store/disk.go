package store

import (
	"io/ioutil"
	"strings"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/log"
	"github.com/ctripcorp/nephele/util"
)

//Disk disk store
type Disk struct {
	Path string
}

func (d *Disk) Read(ctx *context.Context, path string) ([]byte, error) {
	var err error
	log.TraceBegin(ctx, "", "store", "disk.read", "path", path)
	defer log.TraceEnd(ctx, err)
	path = util.TrimPrefixSlash(path)

	if strings.HasSuffix(d.Path, "/") {
		path = util.JoinString(d.Path, path)
	} else {
		path = util.JoinString(d.Path, "/", path)
	}
	var buff []byte
	buff, err = ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return buff, nil
}

//Delete delete file
func (d *Disk) Delete(ctx *context.Context, path string) error {
	return nil
}

//Write write file
func (d *Disk) Write(ctx *context.Context, path string, blob []byte) error {
	return nil
}

//WriteOffset append file
func (d *Disk) WriteOffset(ctx *context.Context, path string, blob []byte, offset int64) error {
	return nil
}
