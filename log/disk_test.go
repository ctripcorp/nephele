package log

import (
	"github.com/ctripcorp/nephele/context"
	"testing"
)

func TestDiskLogger(t *testing.T) {
	dl := &diskLogger{"test.log"}
	dl.Printf(context.Context{}, "debug", "%s!!!", "hello")
	dl.Printw(context.Context{}, "debug", "list", "name", "mag", "gender", "male")
}
