package command

import (
	"encoding/base64"
	"io/ioutil"
	"testing"

	"github.com/ctripcorp/nephele/store"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
)

func TestExecWatermark(t *testing.T) {
	conf, _ := store.DefaultConfig()

	store.Init(conf)

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
	m3 := map[string]string{"n": base64.StdEncoding.EncodeToString([]byte("wm1.png")), "d": "50", "l": "south"}
	if e := w.Verify(&ctx, m3); e != nil {
		t.Error()
	}
	if e := w.Exec(&ctx, wand); e != nil {
		t.Error(e)
	}

	bt1, err := wand.WriteBlob()
	if err != nil {
		t.Error()
	}
	ioutil.WriteFile("newWatermark.jpg", bt1, 0777)
}

// 200j.jpg is 320*180
// func TestExecWatermark1(t *testing.T) {
// 	var w Watermark
// 	w.Name = base64.StdEncoding.EncodeToString([]byte("wm1.png"))
// 	w.Location = ""

// 	w.X = 160
// 	w.Y = 90
// 	bt, err := ioutil.ReadFile("200j.jpg")
// 	if err != nil {
// 		t.Error()
// 	}
// 	wand, err := gm.NewMagickWand(bt)
// 	if err != nil {
// 		t.Error()
// 	}
// 	var ctx context.Context
// 	m3 := map[string]string{"n": base64.StdEncoding.EncodeToString([]byte("wm1.png")), "d": "70", "l": ""}
// 	if w.Verify(&ctx, m3) != nil {
// 		t.Error()
// 	}
// 	if w.Exec(&ctx, wand) != nil {
// 		t.Error()
// 	}

// 	bt1, err := wand.WriteBlob()
// 	if err != nil {
// 		t.Error()
// 	}
// 	ioutil.WriteFile("newWatermark1.jpg", bt1, 0777)
// }

// // test x,y both are 0
// func TestExecWatermark2(t *testing.T) {
// 	var w Watermark
// 	bt, err := ioutil.ReadFile("200j.jpg")
// 	if err != nil {
// 		t.Error()
// 	}
// 	wand, err := gm.NewMagickWand(bt)
// 	if err != nil {
// 		t.Error()
// 	}
// 	var ctx context.Context
// 	m3 := map[string]string{"n": base64.StdEncoding.EncodeToString([]byte("wm1.png")), "d": "90", "l": "sw"}
// 	if w.Verify(&ctx, m3) != nil {
// 		t.Error()
// 	}
// 	if w.Exec(&ctx, wand) != nil {
// 		t.Error()
// 	}

// 	bt1, err := wand.WriteBlob()
// 	if err != nil {
// 		t.Error()
// 	}
// 	ioutil.WriteFile("newWatermark2.jpg", bt1, 0777)
// }
