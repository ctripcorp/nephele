package log

import (
	"errors"
	"testing"

	"github.com/ctripcorp/nephele/context"
)

func TestLogger(t *testing.T) {
	dc, _ := DefaultConfig()
	Init(dc)

	TraceBegin(&context.Context{}, "we are going to gather some personal infomations", "Info", "Collect")

	Debugf(&context.Context{}, "%s!!!", "lets start")
	Infow(&context.Context{}, "this is a info list",
		"name", "mag",
		"gender", "male")

	TraceEnd(&context.Context{}, errors.New("seems something wrong"))
}
