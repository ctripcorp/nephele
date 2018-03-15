package cmd

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/nephele/context"
	"github.com/nephele/img4go/gm"
)

func TestExecStrip(t *testing.T) {
	var r StripCommand
	bt, err := ioutil.ReadFile("200j.jpg")
	if err != nil {
		return
	}
	r.Wand, err = gm.NewMagickWand(bt)
	if err != nil {
		fmt.Println("gm", err.Error())
	}

	var ctx context.Context

	r.Exec(ctx)

	bt1, err := r.Wand.WriteBlob()
	if err != nil {
		fmt.Println(err.Error())
	}
	ioutil.WriteFile("newStrip.jpg", bt1, 0777)
}
