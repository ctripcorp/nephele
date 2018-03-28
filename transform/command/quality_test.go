package command

import (
	"testing"
	"time"

	"github.com/ctripcorp/nephele/context"
)

func TestVerify(t *testing.T) {
	q := &Quality{}
	if e := q.Verify(context.New("", time.Duration(time.Second*3)), map[string]string{"v": "10"}); e != nil {
		t.Error(e)
	}
}
