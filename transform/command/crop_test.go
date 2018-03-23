package command

import "testing"

func TestVerfiy(t *testing.T) {
	type check struct {
		Param map[string]string
		IsErr bool
	}
	checks := []check{
		check{map[string]string{"m": "t", "p": "10"}, false},
		check{map[string]string{"m": "b", "h": "10"}, false},
		check{map[string]string{"m": "l", "p": "10"}, false},
		check{map[string]string{"m": "r", "w": "10"}, false},
		check{map[string]string{"m": "wc", "w": "10"}, false},
		check{map[string]string{"m": "hc", "h": "10"}, false},
		check{map[string]string{"m": "c", "w": "10"}, false},
		check{map[string]string{"m": "crop", "w": "10", "h": "100"}, false},
		check{map[string]string{"m": "resize", "w": "100"}, false},
		check{map[string]string{"m": "b", "p": "100"}, true},
		check{map[string]string{"m": "l", "x": "10"}, true},
		check{map[string]string{"m": "r", "h": "10"}, true},
		check{map[string]string{"m": "wc", "x": "10"}, true},
		check{map[string]string{"m": "t", "w": "100"}, true},
	}
	cropCommand := &Crop{}
	for i, c := range checks {
		err := cropCommand.Verfiy(nil, c.Param)
		if c.IsErr == (err == nil) {
			t.Error("index:", i, " check failed.")
			return
		}
	}
}

func TestM(t *testing.T) {
	type param struct {
		w, h, srcW, srcH uint
		p                int
	}
	type result struct {
		width, height uint
		x, y          int
	}
	type checkH struct {
		Param  param
		Result result
		f      func(w, srcW, srcH uint, p int) (width, height uint, x, y int)
	}
	//cropMT
	checksH := []checkH{
		{f: cropMT, Param: param{h: 100, srcW: 500, srcH: 500}, Result: result{width: 500, height: 400, x: 0, y: 100}},
		{f: cropMT, Param: param{p: 10, srcW: 500, srcH: 500}, Result: result{width: 500, height: 450, x: 0, y: 50}},
		{f: cropMB, Param: param{h: 100, srcW: 500, srcH: 500}, Result: result{width: 500, height: 400, x: 0, y: 0}},
		{f: cropMB, Param: param{p: 10, srcW: 500, srcH: 500}, Result: result{width: 500, height: 450, x: 0, y: 0}},
		{f: cropMHC, Param: param{h: 100, srcW: 500, srcH: 500}, Result: result{width: 500, height: 400, x: 0, y: 50}},
		{f: cropMHC, Param: param{p: 10, srcW: 500, srcH: 500}, Result: result{width: 500, height: 450, x: 0, y: 25}},
	}
	for i, c := range checksH {
		width, height, x, y := c.f(c.Param.h, c.Param.srcW, c.Param.srcH, c.Param.p)
		if !(width == c.Result.width && height == c.Result.height && x == c.Result.x && y == c.Result.y) {
			t.Error("index:", i, " check failed.")
			return
		}
	}

	type checkW struct {
		Param  param
		Result result
		f      func(w, srcW, srcH uint, p int) (width, height uint, x, y int)
	}
	//cropMT
	checksW := []checkW{
		{f: cropML, Param: param{w: 100, srcW: 500, srcH: 500}, Result: result{width: 400, height: 500, x: 100, y: 0}},
		{f: cropML, Param: param{p: 10, srcW: 500, srcH: 500}, Result: result{width: 450, height: 500, x: 50, y: 0}},
		{f: cropMR, Param: param{w: 100, srcW: 500, srcH: 500}, Result: result{width: 400, height: 500, x: 0, y: 0}},
		{f: cropMR, Param: param{p: 10, srcW: 500, srcH: 500}, Result: result{width: 450, height: 500, x: 0, y: 0}},
		{f: cropMWC, Param: param{w: 100, srcW: 500, srcH: 500}, Result: result{width: 400, height: 500, x: 50, y: 0}},
		{f: cropMWC, Param: param{p: 10, srcW: 500, srcH: 500}, Result: result{width: 450, height: 500, x: 25, y: 0}},
	}
	for i, c := range checksW {
		width, height, x, y := c.f(c.Param.w, c.Param.srcW, c.Param.srcH, c.Param.p)
		if !(width == c.Result.width && height == c.Result.height && x == c.Result.x && y == c.Result.y) {
			t.Error("index:", i, " check failed.")
			return
		}
	}

	type checkC struct {
		Param  param
		Result result
		f      func(w, h, srcW, srcH uint, p int) (width, height uint, x, y int)
	}

}
