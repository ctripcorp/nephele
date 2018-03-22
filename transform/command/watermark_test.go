package command

import (
	"io/ioutil"
	"testing"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
)

func TestExecWatermark(t *testing.T) {
	var w Watermark
	w.Name = "wm1"
	w.Location = "ne"

	bt, err := ioutil.ReadFile("200j.jpg")
	if err != nil {
		t.Error()
	}
	wand, err := gm.NewMagickWand(bt)
	if err != nil {
		t.Error()
	}
	var ctx context.Context
	err = w.Exec(&ctx, wand)
	if err != nil {
		t.Error()
	}
	bt1, err := wand.WriteBlob()
	if err != nil {
		t.Error()
	}
	ioutil.WriteFile("newWatermark.jpg", bt1, 0777)
}

// 200j.jpg is 320*180
func TestExecWatermark1(t *testing.T) {
	var w Watermark
	w.Name = "wm1"
	w.Location = ""

	w.X = 160
	w.Y = 90
	bt, err := ioutil.ReadFile("200j.jpg")
	if err != nil {
		t.Error()
	}
	wand, err := gm.NewMagickWand(bt)
	if err != nil {
		t.Error()
	}
	var ctx context.Context
	err = w.Exec(&ctx, wand)
	if err != nil {
		t.Error()
	}
	bt1, err := wand.WriteBlob()
	if err != nil {
		t.Error()
	}
	ioutil.WriteFile("newWatermark1.jpg", bt1, 0777)
}

// test x,y both are 0
func TestExecWatermark2(t *testing.T) {
	var w Watermark
	bt, err := ioutil.ReadFile("200j.jpg")
	if err != nil {
		t.Error()
	}
	wand, err := gm.NewMagickWand(bt)
	if err != nil {
		t.Error()
	}
	var ctx context.Context
	m3 := map[string]string{"n": "wm1", "d": "70", "l": "sw"}
	if w.Verify(&ctx, m3) != nil {
		t.Error()
	}
	if w.Exec(&ctx, wand) != nil {
		t.Error()
	}

	bt1, err := wand.WriteBlob()
	if err != nil {
		t.Error()
	}
	ioutil.WriteFile("newWatermark2.jpg", bt1, 0777)
}
