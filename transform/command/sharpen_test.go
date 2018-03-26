package command

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
)

func TestSharpen(t *testing.T) {
	var s Sharpen
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

	m3 := map[string]string{"r": "10", "s": "122"}
	if s.Verify(&ctx, m3) != nil {
		t.Error()
	}
	if s.Exec(&ctx, wand) != nil {
		fmt.Println(s.Exec(&ctx, wand))
		t.Error()
	}

	blob, err := wand.WriteBlob()
	if err != nil {
		t.Error(err)
		return
	}
	ioutil.WriteFile("newSharpen.jpg", blob, 0777)

}
