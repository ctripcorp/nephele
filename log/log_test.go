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
	TraceBegin(&context.Context{}, "yes, lets start", "Info", "DelayCollect")

	Debugf(&context.Context{}, "%s!!!", "lets start")
	Errorf(&context.Context{}, "%s...", "so just wait a minute")
	Infow(&context.Context{}, "this is a info list",
		"name", "mag",
		"gender", "male")

	Errorw(&context.Context{}, "this is a error list",
		"name", "invalid last name",
		"gender", "not male nor female")

	TraceEndRoot(&context.Context{}, errors.New("seems something wrong"))
}
