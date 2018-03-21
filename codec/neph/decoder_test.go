package neph

import "testing"

func Test_Decode(t *testing.T) {
	decoder := Decoder{}
	if err := decoder.Decode("abc.jpg?x-nephele-process=image/resize,w_100,limit_1/crop,w_100,h_100/format,v_png"); err != nil {
		t.Error(err)
		return
	}

	if decoder.path != "abc.jpg" {
		t.Error("decode path failed.")
		return
	}

	if err := decoder.Decode("abc.jpg?x-nephele-process=image/resize,w_abc,limit_1/crop,w_100,h_100/format,v_png"); err == nil {
		t.Error("check resize failed!")
		return
	}
}
