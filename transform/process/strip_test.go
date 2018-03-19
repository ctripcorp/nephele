package process

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
)

func TestExecStrip(t *testing.T) {
	var r StripCommand
	bt, err := ioutil.ReadFile("200j.jpg")
	if err != nil {
		t.Error()
	}
	r.Wand, err = gm.NewMagickWand(bt)
	if err != nil {
		t.Error()
	}

	var ctx context.Context

	r.Exec(ctx)

	bt1, err := r.Wand.WriteBlob()
	if err != nil {
		fmt.Println(err.Error())
		t.Error()
	}
	ioutil.WriteFile("newStrip.jpg", bt1, 0777)
}
