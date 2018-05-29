package main

import (
	"github.com/ctripcorp/nephele/img4go/gm"
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
	err = mw.CubicResize(hWidth, hHeight)
	if err != nil {
		panic(err)
	}

	// Auto orient based on exif info orientation
	err = mw.AutoOrient()
	if err != nil {
		panic(err)
	}

	// Set output format webp
	err = mw.SetFormat("webp")
	if err != nil {
		panic(err)
	}

	// Set the compression quality to 30 (high quality = low compression)
	err = mw.SetCompressionQuality(30)
	if err != nil {
		panic(err)
	}

	// Strip extra info
	err = mw.Strip()
	if err != nil {
		panic(err)
	}

	err = mw.DisplayImage(os.Getenv("DISPLAY"))
	if err != nil {
		panic(err)
	}

	//b, err := mw.WriteBlob()
	//if err != nil {
	//	panic(err)
	//}
	//
	//ioutil.WriteFile("../images/output.webp", b, 0666)

}
