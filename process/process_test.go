package process

import (
	"testing"
)

func Test_Process(t *testing.T) {
	//check right uri
	uri := "abc.jpg?x-nephele-process=image/resize,w_100,limit_1/crop,w_100,h_100/format,v_png"
	procs, err := parseURI(uri)
	if err != nil {
		t.Error(err)
		return
	}
	procs, err = checkParam(procs)
	if err != nil {
		t.Error(err)
		return
	}
	if len(procs) != 3 &&
		procs[0].Name != Resize && len(procs[0].Param) != 2 &&
		procs[1].Name != Crop && len(procs[1].Param) != 2 &&
		procs[2].Name != Format && len(procs[2].Param) != 1 {
		t.Error("parse uri failed.")
		return
	}

	//check invalid uri
	uri = "abc.jpg?x-nephele-process=image/resize,w_a,limit_1/crop,w_100,h_100/format,v_png"
	procs, err = parseURI(uri)
	if err != nil {
		t.Error(err)
		return
	}
	procs, err = checkParam(procs)
	if err == nil {
		t.Error("check param failed.")
	}
}
