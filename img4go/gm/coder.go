package gm

import (
	"bytes"
	"image/png"
)

// pngDeEncode do a decode and encode routine of src image blob
func pngDeEncode(blob []byte) ([]byte, error) {
	img, err := png.Decode(bytes.NewBuffer(blob))
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	if err = png.Encode(buf, img); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
