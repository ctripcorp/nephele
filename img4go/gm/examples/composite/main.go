package main

import (
	"github.com/giantpoplar/nephele/img4go/gm"
	"io/ioutil"
	"os"
)

func main() {
	source, _ := ioutil.ReadFile("../images/src.jpg")
	watermark, _ := ioutil.ReadFile("../images/ht1.png")
	sourceMW, err := gm.NewMagickWand(source)
	if err != nil {
		panic(err)
	}
	watermarkMW, err := gm.NewMagickWand(watermark)
	if err != nil {
		panic(err)
	}

	// crop source image to size 400x400 at position (100, 100)
	err = sourceMW.Crop(400, 400, 100, 100)
	if err != nil {
		panic(err)
	}

	// scale watermark image to size 100x100
	err = watermarkMW.Scale(100, 100)
	if err != nil {
		panic(err)
	}

	// set composite dissolve percent 50
	err = watermarkMW.Dissolve(50)
	if err != nil {
		panic(err)
	}

	// composite watermark to srouce at position (200, 100)
	err = sourceMW.Composite(watermarkMW, 200, 100)
	if err != nil {
		panic(err)
	}

	err = sourceMW.GWand().DisplayImage(os.Getenv("DISPLAY"))
	if err != nil {
		panic(err)
	}

}
