package command

import (
	"io/ioutil"
	"testing"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
)

func TestRotate(t *testing.T) {
	var r Rotate
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
	m3 := map[string]string{"v": "145"}
	if r.Verify(&ctx, m3) != nil {
		t.Error()
	}
	if r.Exec(&ctx, wand) != nil {
		t.Error()
	}

	bt1, err := wand.WriteBlob()
	if err != nil {
		t.Error()
	}
	ioutil.WriteFile("newRotate.jpg", bt1, 0777)
}
