package command

import (
	"io/ioutil"
	"testing"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
)

func TestExecFormat(t *testing.T) {
	var f Format
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
	m3 := map[string]string{"v": "webp"}
	if f.Verfiy(&ctx, m3) != nil {
		t.Error()
	}
	if f.Exec(&ctx, wand) != nil {
		t.Error()
	}
}

func TestExecFormat1(t *testing.T) {
	var f Format
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
	m3 := map[string]string{"v": "png"}
	if f.Verfiy(&ctx, m3) != nil {
		t.Error()
	}
	if f.Exec(&ctx, wand) != nil {
		t.Error()
	}
}
