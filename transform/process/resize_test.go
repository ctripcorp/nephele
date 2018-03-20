package process

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
)

func TestExec(t *testing.T) {
	var r ResizeCommand
	bt, err := ioutil.ReadFile("200j.jpg")
	if err != nil {
		t.Error()
	}
	r.Wand, err = gm.NewMagickWand(bt)
	if err != nil {
		fmt.Println("gm", err.Error())
		t.Error()
	}
	r.Width = uint(1000)
	r.Height = uint(300)
	r.Method = "lfit"
	r.Limit = 1
	r.Percentage = 500

	var ctx context.Context
	r.Exec(ctx)
	if r.Wand.Width() != 4000 || r.Wand.Height() != 2625 {
		println(r.Wand.Width(), r.Wand.Height())
		t.Fail()
	}
}

func TestExec1(t *testing.T) {
	var r ResizeCommand
	bt, err := ioutil.ReadFile("200j.jpg")
	if err != nil {
		t.Error()
	}
	r.Wand, err = gm.NewMagickWand(bt)
	if err != nil {
		fmt.Println("gm", err.Error())
		t.Error()
	}
	r.Width = uint(1000)
	r.Height = uint(300)
	r.Method = "lfit"
	r.Limit = 1
	r.Percentage = 50

	var ctx context.Context
	r.Exec(ctx)

	if r.Wand.Width() != 399 || r.Wand.Height() != 262 {
		println(r.Wand.Width(), r.Wand.Height())
		t.Fail()
	}
}

func TestExec2(t *testing.T) {
	var r ResizeCommand
	bt, err := ioutil.ReadFile("200j.jpg")
	if err != nil {
		t.Error()
	}
	r.Wand, err = gm.NewMagickWand(bt)
	if err != nil {
		fmt.Println("gm", err.Error())
		t.Error()
	}
	r.Width = uint(100)
	r.Height = uint(300)
	r.Method = "lfit"
	r.Limit = 1
	r.Percentage = 0

	var ctx context.Context
	r.Exec(ctx)

	if r.Wand.Width() != 100 || r.Wand.Height() != 65 {
		println(r.Wand.Width(), r.Wand.Height())
		t.Fail()
	}
}
