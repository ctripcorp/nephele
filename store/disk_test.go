package store

import (
	"os"
	"testing"

	"github.com/ctripcorp/nephele/context"
)

func Test_DiskRead(t *testing.T) {
	dir := os.Getenv("GOROOT")
	disk := &Disk{Path: dir}
	_, err := disk.Read(context.Context{}, "README.md")
	if err != nil {
		t.Error(err)
		return
	}
}
