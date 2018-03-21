package command

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
)

func TestExec(t *testing.T) {
	var r Resize
	bt, err := ioutil.ReadFile("200j.jpg")
	if err != nil {
		t.Error()
	}
	wand, err := gm.NewMagickWand(bt)
	if err != nil {
		fmt.Println("gm", err.Error())
		t.Error()
	}
	r.Width = uint(1000)
	r.Height = uint(300)
	r.Method = "lfit"
	r.Limit = true
	r.Percentage = 500

	var ctx context.Context
	r.Exec(&ctx, wand)
	if wand.Width() != 4000 || wand.Height() != 2625 {
		println(wand.Width(), wand.Height())
		t.Fail()
	}
}
