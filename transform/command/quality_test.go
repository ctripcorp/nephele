package command

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
)

func TestVerify(t *testing.T) {
	q := &Quality{}
	if e := q.Verify(context.New("", time.Duration(time.Second*3)), map[string]string{"v": "10"}); e != nil {
		t.Error(e)
	}
}

func TestQuality(t *testing.T) {
	var q Quality
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

	m3 := map[string]string{"v": "1"}
	if e := q.Verify(&ctx, m3); e != nil {
		t.Error(e)
	}
	if e := q.Exec(&ctx, wand); e != nil {
		t.Error(e)
	}

	blob, err := wand.WriteBlob()
	if err != nil {
		t.Error(err)
		return
	}
	ioutil.WriteFile("newQuality.jpg", blob, 0777)
}
