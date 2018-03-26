package main

import (
	"github.com/giantpoplar/nephele/img4go/gm"
	"io/ioutil"
	"os"
)

func main() {
	blob, err := ioutil.ReadFile("../images/src.jpg")
	if err != nil {
		panic(err)
	}
	mw, err := gm.NewMagickWand(blob)
	if err != nil {
		panic(err)
	}

	// get src image size
	width := mw.Width()
	height := mw.Height()

	// caluculate half the size
	hWidth := uint(width / 2)
	hHeight := uint(height / 2)

	// Resize the image using the Lanczos filter
	err = mw.LanczosResize(hWidth, hHeight)
	if err != nil {
		panic(err)
	}

	// Set the compression quality to 95 (high quality = low compression)
	err = mw.SetCompressionQuality(90)
	if err != nil {
		panic(err)
	}

	err = mw.GWand().DisplayImage(os.Getenv("DISPLAY"))
	if err != nil {
		panic(err)
	}

}
