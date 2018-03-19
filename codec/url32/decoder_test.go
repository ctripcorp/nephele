package url32

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

	if len(decoder.processes) != 3 &&
		decoder.processes[0].Name != Resize && len(decoder.processes[0].Param) != 2 &&
		decoder.processes[1].Name != Crop && len(decoder.processes[1].Param) != 2 &&
		decoder.processes[2].Name != Format && len(decoder.processes[2].Param) != 1 {
		t.Error("parse uri failed.")
		return
	}
}
