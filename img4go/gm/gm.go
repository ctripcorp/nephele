package gm

import (
	"github.com/ctrip-nephele/gmagick"
)

// Wrapper of gmagick's MagickWand
type MagickWand struct {
	*gmagick.MagickWand
}

// New a MagickWand from the image blob
func NewMagickWand(blob []byte) (*MagickWand, error) {
	mw := &MagickWand{gmagick.NewMagickWand()}
	err := mw.ReadImageBlob(blob)
	// In our practice, GM will throw coder error in decoding some png files.
	// This can be solved by recoding image use Go's image/png package.
	if err != nil && exceptionType(err) == "ERROR_CODER" {
		if newBlob, e := recodePNG(blob); e == nil {
			err = mw.ReadImageBlob(newBlob)
		}
	}
	if err != nil {
		return nil, err
	}
	return mw, nil
}

// AutoOrient adjusts the current image based on EXIF orientation so that it
// is suitable for viewing.
func (mw *MagickWand) AutoOrient() error {
	orientation := mw.GetImageOrientation()
	return mw.AutoOrientImage(orientation)
}

// Composite composites the compositeImage onto current image at position (x, y)
func (mw *MagickWand) Composite(compositeImage *MagickWand, x, y int) error {
	return mw.CompositeImage(compositeImage.MagickWand, gmagick.COMPOSITE_OP_OVER, x, y)
}

// Corp extracts a region of the image
func (mw *MagickWand) Crop(width, height uint, x, y int) error {
	return mw.CropImage(width, height, x, y)
}

// CheckRGBColorspace checks wheather the wand is RGB colorspace
func (mw *MagickWand) CheckRGBColorspace() bool {
	if mw.GetImageColorspace() == gmagick.COLORSPACE_RGB {
		return true
	}
	return false
}

// Format returns the format of a particular image in a sequence.
func (mw *MagickWand) Format() string {
	return mw.GetImageFormat()
}

// Height returns the image height
func (mw *MagickWand) Height() uint {
	return mw.GetImageHeight()
}

// Width returns the image width
func (mw *MagickWand) Width() uint {
	return mw.GetImageWidth()
}

// Size returns the size associated with the magick wand.
func (mw *MagickWand) Size() (columns, rows uint, err error) {
	return mw.Width(), mw.Height(), nil
}

// PreserveJPEGSamplingFactor will use "sampling-factor" settings as the input file
// in encoding output file, if both the input and output are JPEG fromat and sampling factor equals
// to "1x1,1x1,1x1", or default "2x2, 1x1, 1x1" sampling factors will be used.
func (mw *MagickWand) PreserveJPEGSamplingFactor() error {
	factors := mw.GetImageAttribute("JPEG-Sampling-factors")
	if factors == "1x1,1x1,1x1" {
		return mw.SetSamplingFactors([]float64{1, 1})
	}
	return nil
}

// PreserveJPEGSettings will use the same "quality" and "sampling-factor" settings as
// the input file in encoding output file, if both input and output are JPEG format.
// Manual set quality and sampling-factor will be ignored.
func (mw *MagickWand) PreserveJPEGSettings() error {
	return mw.SetImageOption("jpeg", "preserve-settings", "true")
}

// RGBCharPixels extracts char RGB pixel data from the image
func (mw *MagickWand) RGBCharPixels(xOffset, yOffset int, columns, rows uint) ([]byte, error) {
	pixels := make([]byte, columns*rows)
	err := mw.GetImagePixels(xOffset, yOffset, columns, rows, "RGB", gmagick.CharPixel, pixels)
	return pixels, err
}

// CublicResize resizes image use cubic filter
func (mw *MagickWand) CubicResize(cols, rows uint) error {
	return mw.ResizeImage(cols, rows, gmagick.FILTER_CUBIC, 0.5)
}

// LanczosResize resizes image use lanczos filter
func (mw *MagickWand) LanczosResize(cols, rows uint) error {
	return mw.ResizeImage(cols, rows, gmagick.FILTER_LANCZOS, 1.0)
}

// Rotates an image the specified number of degrees
func (mw *MagickWand) Rotate(degree float64) error {
	background := gmagick.NewPixelWand()
	background.SetColor("#000000")
	return mw.RotateImage(background, degree)
}

// Scales the size of an image to the given dimensions.
func (mw *MagickWand) Scale(cols, rows uint) error {
	return mw.ScaleImage(cols, rows)
}

// SetFormat sets the format of the magick wand.
func (mw *MagickWand) SetFormat(format string) error {
	return mw.SetImageFormat(format)
}

// Sets pixel data in the image at the location you specify
func (mw *MagickWand) SetRGBCharPixels(xOffset, yOffset int, columns, rows uint, pixels []byte) error {
	return mw.SetImagePixels(xOffset, yOffset, columns, rows, "RGB", gmagick.CharPixel, pixels)
}

// Strip removes all profiles and text attributes from this image.
func (mw *MagickWand) Strip() error {
	return mw.StripImage()
}

// Sharpen sharpens the image use Gaussian convolution with the
// given radius and standard deviation (sigma)
func (mw *MagickWand) Sharpen(radius, sigma float64) error {
	return mw.SharpenImage(radius, sigma)
}

// WriteBlob returns the image as a byte array (a formatted GIF, JPEG, PNG, etc. in memory),
func (mw *MagickWand) WriteBlob() ([]byte, error) {
	blob := mw.WriteImageBlob()
	if len(blob) == 0 {
		return nil, mw.GetLastError()
	}
	return blob, nil
}
