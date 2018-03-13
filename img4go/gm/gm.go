package gm

import "github.com/happysheeps/gmagick"

// Wrapper of gmagick's MagickWand
type MagickWand struct {
	mw *gmagick.MagickWand
}

// New a MagickWand from the image blob
func NewMagickWand(blob []byte) (*MagickWand, error) {
	mw := &MagickWand{}
	mw.mw = gmagick.NewMagickWand()
	err := mw.mw.ReadImageBlob(blob)
	// In our practice, GM will throw coder error in decoding some png files.
	// This can be solved by recoding image use Go's image/png package.
	if err != nil && exceptionType(err) == "ERROR_CODER" {
		if newBlob, e := pngDeEncode(blob); e == nil {
			err = mw.mw.ReadImageBlob(newBlob)
		}
	}
	if err != nil {
		return nil, err
	}
	return mw, nil
}

// GWand returns raw gmagick wand
func (mw *MagickWand) GWand() *gmagick.MagickWand {
	return mw.mw
}

// AutoOrient adjusts the current image based on EXIF orientation so that it
// is suitable for viewing.
func (mw *MagickWand) AutoOrient() error {
	orientation := mw.mw.GetImageOrientation()
	return mw.mw.AutoOrientImage(orientation)
}

// Composite composites the compositeImage onto current image at position (x, y)
func (mw *MagickWand) Composite(compositeImage *MagickWand, x, y int) error {
	return mw.mw.CompositeImage(compositeImage.mw, gmagick.COMPOSITE_OP_OVER, x, y)
}

// Corp extracts a region of the image
func (mw *MagickWand) Crop(width, height uint, x, y int) error {
	return mw.mw.CropImage(width, height, x, y)
}

// CheckRGBColorspace checks wheather the wand is RGB colorspace
func (mw *MagickWand) CheckRGBColorspace() bool {
	if mw.mw.GetImageColorspace() == gmagick.COLORSPACE_RGB {
		return true
	}
	return false
}

// Dissolve sets current image dissolved percent in composite.
// It multiplies current image opacity by the percent.
//
// percent: range from 0(completely transparent)-100(opacity)
func (mw *MagickWand) Dissolve(percent int) error {
	return mw.mw.Dissolve(percent)
}

// Format returns the format of a particular image in a sequence.
func (mw *MagickWand) Format() string {
	return mw.mw.GetImageFormat()
}

// Height returns the image height
func (mw *MagickWand) Height() uint {
	return mw.mw.GetImageHeight()
}

// Width returns the image width
func (mw *MagickWand) Width() uint {
	return mw.mw.GetImageWidth()
}

// Size returns the size associated with the magick wand.
func (mw *MagickWand) Size() (columns, rows uint, err error) {
	columns, rows, err = mw.mw.GetSize()
	return
}

// PreserveJPEGSamplingFactor will use "sampling-factor" settings as the input file
// in encoding output file, if both the input and output are JPEG fromat and sampling factor equals
// to "1x1,1x1,1x1", or default "2x2, 1x1, 1x1" sampling factors will be used.
func (mw *MagickWand) PreserveJPEGSamplingFactor() error {
	factors := mw.mw.GetImageAttribute("JPEG-Sampling-factors")
	if factors == "1x1,1x1,1x1" {
		return mw.mw.SetSamplingFactors([]float64{1, 1})
	}
	return nil
}

// PreserveJPEGSettings will use the same "quality" and "sampling-factor" settings as
// the input file in encoding output file, if both input and output are JPEG format.
// Manual set quality and sampling-factor will be ignored.
func (mw *MagickWand) PreserveJPEGSettings() error {
	return mw.mw.SetImageOption("jpeg", "preserve-settings", "true")
}

// RGBCharPixels extracts char RGB pixel data from the image
func (mw *MagickWand) RGBCharPixels(xOffset, yOffset int, columns, rows uint) ([]byte, error) {
	pixels := make([]byte, columns*rows)
	err := mw.mw.GetImagePixels(xOffset, yOffset, columns, rows, "RGB", gmagick.CharPixel, pixels)
	return pixels, err
}

// CublicResize resizes image use cubic filter
func (mw *MagickWand) CubicResize(cols, rows uint) error {
	return mw.mw.ResizeImage(cols, rows, gmagick.FILTER_CUBIC, 0.5)
}

// LanczosResize resizes image use lanczos filter
func (mw *MagickWand) LanczosResize(cols, rows uint) error {
	return mw.mw.ResizeImage(cols, rows, gmagick.FILTER_LANCZOS, 1.0)
}

// Rotates an image the specified number of degrees
func (mw *MagickWand) Rotate(degree float64) error {
	background := gmagick.NewPixelWand()
	background.SetColor("#000000")
	return mw.mw.RotateImage(background, degree)
}

// Scales the size of an image to the given dimensions.
func (mw *MagickWand) Scale(cols, rows uint) error {
	return mw.mw.ScaleImage(cols, rows)
}

// SetCompressionQuality sets the wand compression quality.
func (mw *MagickWand) SetCompressionQuality(quality uint) error {
	return mw.mw.SetCompressionQuality(quality)
}

// SetFormat sets the format of the magick wand.
func (mw *MagickWand) SetFormat(format string) error {
	return mw.mw.SetImageFormat(format)
}

// Sets pixel data in the image at the location you specify
func (mw *MagickWand) SetRGBCharPixels(xOffset, yOffset int, columns, rows uint, pixels []byte) error {
	return mw.mw.SetImagePixels(xOffset, yOffset, columns, rows, "RGB", gmagick.CharPixel, pixels)
}

// Strip removes all profiles and text attributes from this image.
func (mw *MagickWand) Strip() error {
	return mw.mw.StripImage()
}

// Sharpen sharpens the image use Gaussian convolution with the
// given radius and standard deviation (sigma)
func (mw *MagickWand) Sharpen(radius, sigma float64) error {
	return mw.mw.SharpenImage(radius, sigma)
}

// WriteBlob returns the image as a byte array (a formatted GIF, JPEG, PNG, etc. in memory),
func (mw *MagickWand) WriteBlob() ([]byte, error) {
	blob := mw.mw.WriteImageBlob()
	if len(blob) == 0 {
		return nil, mw.mw.GetLastError()
	}
	return blob, nil
}
