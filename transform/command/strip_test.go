package command

import (
	"io/ioutil"
	"testing"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
)

func TestStrip(t *testing.T) {
	var r Strip
	bt, err := ioutil.ReadFile("200j.jpg")
	if err != nil {
		t.Error(err)
		return
	}
	wand, err := gm.NewMagickWand(bt)
	if err != nil {
		t.Error(err)
		return
	}

	var ctx context.Context
	r.Exec(&ctx, wand)

	blob, err := wand.WriteBlob()
	if err != nil {
		t.Error(err)
		return
	}
	ioutil.WriteFile("newStrip.jpg", blob, 0777)
}
