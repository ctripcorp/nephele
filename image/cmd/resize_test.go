package cmd

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/nephele/context"
	"github.com/nephele/img4go/gm"
)

// type ResizeCommand struct {
// 	Wand              *gm.MagickWand
// 	Width             uint
// 	Height            uint
// 	Method            string //Lfit/Fixed
// 	Limit, Percentage int
// }

func TestExec(t *testing.T) {
	var r ResizeCommand
	bt, err := ioutil.ReadFile("200j.jpg")
	if err != nil {
		return
	}
	r.Wand, err = gm.NewMagickWand(bt)
	if err != nil {
		fmt.Println("gm", err.Error())
	}
	r.Width = uint(1000)
	r.Height = uint(300)
	r.Method = "lfit"
	r.Limit = 1
	r.Percentage = 500

	var ctx context.Context

	r.Exec(ctx)

	bt1, err := r.Wand.WriteBlob()
	if err != nil {
		fmt.Println(err.Error())
	}
	ioutil.WriteFile("new.jpg", bt1, 0777)

}

func TestExec1(t *testing.T) {
	var r ResizeCommand
	bt, err := ioutil.ReadFile("200j.jpg")
	if err != nil {
		return
	}
	r.Wand, err = gm.NewMagickWand(bt)
	if err != nil {
		fmt.Println("gm", err.Error())
	}
	r.Width = uint(1000)
	r.Height = uint(300)
	r.Method = "lfit"
	r.Limit = 1
	r.Percentage = 50

	var ctx context.Context

	r.Exec(ctx)

	bt1, err := r.Wand.WriteBlob()
	if err != nil {
		fmt.Println(err.Error())
	}
	ioutil.WriteFile("new1.jpg", bt1, 0777)
}

func TestExec2(t *testing.T) {
	var r ResizeCommand
	bt, err := ioutil.ReadFile("200j.jpg")
	if err != nil {
		return
	}
	r.Wand, err = gm.NewMagickWand(bt)
	if err != nil {
		fmt.Println("gm", err.Error())
	}
	r.Width = uint(100)
	r.Height = uint(300)
	r.Method = "lfit"
	r.Limit = 1
	r.Percentage = 0

	var ctx context.Context

	r.Exec(ctx)

	bt1, err := r.Wand.WriteBlob()
	if err != nil {
		fmt.Println(err.Error())
	}
	ioutil.WriteFile("new2.jpg", bt1, 0777)
}
