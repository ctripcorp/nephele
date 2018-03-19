// +build darwin,cgo
package gm

// Wrapper of gmagick's MagickWand
type MagickWand struct {
}

// New a MagickWand from the image blob
func NewMagickWand(blob []byte) (*MagickWand, error) {
	return nil, nil
}

// AutoOrient adjusts the current image based on EXIF orientation so that it
// is suitable for viewing.
func (mw *MagickWand) AutoOrient() error {
	return nil
}

// Composite composites the compositeImage onto current image at position (x, y)
func (mw *MagickWand) Composite(compositeImage *MagickWand, x, y int) error {
	return nil
}

// Corp extracts a region of the image
func (mw *MagickWand) Crop(width, height uint, x, y int) error {
	return nil
}

// CheckRGBColorspace checks wheather the wand is RGB colorspace
func (mw *MagickWand) CheckRGBColorspace() bool {
	return false
}

// Dissolve sets current image dissolved percent in composite.
// It multiplies current image opacity by the percent.
//
// percent: range from 0(completely transparent)-100(opacity)
func (mw *MagickWand) Dissolve(percent int) error {
	return nil
}

// Format returns the format of a particular image in a sequence.
func (mw *MagickWand) Format() string {
	return ""
}

// Height returns the image height
func (mw *MagickWand) Height() uint {
	return 0
}

// Width returns the image width
func (mw *MagickWand) Width() uint {
	return 0
}

// Size returns the size associated with the magick wand.
func (mw *MagickWand) Size() (columns, rows uint, err error) {
	return 0, 0, nil
}

// PreserveJPEGSamplingFactor will use "sampling-factor" settings as the input file
// in encoding output file, if both the input and output are JPEG fromat and sampling factor equals
// to "1x1,1x1,1x1", or default "2x2, 1x1, 1x1" sampling factors will be used.
func (mw *MagickWand) PreserveJPEGSamplingFactor() error {
	return nil
}

// PreserveJPEGSettings will use the same "quality" and "sampling-factor" settings as
// the input file in encoding output file, if both input and output are JPEG format.
// Manual set quality and sampling-factor will be ignored.
func (mw *MagickWand) PreserveJPEGSettings() error {
	return nil
}

// RGBCharPixels extracts char RGB pixel data from the image
func (mw *MagickWand) RGBCharPixels(xOffset, yOffset int, columns, rows uint) ([]byte, error) {
	return nil, nil
}

// CublicResize resizes image use cubic filter
func (mw *MagickWand) CubicResize(cols, rows uint) error {
	return nil
}

// LanczosResize resizes image use lanczos filter
func (mw *MagickWand) LanczosResize(cols, rows uint) error {
	return nil
}

// Rotates an image the specified number of degrees
func (mw *MagickWand) Rotate(degree float64) error {
	return nil
}

// Scales the size of an image to the given dimensions.
func (mw *MagickWand) Scale(cols, rows uint) error {
	return nil
}

// SetCompressionQuality sets the wand compression quality.
func (mw *MagickWand) SetCompressionQuality(quality uint) error {
	return nil
}

// SetFormat sets the format of the magick wand.
func (mw *MagickWand) SetFormat(format string) error {
	return nil
}

// Sets pixel data in the image at the location you specify
func (mw *MagickWand) SetRGBCharPixels(xOffset, yOffset int, columns, rows uint, pixels []byte) error {
	return nil
}

// Strip removes all profiles and text attributes from this image.
func (mw *MagickWand) Strip() error {
	return nil
}

// Sharpen sharpens the image use Gaussian convolution with the
// given radius and standard deviation (sigma)
func (mw *MagickWand) Sharpen(radius, sigma float64) error {
	return nil
}

// WriteBlob returns the image as a byte array (a formatted GIF, JPEG, PNG, etc. in memory),
func (mw *MagickWand) WriteBlob() ([]byte, error) {
	return nil, nil
}
