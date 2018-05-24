package process

import (
	"github.com/ctrip-nephele/gmagick"
)

const (
	imageIsBlob = iota
	imageIsWand
)

type image struct {
	state int
	blob  []byte
	wand  *gmagick.MagickWand
}

func (i *image) Bytes() ([]byte, error) {
	if i.state == imageIsWand {
		i.blob = i.wand.WriteImageBlob()
		if len(i.blob) == 0 {
			return nil, i.wand.GetLastError()
		}
		i.state = imageIsBlob
	}
	return i.blob, nil
}

func (i *image) SetBlob(blob []byte) {
    i.blob = blob
}

func (i *image) Wand() (*gmagick.MagickWand, error) {
	if i.state == imageIsBlob {
		i.wand = gmagick.NewMagickWand()
		err := i.wand.ReadImageBlob(i.blob)
		if err != nil {
			return nil, err
		}
		i.state = imageIsWand
	}
	return i.wand, nil
}
