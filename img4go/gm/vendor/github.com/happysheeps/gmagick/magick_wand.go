package gmagick

/*
#include <unistd.h>
#include <wand/wand_api.h>
*/
import "C"

import (
	"errors"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"syscall"
	"unsafe"
)

// This struct represents the MagickWand C API of GraphicsMagick
type MagickWand struct {
	mw   *C.MagickWand
	init sync.Once
}

func newMagickWand(cmw *C.MagickWand) *MagickWand {
	mw := &MagickWand{mw: cmw}
	runtime.SetFinalizer(mw, Destroy)
	mw.IncreaseCount()

	return mw
}

// Returns a wand required for all other methods in the API. A fatal exception is thrown if there is not enough memory to allocate the wand.
func NewMagickWand() *MagickWand {
	return newMagickWand(C.NewMagickWand())
}

// Returns low level C MagickWand
func (mw *MagickWand) CWand() *C.MagickWand {
	return mw.mw
}


// Makes an exact copy of the MagickWand object
func (mw *MagickWand) Clone() *MagickWand {
	return newMagickWand(C.CloneMagickWand(mw.mw))
}

// Deallocates memory associated with an MagickWand
func (mw *MagickWand) Destroy() {
	if mw.mw == nil {
		return
	}

	mw.init.Do(func() {
		C.DestroyMagickWand(mw.mw)
		mw.mw = nil
		mw.DecreaseCount()
	})
}

// Increase MagickWand ref counter and set according "can`t be terminated status"
func (mw *MagickWand) IncreaseCount() {
	atomic.AddInt64(&magickWandCounter, int64(1))
	unsetCanTerminate()
}

// Decrease MagickWand ref counter and set according "can be terminated status"
func (mw *MagickWand) DecreaseCount() {
	atomic.AddInt64(&magickWandCounter, int64(-1))
	setCanTerminate()
}

// Returns the position of the iterator in the image list
func (mw *MagickWand) GetImageIndex() uint {
	return uint(C.MagickGetImageIndex(mw.mw))
}

// Selects an individual threshold for each pixel based on the range of
// intensity values in its local neighborhood. This allows for thresholding
// of an image whose global intensity histogram doesn't contain distinctive
// peaks.
func (mw *MagickWand) AdaptiveThresholdImage(width, height uint, offset int) error {
	ok := C.MagickAdaptiveThresholdImage(mw.mw, C.ulong(width), C.ulong(height), C.long(offset))
	return mw.getLastErrorIfFailed(ok)
}

// Adds a clone of the images from the second wand and inserts them into the
// first wand. Use SetLastIterator(), to append new images into an existing
// wand, current image will be set to last image so later adds with also be
// appened to end of wand. Use SetFirstIterator() to prepend new images into
// wand, any more images added will also be prepended before other images in
// the wand. However the order of a list of new images will not change.
// Otherwise the new images will be inserted just after the current image, and
// any later image will also be added after this current image but before the
// previously added images. Caution is advised when multiple image adds are
// inserted into the middle of the wand image list.
func (mw *MagickWand) AddImage(wand *MagickWand) error {
	ok := C.MagickAddImage(mw.mw, wand.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Adds random noise to the image
func (mw *MagickWand) AddNoiseImage(noiseType NoiseType) error {
	ok := C.MagickAddNoiseImage(mw.mw, C.NoiseType(noiseType))
	return mw.getLastErrorIfFailed(ok)
}

// Transforms an image as dictaded by the affine matrix of the drawing wand
func (mw *MagickWand) AffineTransformImage(drawingWand *DrawingWand) error {
	ok := C.MagickAffineTransformImage(mw.mw, drawingWand.dw)
	return mw.getLastErrorIfFailed(ok)
}

// Animates an image or image sequence
func (mw *MagickWand) AnimateImages(server string) error {
	csserver := C.CString(server)
	defer C.free(unsafe.Pointer(csserver))
	ok := C.MagickAnimateImages(mw.mw, csserver)
	return mw.getLastErrorIfFailed(ok)
}

// Annotates an image with text
//
// x: ordinate to left of text
//
// y: ordinate to text baseline
//
// angle: rotate text relative to this angle
//
func (mw *MagickWand) AnnotateImage(drawingWand *DrawingWand, x, y, angle float64, text string) error {
	cstext := C.CString(text)
	defer C.free(unsafe.Pointer(cstext))
	ok := C.MagickAnnotateImage(mw.mw, drawingWand.dw, C.double(x), C.double(y), C.double(angle), cstext)
	return mw.getLastErrorIfFailed(ok)
}

// Append the images in a wand from the current image onwards, creating a new
// wand with the single image result. This is affected by the gravity and
// background setting of the first image. Typically you would call either
// ResetIterator() or SetFirstImage() before calling this function to ensure
// that all the images in the wand's image list will be appended together.
// By default, images are stacked left-to-right. Set topToBottom to true to
// stack them top-to-bottom.
func (mw *MagickWand) AppendImages(topToBottom bool) *MagickWand {
	return newMagickWand(C.MagickAppendImages(mw.mw, b2i(topToBottom)))
}

// adjusts the current image so that its orientation is suitable for viewing
// (i.e. top-left orientation).
//
// currentOrientation: Current image orientation
func (mw *MagickWand) AutoOrientImage(currentOrientation OrientationType) error {
	ok := C.MagickAutoOrientImage(mw.mw, C.OrientationType(currentOrientation))
	return mw.getLastErrorIfFailed(ok)
}

// Average a set of images
func (mw *MagickWand) AverageImages() *MagickWand {
	return newMagickWand(C.MagickAverageImages(mw.mw))
}

// This is like ThresholdImage() but forces all pixels below the threshold
// into black while leaving all pixels above the threshold unchanged.
func (mw *MagickWand) BlackThresholdImage(threshold *PixelWand) error {
	ok := C.MagickBlackThresholdImage(mw.mw, threshold.pw)
	return mw.getLastErrorIfFailed(ok)
}

// Blurs an image. We convolve the image with a gaussian operator of the
// given radius and standard deviation (sigma). For reasonable results, the
// radius should be larger than sigma. Use a radius of 0 and BlurImage()
// selects a suitable radius for you.
//
// radius: the radius of the, in pixels, not counting the center pixel.
//
// sigma: the standard deviation of the, in pixels
//
func (mw *MagickWand) BlurImage(radius, sigma float64) error {
	ok := C.MagickBlurImage(mw.mw, C.double(radius), C.double(sigma))
	return mw.getLastErrorIfFailed(ok)
}

// Surrounds the image with a border of the color defined by the bordercolor
// pixel wand.
func (mw *MagickWand) BorderImage(borderColor *PixelWand, width, height uint) error {
	ok := C.MagickBorderImage(mw.mw, borderColor.pw, C.ulong(width), C.ulong(height))
	return mw.getLastErrorIfFailed(ok)
}

// Simulates a charcoal drawing
//
// radius: the radius of the Gaussian, in pixels, not counting the center pixel
//
// sigma: the standard deviation of the Gaussian, in pixels
//
func (mw *MagickWand) CharcoalImage(radius, sigma float64) error {
	ok := C.MagickCharcoalImage(mw.mw, C.double(radius), C.double(sigma))
	return mw.getLastErrorIfFailed(ok)
}

// Removes a region of an image and collapses the image to occupy the removed
// portion.
//
// width, height: the region width and height
//
// x, y: the region x and y offsets
//
func (mw *MagickWand) ChopImage(width, height uint, x, y int) error {
	ok := C.MagickChopImage(mw.mw, C.ulong(width), C.ulong(height), C.long(x), C.long(y))
	return mw.getLastErrorIfFailed(ok)
}

// Clips along the first path from the 8BIM profile, if present
func (mw *MagickWand) ClipImage() error {
	ok := C.MagickClipImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Composites a set of images while respecting any page offsets and disposal
// methods. GIF, MIFF, and MNG animation sequences typically start with an
// image background and each subsequent image varies in size and offset.
// CoalesceImages() returns a new sequence where each image in the sequence
// is the same size as the first and composited with the next image in the
// sequence.
func (mw *MagickWand) CoalesceImages() *MagickWand {
	return newMagickWand(C.MagickCoalesceImages(mw.mw))
}

// Blends the fill color with each pixel in the image
func (mw *MagickWand) ColorizeImage(colorize, opacity *PixelWand) error {
	ok := C.MagickColorizeImage(mw.mw, colorize.pw, opacity.pw)
	return mw.getLastErrorIfFailed(ok)
}

// Adds a comment to your image
func (mw *MagickWand) CommentImage(comment string) error {
	cscomment := C.CString(comment)
	defer C.free(unsafe.Pointer(cscomment))
	ok := C.MagickCommentImage(mw.mw, cscomment)
	return mw.getLastErrorIfFailed(ok)
}

// Compares one or more image channels of an image to a reconstructed image
// and returns the difference image
func (mw *MagickWand) CompareImageChannels(reference *MagickWand, channel ChannelType, metric MetricType) (wand *MagickWand, distortion float64) {
	cmw := C.MagickCompareImageChannels(mw.mw, reference.mw, C.ChannelType(channel), C.MetricType(metric), (*C.double)(&distortion))
	wand = newMagickWand(cmw)
	return
}

// CompareImages() compares an image to a reconstructed image and returns the
// specified difference image. Returns the new MagickWand and the computed
// distortion between the images
func (mw *MagickWand) CompareImages(reference *MagickWand, metric MetricType) (wand *MagickWand, distortion float64) {
	cmw := C.MagickCompareImages(mw.mw, reference.mw, C.MetricType(metric), (*C.double)(&distortion))
	wand = newMagickWand(cmw)
	return
}

// Composite one image onto another at the specified offset.
// source: The magick wand holding source image.
// compose: This operator affects how the composite is applied to the image.
// The default is Over.
//
// x: the column offset of the composited image.
//
// y: the row offset of the composited image.
//
func (mw *MagickWand) CompositeImage(source *MagickWand, compose CompositeOperator, x, y int) error {
	ok := C.MagickCompositeImage(mw.mw, source.mw, C.CompositeOperator(compose), C.long(x), C.long(y))
	return mw.getLastErrorIfFailed(ok)
}

// Enhances the intensity differences between the lighter and darker elements
// of the image. Set sharpen to a value other than 0 to increase the image
// contrast otherwise the contrast is reduced.
//
// sharpen: increase or decrease image contrast
//
func (mw *MagickWand) ContrastImage(sharpen bool) error {
	ok := C.MagickContrastImage(mw.mw, b2i(sharpen))
	return mw.getLastErrorIfFailed(ok)
}

// Applies a custom convolution kernel to the image.
//
// order: the number of cols and rows in the filter kernel
//
// kernel: an array of doubles, representing the convolution kernel
//
func (mw *MagickWand) ConvolveImage(order uint, kernel []float64) error {
	ok := C.MagickConvolveImage(mw.mw, C.ulong(order), (*C.double)(&kernel[0]))
	return mw.getLastErrorIfFailed(ok)
}

// Extracts a region of the image
func (mw *MagickWand) CropImage(width, height uint, x, y int) error {
	ok := C.MagickCropImage(mw.mw, C.ulong(width), C.ulong(height), C.long(x), C.long(y))
	return mw.getLastErrorIfFailed(ok)
}

// Displaces an Image's colormap by a given number of positions. If you cycle
// the colormap a number of times you can produce a psychodelic effect.
func (mw *MagickWand) CycleColormapImage(displace int) error {
	ok := C.MagickCycleColormapImage(mw.mw, C.long(displace))
	return mw.getLastErrorIfFailed(ok)
}

// Compares each image with the next in a sequence and returns the maximum
// bouding region of any pixel differences it discovers.
func (mw *MagickWand) DeconstructImages() *MagickWand {
	return newMagickWand(C.MagickDeconstructImages(mw.mw))
}

// Describes an image by formatting its attributes
// to an allocated string which must be freed by the user.  Attributes
// include the image width, height, size, and others.  The string is
// similar to the output of 'identify -verbose'.
func (mw *MagickWand) DescribeImage() string {
	description := C.MagickDescribeImage(mw.mw)
	defer C.MagickRelinquishMemory(unsafe.Pointer(description))
	return C.GoString(description)
}

// Reduces the speckle noise in an image while perserving the edges of the
// original image.
func (mw *MagickWand) DespeckleImage() error {
	ok := C.MagickDespeckleImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Displays and image
func (mw *MagickWand) DisplayImage(server string) error {
	cstring := C.CString(server)
	defer C.free(unsafe.Pointer(cstring))
	ok := C.MagickDisplayImage(mw.mw, cstring)
	return mw.getLastErrorIfFailed(ok)
}

// Displays and image or image sequence
func (mw *MagickWand) DisplayImages(server string) error {
	cstring := C.CString(server)
	defer C.free(unsafe.Pointer(cstring))
	ok := C.MagickDisplayImages(mw.mw, cstring)
	return mw.getLastErrorIfFailed(ok)
}

// Draws vectors on the image as described by DrawingWand.
func (mw *MagickWand) DrawImage(dw *DrawingWand) error {
	ok := C.MagickDrawImage(mw.mw, dw.dw)
	return mw.getLastErrorIfFailed(ok)
}

// Enhance edges within the image with a convolution filter of the given
// radius. Use a radius of 0 and Edge() selects a suitable radius for you.
//
// radius: the radius of the pixel neighborhood
//
func (mw *MagickWand) EdgeImage(radius float64) error {
	ok := C.MagickEdgeImage(mw.mw, C.double(radius))
	return mw.getLastErrorIfFailed(ok)
}

// Returns a grayscale image with a three-dimensional effect. We convolve the
// image with a Gaussian operator of the given radius and standard deviation
// (sigma). For reasonable results, radius should be larger than sigma. Use a
// radius of 0 and Emboss() selects a suitable radius for you.
//
// radius: the radius of the Gaussian, in pixels, not counting the center pixel
//
// sigma: the standard deviation of the Gaussian, in pixels
//
func (mw *MagickWand) EmbossImage(radius, sigma float64) error {
	ok := C.MagickEmbossImage(mw.mw, C.double(radius), C.double(sigma))
	return mw.getLastErrorIfFailed(ok)
}

// Applies a digital filter that improves the quality of a noisy image
func (mw *MagickWand) EnhanceImage() error {
	ok := C.MagickEnhanceImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Equalizes the image histogram.
func (mw *MagickWand) EqualizeImage() error {
	ok := C.MagickEqualizeImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Extends the image as defined by the geometry, gravitt, and wand background
// color. Set the (x,y) offset of the geometry to move the original wand
// relative to the extended wand.
//
// width: the region width.
//
// height: the region height.
//
// x: the region x offset.
//
// y: the region y offset.
//
func (mw *MagickWand) ExtentImage(width, height uint, x, y int) error {
	ok := C.MagickExtentImage(mw.mw, C.size_t(width), C.size_t(height), C.ssize_t(x), C.ssize_t(y))
	return mw.getLastErrorIfFailed(ok)
}

// Creates a vertical mirror image by reflecting the pixels around the central
// x-axis.
func (mw *MagickWand) FlipImage() error {
	ok := C.MagickFlipImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Creates a horizontal mirror image by reflecting the pixels around the
// central y-axis.
func (mw *MagickWand) FlopImage() error {
	ok := C.MagickFlopImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Adds a simulated three-dimensional border around the image. The width and
// height specify the border width of the vertical and horizontal sides of the
// frame. The inner and outer bevels indicate the width of the inner and outer
// shadows of the frame.
//
// matteColor: the frame color pixel wand.
//
// width: the border width.
//
// height: the border height.
//
// innerBevel: the inner bevel width.
//
// outerBevel: the outer bevel width.
//
func (mw *MagickWand) FrameImage(matteColor *PixelWand, width, height uint, innerBevel, outerBevel int) error {
	ok := C.MagickFrameImage(mw.mw, matteColor.pw, C.ulong(width), C.ulong(height), C.long(innerBevel), C.long(outerBevel))
	return mw.getLastErrorIfFailed(ok)
}

// Evaluate expression for each pixel in the image.
func (mw *MagickWand) FxImage(expression string) (fxmw *MagickWand, err error) {
	csexpression := C.CString(expression)
	defer C.free(unsafe.Pointer(csexpression))
	fxmw = newMagickWand(C.MagickFxImage(mw.mw, csexpression))
	err = mw.GetLastError()
	return
}

// Evaluate expression for each pixel in the image's channel
func (mw *MagickWand) FxImageChannel(channel ChannelType, expression string) *MagickWand {
	csexpression := C.CString(expression)
	defer C.free(unsafe.Pointer(csexpression))

	return newMagickWand(C.MagickFxImageChannel(mw.mw, C.ChannelType(channel), csexpression))
}

// Gamma-corrects an image. The same image viewed on different devices will
// have perceptual differences in the way the image's intensities are
// represented on the screen. Specify individual gamma levels for the red,
// green, and blue channels, or adjust all three with the gamma parameter.
// Values typically range from 0.8 to 2.3. You can also reduce the influence
// of a particular channel with a gamma value of 0.
func (mw *MagickWand) GammaImage(gamma float64) error {
	ok := C.MagickGammaImage(mw.mw, C.double(gamma))
	return mw.getLastErrorIfFailed(ok)
}

// Gamma-corrects an image's channel. The same image viewed on different
// devices will have perceptual differences in the way the image's intensities
// are represented on the screen. Specify individual gamma levels for the red,
// green, and blue channels, or adjust all three with the gamma parameter.
// Values typically range from 0.8 to 2.3. You can also reduce the influence
// of a particular channel with a gamma value of 0.
func (mw *MagickWand) GammaImageChannel(channel ChannelType, gamma float64) error {
	ok := C.MagickGammaImageChannel(mw.mw, C.ChannelType(channel), C.double(gamma))
	return mw.getLastErrorIfFailed(ok)
}

// Returns the filename associated with an image sequence.
func (mw *MagickWand) GetFilename() string {
	p := C.MagickGetFilename(mw.mw)
	defer C.MagickRelinquishMemory(unsafe.Pointer(p))
	return C.GoString(p)
}

// Returns an image attribute as a string
func (mw *MagickWand) GetImageAttribute(attributeName string) string {
	n := C.CString(attributeName)
	defer C.free(unsafe.Pointer(n))
	p := C.MagickGetImageAttribute(mw.mw, n)
	defer C.MagickRelinquishMemory(unsafe.Pointer(p))
	return C.GoString(p)
}

// Returns the image orientation type
func (mw *MagickWand) GetImageOrientation() OrientationType {
	return OrientationType(C.MagickGetImageOrientation(mw.mw))
}

// Returns the image background color.
func (mw *MagickWand) GetImageBackgroundColor() (bgColor *PixelWand, err error) {
	cbgcolor := NewPixelWand()
	ok := C.MagickGetImageBackgroundColor(mw.mw, cbgcolor.pw)
	return cbgcolor, mw.getLastErrorIfFailed(ok)
}

// Returns the chromaticy blue primary point for the image.
//
// x: the chromaticity blue primary x-point.
//
// y: the chromaticity blue primary y-point.
//
func (mw *MagickWand) GetImageBluePrimary() (x, y float64, err error) {
	ok := C.MagickGetImageBluePrimary(mw.mw, (*C.double)(&x), (*C.double)(&y))
	err = mw.getLastErrorIfFailed(ok)
	return
}

// Returns the image border color.
func (mw *MagickWand) GetImageBorderColor() (borderColor *PixelWand, err error) {
	cbc := NewPixelWand()
	ok := C.MagickGetImageBorderColor(mw.mw, cbc.pw)
	return cbc, mw.getLastErrorIfFailed(ok)
}

// Gets the depth for one or more image channels.
func (mw *MagickWand) GetImageChannelDepth(channel ChannelType) uint {
	return uint(C.MagickGetImageChannelDepth(mw.mw, C.ChannelType(channel)))
}

// Gets the mean and standard deviation of one or more image channels.
func (mw *MagickWand) GetImageChannelMean(channel ChannelType) (mean, stdev float64, err error) {
	ok := C.MagickGetImageChannelMean(mw.mw, C.ChannelType(channel), (*C.double)(&mean), (*C.double)(&stdev))
	err = mw.getLastErrorIfFailed(ok)
	return
}

// Returns the color of the specified colormap index.
func (mw *MagickWand) GetImageColormapColor(index uint) (color *PixelWand, err error) {
	cpw := NewPixelWand()
	ok := C.MagickGetImageColormapColor(mw.mw, C.ulong(index), cpw.pw)
	return cpw, mw.getLastErrorIfFailed(ok)
}

// Gets the number of unique colors in the image.
func (mw *MagickWand) GetImageColors() uint {
	return uint(C.MagickGetImageColors(mw.mw))
}

// Gets the image colorspace.
func (mw *MagickWand) GetImageColorspace() ColorspaceType {
	return ColorspaceType(C.MagickGetImageColorspace(mw.mw))
}

// Returns the composite operator associated with the image.
func (mw *MagickWand) GetImageCompose() CompositeOperator {
	return CompositeOperator(C.MagickGetImageCompose(mw.mw))
}

// Gets the image compression.
func (mw *MagickWand) GetImageCompression() CompressionType {
	return CompressionType(C.MagickGetImageCompression(mw.mw))
}

// Gets the image delay.
func (mw *MagickWand) GetImageDelay() uint {
	return uint(C.MagickGetImageDelay(mw.mw))
}

// Gets the image depth.
func (mw *MagickWand) GetImageDepth() uint {
	return uint(C.MagickGetImageDepth(mw.mw))
}

// Gets the image disposal method.
func (mw *MagickWand) GetImageDispose() DisposeType {
	return DisposeType(C.MagickGetImageDispose(mw.mw))
}

// Returns the filename of a particular image in a sequence.
func (mw *MagickWand) GetImageFilename() string {
	p := C.MagickGetImageFilename(mw.mw)
	defer C.MagickRelinquishMemory(unsafe.Pointer(p))
	return C.GoString(p)
}

// Returns the format of a particular image in a sequence.
func (mw *MagickWand) GetImageFormat() string {
	p := C.MagickGetImageFormat(mw.mw)
	defer C.MagickRelinquishMemory(unsafe.Pointer(p))
	return C.GoString(p)
}

// Gets the image fuzz.
func (mw *MagickWand) GetImageFuzz() float64 {
	return float64(C.MagickGetImageFuzz(mw.mw))
}

// Gets the image gamma.
func (mw *MagickWand) GetImageGamma() float64 {
	return float64(C.MagickGetImageGamma(mw.mw))
}

// Gets the image geometry string.  NULL is
// returned if the image does not contain a geometry string.
func (mw *MagickWand) MagickGetImageGeometry() string {
	p := C.MagickGetImageGeometry(mw.mw)
	if p != nil {
		defer C.MagickRelinquishMemory(unsafe.Pointer(p))
		return C.GoString(p)
	}
	return ""
}

// Gets the image gravity.
func (mw *MagickWand) GetImageGravity() GravityType {
	return GravityType(C.MagickGetImageGravity(mw.mw))
}

// Gets the image at the current image index.
func (mw *MagickWand) GetImage() *MagickWand {
	return newMagickWand(C.MagickGetImage(mw.mw))
}

// Returns the chromaticy green primary point.
//
// x: the chromaticity green primary x-point.
//
// y: the chromaticity green primary y-point.
//
func (mw *MagickWand) GetImageGreenPrimary() (x, y float64, err error) {
	ok := C.MagickGetImageGreenPrimary(mw.mw, (*C.double)(&x), (*C.double)(&y))
	err = mw.getLastErrorIfFailed(ok)
	return
}

// Returns the image height.
func (mw *MagickWand) GetImageHeight() uint {
	return uint(C.MagickGetImageHeight(mw.mw))
}

// Returns the image histogram as an array of PixelWand wands.
//
// numberColors: the number of unique colors in the image and the number of
// pixel wands returned.
func (mw *MagickWand) GetImageHistogram() (numberColors uint, pws []PixelWand) {
	cnc := C.ulong(0)
	p := C.MagickGetImageHistogram(mw.mw, &cnc)
	defer C.MagickRelinquishMemory(unsafe.Pointer(p))
	q := uintptr(unsafe.Pointer(p))
	for {
		p = (**C.PixelWand)(unsafe.Pointer(q))
		if *p == nil {
			break
		}
		pws = append(pws, *newPixelWand(*p))
		q += unsafe.Sizeof(q)
	}
	numberColors = uint(cnc)
	return
}

// Gets the image interlace scheme.
func (mw *MagickWand) GetImageInterlaceScheme() InterlaceType {
	return InterlaceType(C.MagickGetImageInterlaceScheme(mw.mw))
}

// Gets the image iterations.
func (mw *MagickWand) GetImageIterations() uint {
	return uint(C.MagickGetImageIterations(mw.mw))
}

// Returns the image matte color.
func (mw *MagickWand) GetImageMatteColor() (matteColor *PixelWand, err error) {
	cptrpw := NewPixelWand()
	ok := C.MagickGetImageMatteColor(mw.mw, cptrpw.pw)
	return cptrpw, mw.getLastErrorIfFailed(ok)
}

// Returns the page geometry associated with the image.
//
// w, h: the page width and height
//
// x, h: the page x-offset and y-offset.
//
func (mw *MagickWand) GetImagePage() (w, h uint, x, y int, err error) {
	var cw, ch C.ulong
	var cx, cy C.long
	ok := C.MagickGetImagePage(mw.mw, &cw, &ch, &cx, &cy)
	return uint(cw), uint(ch), int(cx), int(cy), mw.getLastErrorIfFailed(ok)
}

// extracts pixel data from an image
//
// xOffset, yOffset: offset (from top left) on base canvas image on which to composite image data
//
// columns, rows: dimensions of image
//
// pixelMap: ordering of the pixel array
//
// storageType: define the data type of the pixels. Float and double types are expected to be normalized [0..1] otherwise [0..MaxRGB]
//
// pixels: contain the pixel components as defined by pixelMap and storageType
func (mw *MagickWand) GetImagePixels(xOffset, yOffset int, columns, rows uint, pixelMap string, storageType StorageType, pixels []byte) error {
	cspixelMap := C.CString(pixelMap)
	defer C.free(unsafe.Pointer(cspixelMap))
	ok := C.MagickGetImagePixels(mw.mw, C.long(xOffset), C.long(yOffset), C.ulong(columns), C.ulong(rows), cspixelMap, C.StorageType(storageType), (*C.uchar)(unsafe.Pointer(&pixels[0])))
	return mw.getLastErrorIfFailed(ok)
}


// Returns the chromaticy red primary point.
//
// x, y: the chromaticity red primary x/y-point.
//
func (mw *MagickWand) GetImageRedPrimary() (x, y float64, err error) {
	var cdx, cdy C.double
	ok := C.MagickGetImageRedPrimary(mw.mw, &cdx, &cdy)
	return float64(cdx), float64(cdy), mw.getLastErrorIfFailed(ok)
}

// Gets the image X and Y resolution.
func (mw *MagickWand) GetImageResolution() (x, y float64, err error) {
	var dx, dy C.double
	ok := C.MagickGetImageResolution(mw.mw, &dx, &dy)
	return float64(dx), float64(dy), mw.getLastErrorIfFailed(ok)
}

// Gets the image type that will be used when the
// image is saved. This may be different to the current image type, returned
// by GetImageType().
func (mw *MagickWand) GetImageSavedType() ImageType {
	return ImageType(C.MagickGetImageSavedType(mw.mw))
}

// Generates an SHA-256 message digest for the image pixel stream.
func (mw *MagickWand) GetImageSignature() string {
	p := C.MagickGetImageSignature(mw.mw)
	defer C.MagickRelinquishMemory(unsafe.Pointer(p))
	return C.GoString(p)
}

// Gets the potential image type
// To ensure the image type matches its potential, use SetImageType():
// wand.SetImageType(wand.GetImageType())
func (mw *MagickWand) GetImageType() ImageType {
	return ImageType(C.MagickGetImageType(mw.mw))
}

// Gets the image units of resolution.
func (mw *MagickWand) GetImageUnits() ResolutionType {
	return ResolutionType(C.MagickGetImageUnits(mw.mw))
}

// Returns the chromaticy white point.
//
// x, y: the chromaticity white x/y-point.
//
func (mw *MagickWand) GetImageWhitePoint() (x, y float64, err error) {
	ok := C.MagickGetImageWhitePoint(mw.mw, (*C.double)(&x), (*C.double)(&y))
	err = mw.getLastErrorIfFailed(ok)
	return
}

// Returns the image width.
func (mw *MagickWand) GetImageWidth() uint {
	return uint(C.MagickGetImageWidth(mw.mw))
}

// Returns the number of images associated with a magick wand.
func (mw *MagickWand) GetNumberImages() uint {
	return uint(C.MagickGetNumberImages(mw.mw))
}

// Gets the horizontal and vertical sampling factor.
func (mw *MagickWand) GetSamplingFactors() (factors []float64) {
	num := C.ulong(0)
	pd := C.MagickGetSamplingFactors(mw.mw, &num)
	defer C.MagickRelinquishMemory(unsafe.Pointer(pd))
	factors = sizedDoubleArrayToFloat64Slice(pd, num)
	return
}

// Returns the size associated with the magick wand.
func (mw *MagickWand) GetSize() (uint, uint, error) {
	w := C.ulong(0)
	h := C.ulong(0)
	ok := C.MagickGetSize(mw.mw, &w, &h)
	return uint(w), uint(h), mw.getLastErrorIfFailed(ok)
}

// Replaces colors in the image from a Hald color lookup table. A Hald color
// lookup table is a 3-dimensional color cube mapped to 2 dimensions. Create
// it with the HALD coder. You can apply any color transformation to the Hald
// image and then use this method to apply the transform to the image.
func (mw *MagickWand) HaldClutImage(hald *MagickWand) error {
	ok := C.MagickHaldClutImage(mw.mw, hald.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Returns true if the wand has more images when traversing the list in the
// forward direction
func (mw *MagickWand) HasNextImage() bool {
	return 1 == C.MagickHasNextImage(mw.mw)
}

// Returns true if the wand has more images when traversing the list in the
// reverse direction
func (mw *MagickWand) HasPreviousImage() bool {
	return 1 == C.MagickHasPreviousImage(mw.mw)
}

// Creates a new image that is a copy of an existing one with the image pixels
// "implode" by the specified percentage.
func (mw *MagickWand) ImplodeImage(radius float64) error {
	ok := C.MagickImplodeImage(mw.mw, C.double(radius))
	return mw.getLastErrorIfFailed(ok)
}

// Adds a label to your image.
func (mw *MagickWand) LabelImage(label string) error {
	cslabel := C.CString(label)
	defer C.free(unsafe.Pointer(cslabel))
	ok := C.MagickLabelImage(mw.mw, cslabel)
	return mw.getLastErrorIfFailed(ok)
}

// Adjusts the levels of an image by scaling the colors falling between
// specified white and black points to the full available quantum range. The
// parameters provided represent the black, mid, and white points. The black
// point specifies the darkest color in the image. Colors darker than the
// black point are set to zero. Mid point specifies a gamma correction to
// apply to the image. White point specifies the lightest color in the image.
// Colors brighter than the white point are set to the maximum quantum value.
func (mw *MagickWand) LevelImage(blackPoint, gamma, whitePoint float64) error {
	ok := C.MagickLevelImage(mw.mw, C.double(blackPoint), C.double(gamma), C.double(whitePoint))
	return mw.getLastErrorIfFailed(ok)
}

// Adjusts the levels of an image by scaling the colors falling between
// specified white and black points to the full available quantum range. The
// parameters provided represent the black, mid, and white points. The black
// point specifies the darkest color in the image. Colors darker than the
// black point are set to zero. Mid point specifies a gamma correction to
// apply to the image. White point specifies the lightest color in the image.
// Colors brighter than the white point are set to the maximum quantum value.
func (mw *MagickWand) LevelImageChannel(channel ChannelType, blackPoint, gamma, whitePoint float64) error {
	ok := C.MagickLevelImageChannel(mw.mw, C.ChannelType(channel), C.double(blackPoint), C.double(gamma), C.double(whitePoint))
	return mw.getLastErrorIfFailed(ok)
}

// This is a convenience method that scales an image proportionally to twice
// its original size.
func (mw *MagickWand) MagnifyImage() error {
	ok := C.MagickMagnifyImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Replaces the colors of an image with the closest color from a reference image.
//
// from: reference image
//
// dither: Set this integer value to something other than zero to dither the mapped image.
//
func (mw *MagickWand) MapImage(from *MagickWand, dither uint) error {
	ok := C.MagickMapImage(mw.mw, from.mw, C.uint(dither))
	return mw.getLastErrorIfFailed(ok)
}

// Changes the transparency value of any pixel that
// matches target and is an immediate neighbor.  If the method
// FillToBorderMethod is specified, the transparency value is changed for any
// neighbor pixel that does not match the bordercolor member of image.
func (mw *MagickWand) MatteFloodfillImage(opacity Quantum, fuzz float64, borderColor *PixelWand, x int, y int) error {
	ok := C.MagickMatteFloodfillImage(mw.mw, C.Quantum(opacity), C.double(fuzz), borderColor.pw, C.long(x), C.long(y))
	return mw.getLastErrorIfFailed(ok)
}

// Applies a digital filter that improves the quality
// of a noisy image.  Each pixel is replaced by the median in a set of
// neighboring pixels as defined by radius.
func (mw *MagickWand) MedianFilterImage(radius float64) error {
	ok := C.MagickMedianFilterImage(mw.mw, C.double(radius))
	return mw.getLastErrorIfFailed(ok)
}

// This is a convenience method that scales an image proportionally to
// one-half its original size
func (mw *MagickWand) MinifyImage() error {
	ok := C.MagickMinifyImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Lets you control the brightness, saturation, and hue of an image. Hue is
// the percentage of absolute rotation from the current position. For example
// 50 results in a counter-clockwise rotation of 90 degrees, 150 results in a
// clockwise rotation of 90 degrees, with 0 and 200 both resulting in a
// rotation of 180 degrees. To increase the color brightness by 20 and
// decrease the color saturation by 10 and leave the hue unchanged, use: 120,
// 90, 100.
//
// brightness: the percent change in brighness.
//
// saturation: the percent change in saturation.
//
// hue: the percent change in hue.
//
func (mw *MagickWand) ModulateImage(brightness, saturation, hue float64) error {
	ok := C.MagickModulateImage(mw.mw, C.double(brightness), C.double(saturation), C.double(hue))
	return mw.getLastErrorIfFailed(ok)
}

// Creates a composite image by combining several separate images. The images
// are tiled on the composite image with the name of the image optionally
// appearing just below the individual tile.
//
// dw: the drawing wand. The font name, size, and color are obtained from this
// wand.
//
// tileGeo: the number of tiles per row and page (e.g. 6x4+0+0).
//
// thumbGeo: Preferred image size and border size of each thumbnail (e.g.
// 120x120+4+3>).
//
// mode: Thumbnail framing mode: Frame, Unframe, or Concatenate.
//
// frame: Surround the image with an ornamental border (e.g. 15x15+3+3). The
// frame color is that of the thumbnail's matte color.
//
func (mw *MagickWand) MontageImage(dw *DrawingWand, tileGeo string, thumbGeo string, mode MontageMode, frame string) *MagickWand {
	cstile := C.CString(tileGeo)
	defer C.free(unsafe.Pointer(cstile))
	csthumb := C.CString(thumbGeo)
	defer C.free(unsafe.Pointer(csthumb))
	csframe := C.CString(frame)
	defer C.free(unsafe.Pointer(csframe))

	return newMagickWand(C.MagickMontageImage(mw.mw, dw.dw, cstile, csthumb, C.MontageMode(mode), csframe))
}

// Method morphs a set of images. Both the image pixels and size are linearly
// interpolated to give the appearance of a meta-morphosis from one image to
// the next.
//
// numFrames: the number of in-between images to generate.
func (mw *MagickWand) MorphImages(numFrames uint) *MagickWand {
	return newMagickWand(C.MagickMorphImages(mw.mw, C.ulong(numFrames)))
}

// inlays an image sequence to form a single coherent
// picture.  It returns a wand with each image in the sequence composited at
// the location defined by the page offset of the image.
func (mw *MagickWand) MosaicImages() *MagickWand {
	return newMagickWand(C.MagickMosaicImages(mw.mw))
}

// Simulates motion blur. We convolve the image with a Gaussian operator of
// the given radius and standard deviation (sigma). For reasonable results,
// radius should be larger than sigma. Use a radius of 0 and MotionBlurImage()
// selects a suitable radius for you. Angle gives the angle of the blurring
// motion.
//
// radius: the radius of the Gaussian, in pixels, not counting the center pixel.
//
// sigma: the standard deviation of the Gaussian, in pixels.
//
// angle: apply the effect along this angle.
//
func (mw *MagickWand) MotionBlurImage(radius, sigma, angle float64) error {
	ok := C.MagickMotionBlurImage(mw.mw, C.double(radius), C.double(sigma), C.double(angle))
	return mw.getLastErrorIfFailed(ok)
}

// Negates the colors in the reference image. The Grayscale option means that
// only grayscale values within the image are negated. You can also reduce the
// influence of a particular channel with a gamma value of 0.
//
// gray: If true, only negate grayscale pixels within the image.
//
func (mw *MagickWand) NegateImageChannel(channel ChannelType, gray bool) error {
	ok := C.MagickNegateImageChannel(mw.mw, C.ChannelType(channel), b2i(gray))
	return mw.getLastErrorIfFailed(ok)
}

// Negates the colors in the reference image. The Grayscale option means that
// only grayscale values within the image are negated. You can also reduce the
// influence of a particular channel with a gamma value of 0.
//
// gray: If true, only negate grayscale pixels within the image.
//
func (mw *MagickWand) NegateImage(gray bool) error {
	ok := C.MagickNegateImage(mw.mw, b2i(gray))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the next image in the wand as the current image. It is typically used
// after ResetIterator(), after which its first use will set the first image
// as the current image (unless the wand is empty). It will return false when
// no more images are left to be returned which happens when the wand is empty,
// or the current image is the last image. When the above condition (end of
// image list) is reached, the iterator is automaticall set so that you can
// start using PreviousImage() to again/ iterate over the images in the
// reverse direction, starting with the last image (again). You can jump to
// this condition immeditally using SetLastIterator().
func (mw *MagickWand) NextImage() bool {
	return 1 == C.MagickNextImage(mw.mw)
}

// Enhances the contrast of a color image by adjusting the pixels color to
// span the entire range of colors available. You can also reduce the
// influence of a particular channel with a gamma value of 0.
func (mw *MagickWand) NormalizeImage() error {
	ok := C.MagickNormalizeImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Applies a special effect filter that simulates an oil painting. Each pixel
// is replaced by the most frequent color occurring in a circular region
// defined by radius.
//
// radius: the radius of the circular neighborhood.
//
func (mw *MagickWand) OilPaintImage(radius float64) error {
	ok := C.MagickOilPaintImage(mw.mw, C.double(radius))
	return mw.getLastErrorIfFailed(ok)
}

//
func (mw *MagickWand) OpaqueImage(target, fill *PixelWand, fuzz float64) error {
	ok := C.MagickOpaqueImage(mw.mw, target.pw, fill.pw, C.double(fuzz))
	return mw.getLastErrorIfFailed(ok)
}

// This is like ReadImage() except the only valid information returned is the
// image width, height, size, and format. It is designed to efficiently obtain
// this information from a file without reading the entire image sequence into
// memory.
func (mw *MagickWand) PingImage(filename string) error {
	csfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(csfilename))
	ok := C.MagickPingImage(mw.mw, csfilename)
	return mw.getLastErrorIfFailed(ok)
}

// Tiles 9 thumbnails of the specified image with an image processing
// operation applied at varying strengths. This helpful to quickly pin-point
// an appropriate parameter for an image processing operation.
func (mw *MagickWand) PreviewImages(preview PreviewType) *MagickWand {
	return newMagickWand(C.MagickPreviewImages(mw.mw, C.PreviewType(preview)))
}

// Sets the previous image in the wand as the current image. It is typically
// used after SetLastIterator(), after which its first use will set the last
// image as the current image (unless the wand is empty). It will return false
// when no more images are left to be returned which happens when the wand is
// empty, or the current image is the first image. At that point the iterator
// is than reset to again process images in the forward direction, again
// starting with the first image in list. Images added at this point are
// prepended. Also at that point any images added to the wand using AddImages()
// or ReadImages() will be prepended before the first image. In this sense the
// condition is not quite exactly the same as ResetIterator().
func (mw *MagickWand) PreviousImage() bool {
	return 1 == C.MagickPreviousImage(mw.mw)
}

// Analyzes the colors within a reference image and chooses a fixed number of
// colors to represent the image. The goal of the algorithm is to minimize the
// color difference between the input and output image while minimizing the
// processing time.
//
// numColors: the number of colors.
//
// colorspace: Perform color reduction in this colorspace, typically
// RGBColorspace.
//
// treedepth: Normally, this integer value is zero or one. A zero or one tells
// Quantize to choose a optimal tree depth of Log4(number_colors). A tree of
// this depth generally allows the best representation of the reference image
// with the least amount of memory and the fastest computational speed. In
// some cases, such as an image with low color dispersion (a few number of
// colors), a value other than Log4(number_colors) is required. To expand the
// color tree completely, use a value of 8.
//
// dither: A value other than zero distributes the difference between an
// original image and the corresponding color reduced image to neighboring
// pixels along a Hilbert curve.
//
// measureError: A value other than zero measures the difference between the
// original and quantized images. This difference is the total quantization
// error. The error is computed by summing over all pixels in an image the
// distance squared in RGB space between each reference pixel value and its
// quantized value.
//
func (mw *MagickWand) QuantizeImage(numColors uint, colorspace ColorspaceType, treedepth uint, dither bool, measureError bool) error {
	ok := C.MagickQuantizeImage(mw.mw, C.ulong(numColors), C.ColorspaceType(colorspace), C.ulong(treedepth), b2i(dither), b2i(measureError))
	return mw.getLastErrorIfFailed(ok)
}

// Analyzes the colors within a sequence of images and chooses a fixed number
// of colors to represent the image. The goal of the algorithm is to minimize
// the color difference between the input and output image while minimizing the
// processing time.
//
// numColors: the number of colors.
//
// colorspace: Perform color reduction in this colorspace, typically
// RGBColorspace.
//
// treedepth: Normally, this integer value is zero or one. A zero or one tells
// Quantize to choose a optimal tree depth of Log4(number_colors). A tree of
// this depth generally allows the best representation of the reference image
// with the least amount of memory and the fastest computational speed. In
// some cases, such as an image with low color dispersion (a few number of
// colors), a value other than Log4(number_colors) is required. To expand the
// color tree completely, use a value of 8.
//
// dither: A value other than zero distributes the difference between an
// original image and the corresponding color reduced image to neighboring
// pixels along a Hilbert curve.
//
// measureError: A value other than zero measures the difference between the
// original and quantized images. This difference is the total quantization
// error. The error is computed by summing over all pixels in an image the
// distance squared in RGB space between each reference pixel value and its
// quantized value.
//
func (mw *MagickWand) QuantizeImages(numColors uint, colorspace ColorspaceType, treedepth uint, dither bool, measureError bool) error {
	ok := C.MagickQuantizeImages(mw.mw, C.ulong(numColors), C.ColorspaceType(colorspace), C.ulong(treedepth), b2i(dither), b2i(measureError))
	return mw.getLastErrorIfFailed(ok)
}

// Returns a FontMetrics struct
func (mw *MagickWand) QueryFontMetrics(dw *DrawingWand, textLine string) *FontMetrics {
	cstext := C.CString(textLine)
	defer C.free(unsafe.Pointer(cstext))
	cdoubles := C.MagickQueryFontMetrics(mw.mw, dw.dw, cstext)
	defer C.MagickRelinquishMemory(unsafe.Pointer(cdoubles))
	doubles := sizedDoubleArrayToFloat64Slice(cdoubles, 7)
	return NewFontMetricsFromArray(doubles)
}

// Returns any font that match the specified pattern (e.g. "*" for all)
func (mw *MagickWand) QueryFonts(pattern string) (fonts []string) {
	cspattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(cspattern))
	var num C.ulong
	copts := C.MagickQueryFonts(cspattern, &num)
	defer C.MagickRelinquishMemory(unsafe.Pointer(copts))
	fonts = sizedCStringArrayToStringSlice(copts, num)
	return
}

// Returns any supported image format that match the specified pattern (e.g. "*" for all)
func (mw *MagickWand) QueryFormats(pattern string) (formats []string) {
	cspattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(cspattern))
	var num C.ulong
	copts := C.MagickQueryFormats(cspattern, &num)
	defer C.MagickRelinquishMemory(unsafe.Pointer(copts))
	formats = sizedCStringArrayToStringSlice(copts, num)
	return
}

// Radial blurs an image.
func (mw *MagickWand) RadialBlurImage(angle float64) error {
	ok := C.MagickRadialBlurImage(mw.mw, C.double(angle))
	return mw.getLastErrorIfFailed(ok)
}

// Creates a simulated three-dimensional button-like effect by lightening and
// darkening the edges of the image. Members width and height of raise_info
// define the width of the vertical and horizontal edge of the effect. width,
//
// height, x, y: Define the dimensions of the area to raise.
//
// raise: A value other than zero creates a 3-D raise effect, otherwise it has
// a lowered effect.
//
func (mw *MagickWand) RaiseImage(width uint, height uint, x int, y int, raise bool) error {
	ok := C.MagickRaiseImage(mw.mw, C.ulong(width), C.ulong(height), C.long(x), C.long(y), b2i(raise))
	return mw.getLastErrorIfFailed(ok)
}

// Reads an image or image sequence. The images are inserted at the current
// image pointer position. Use SetFirstIterator(), SetLastIterator, or
// SetImageIndex() to specify the current image pointer position at the
// beginning of the image list, the end, or anywhere in-between respectively.
func (mw *MagickWand) ReadImage(filename string) error {
	csfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(csfilename))
	ok := C.MagickReadImage(mw.mw, csfilename)
	return mw.getLastErrorIfFailed(ok)
}

// Reads an image or image sequence from a blob.
func (mw *MagickWand) ReadImageBlob(blob []byte) error {
	if len(blob) == 0 {
		return errors.New("zero-length blob not permitted")
	}
	ok := C.MagickReadImageBlob(mw.mw, (*C.uchar)(unsafe.Pointer(&blob[0])), C.size_t(len(blob)))
	return mw.getLastErrorIfFailed(ok)
}

// Reads an image or image sequence from an open file descriptor.
func (mw *MagickWand) ReadImageFile(img *os.File) error {
	file, err := cfdopen(img, "rb")
	if err != nil {
		return err
	}
	defer C.fclose(file)
	ok := C.MagickReadImageFile(mw.mw, file)
	return mw.getLastErrorIfFailed(ok)
}

// Smooths the contours o an image while still
// preserving edge information. The algorithm works by replacing each pixel
// with its neighbor closes in value. A neighbor is defined by radius. Use
// a radius of 0 and the function selects a suitable radius for you.
func (mw *MagickWand) ReduceNoiseImage(radius float64) error {
	ok := C.MagickReduceNoiseImage(mw.mw, C.double(radius))
	return mw.getLastErrorIfFailed(ok)
}

// Removes an image from the image list.
func (mw *MagickWand) RemoveImage() error {
	ok := C.MagickRemoveImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Removes the named image profile and returns it.
//
// name: name of profile to return: ICC, IPTC, or generic profile.
//
func (mw *MagickWand) RemoveImageProfile(name string) []byte {
	csname := C.CString(name)
	defer C.free(unsafe.Pointer(csname))
	clen := C.ulong(0)
	profile := C.MagickRemoveImageProfile(mw.mw, csname, &clen)
	defer C.MagickRelinquishMemory(unsafe.Pointer(profile))
	return C.GoBytes(unsafe.Pointer(profile), C.int(clen))
}

// Resample image to desired resolution.
//
// xRes/yRes: the new image x/y resolution.
//
// filter: Image filter to use.
//
// blur: the blur factor where > 1 is blurry, < 1 is sharp.
//
func (mw *MagickWand) ResampleImage(xRes, yRes float64, filter FilterType, blur float64) error {
	ok := C.MagickResampleImage(mw.mw, C.double(xRes), C.double(yRes), C.FilterTypes(filter), C.double(blur))
	return mw.getLastErrorIfFailed(ok)
}

// This method resets the wand iterator.
// It is typically used either before iterating though images, or before calling specific methods such as AppendImages()
// to append all images together.
// Afterward you can use NextImage() to iterate over all the images in a wand container, starting with the first image.
// Using this before AddImages() or ReadImages() will cause new images to be inserted between the first and second image.
func (mw *MagickWand) ResetIterator() {
	C.MagickResetIterator(mw.mw)
}

// Scales an image to the desired dimensions
//
// cols: the number of cols in the scaled image.
//
// rows: the number of rows in the scaled image.
//
// filter: Image filter to use.
//
// blur: the blur factor where > 1 is blurry, < 1 is sharp.
//
func (mw *MagickWand) ResizeImage(cols, rows uint, filter FilterType, blur float64) error {
	ok := C.MagickResizeImage(mw.mw, C.ulong(cols), C.ulong(rows), C.FilterTypes(filter), C.double(blur))
	return mw.getLastErrorIfFailed(ok)
}

// Offsets an image as defined by x and y.
//
// x: the x offset.
//
// y: the y offset.
//
func (mw *MagickWand) RollImage(x, y int) error {
	ok := C.MagickRollImage(mw.mw, C.long(x), C.long(y))
	return mw.getLastErrorIfFailed(ok)
}

// Rotates an image the specified number of degrees. Empty triangles left over
// from rotating the image are filled with the background color.
//
// background: the background pixel wand.
//
// degrees: the number of degrees to rotate the image.
//
func (mw *MagickWand) RotateImage(background *PixelWand, degrees float64) error {
	ok := C.MagickRotateImage(mw.mw, background.pw, C.double(degrees))
	return mw.getLastErrorIfFailed(ok)
}

// Scales an image to the desired dimensions with pixel sampling. Unlike other
// scaling methods, this method does not introduce any additional color into
// the scaled image.
func (mw *MagickWand) SampleImage(cols, rows uint) error {
	ok := C.MagickSampleImage(mw.mw, C.ulong(cols), C.ulong(rows))
	return mw.getLastErrorIfFailed(ok)
}

// Scales the size of an image to the given dimensions.
func (mw *MagickWand) ScaleImage(cols, rows uint) error {
	ok := C.MagickScaleImage(mw.mw, C.ulong(cols), C.ulong(rows))
	return mw.getLastErrorIfFailed(ok)
}

// Separates a channel from the image and returns a grayscale image. A channel
// is a particular color component of each pixel in the image.
func (mw *MagickWand) SeparateImageChannel(channel ChannelType) error {
	ok := C.MagickSeparateImageChannel(mw.mw, C.ChannelType(channel))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the wand compression quality.
func (mw *MagickWand) SetCompressionQuality(quality uint) error {
	ok := C.MagickSetCompressionQuality(mw.mw, C.ulong(quality))
	return mw.getLastErrorIfFailed(ok)
}

// Replaces the last image returned by SetImageIndex(), NextImage(),
// PreviousImage() with the images from the specified wand.
func (mw *MagickWand) SetImage(source *MagickWand) error {
	ok := C.MagickSetImage(mw.mw, source.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Sets the format of the magick wand.
func (mw *MagickWand) SetFormat(format string) error {
	csformat := C.CString(format)
	defer C.free(unsafe.Pointer(csformat))
	ok := C.MagickSetFormat(mw.mw, csformat)
	return mw.getLastErrorIfFailed(ok)
}

// Sets an image attribute
func (mw *MagickWand) SetImageAttribute(name, value string) error {
	n := C.CString(name)
	v := C.CString(value)
	defer C.free(unsafe.Pointer(n))
	defer C.free(unsafe.Pointer(v))
	ok := C.MagickSetImageAttribute(mw.mw, n, v)
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image background color.
func (mw *MagickWand) SetImageBackgroundColor(background *PixelWand) error {
	ok := C.MagickSetImageBackgroundColor(mw.mw, background.pw)
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image chromaticity blue primary point.
func (mw *MagickWand) SetImageBluePrimary(x, y float64) error {
	ok := C.MagickSetImageBluePrimary(mw.mw, C.double(x), C.double(y))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image border color.
func (mw *MagickWand) SetImageBorderColor(border *PixelWand) error {
	ok := C.MagickSetImageBorderColor(mw.mw, border.pw)
	return mw.getLastErrorIfFailed(ok)
}

// Sets the depth of a particular image channel.
//
// depth: the image depth in bits.
//
func (mw *MagickWand) SetImageChannelDepth(channel ChannelType, depth uint) error {
	ok := C.MagickSetImageChannelDepth(mw.mw, C.ChannelType(channel), C.ulong(depth))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the color of the specified colormap index.
//
// index: the offset into the image colormap.
//
// color: return the colormap color in this wand.
//
func (mw *MagickWand) SetImageColormapColor(index uint, color *PixelWand) error {
	ok := C.MagickSetImageColormapColor(mw.mw, C.ulong(index), color.pw)
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image colorspace.
func (mw *MagickWand) SetImageColorspace(colorspace ColorspaceType) error {
	ok := C.MagickSetImageColorspace(mw.mw, C.ColorspaceType(colorspace))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image composite operator, useful for specifying how to composite
/// the image thumbnail when using the MontageImage() method.
func (mw *MagickWand) SetImageCompose(compose CompositeOperator) error {
	ok := C.MagickSetImageCompose(mw.mw, C.CompositeOperator(compose))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image compression.
func (mw *MagickWand) SetImageCompression(compression CompressionType) error {
	ok := C.MagickSetImageCompression(mw.mw, C.CompressionType(compression))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image delay.
//
// delay: the image delay in ticks-per-second units.
//
func (mw *MagickWand) SetImageDelay(delay uint) error {
	ok := C.MagickSetImageDelay(mw.mw, C.ulong(delay))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image depth.
//
// depth: the image depth in bits: 8, 16, or 32.
//
func (mw *MagickWand) SetImageDepth(depth uint) error {
	ok := C.MagickSetImageDepth(mw.mw, C.ulong(depth))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image disposal method.
func (mw *MagickWand) SetImageDispose(dispose DisposeType) error {
	ok := C.MagickSetImageDispose(mw.mw, C.DisposeType(dispose))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the filename of a particular image in a sequence.
func (mw *MagickWand) SetImageFilename(filename string) error {
	csfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(csfilename))
	ok := C.MagickSetImageFilename(mw.mw, csfilename)
	return mw.getLastErrorIfFailed(ok)
}

// Sets the format of a particular image in a sequence.
//
// format: the image format.
//
func (mw *MagickWand) SetImageFormat(format string) error {
	csformat := C.CString(format)
	defer C.free(unsafe.Pointer(csformat))
	ok := C.MagickSetImageFormat(mw.mw, csformat)
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image fuzz.
func (mw *MagickWand) SetImageFuzz(fuzz float64) error {
	ok := C.MagickSetImageFuzz(mw.mw, C.double(fuzz))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image gamma.
func (mw *MagickWand) SetImageGamma(gamma float64) error {
	ok := C.MagickSetImageGamma(mw.mw, C.double(gamma))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image gravity type.
func (mw *MagickWand) SetImageGravity(gravity GravityType) error {
	ok := C.MagickSetImageGravity(mw.mw, C.GravityType(gravity))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image chromaticity green primary point.
func (mw *MagickWand) SetImageGreenPrimary(x, y float64) error {
	ok := C.MagickSetImageGreenPrimary(mw.mw, C.double(x), C.double(y))
	return mw.getLastErrorIfFailed(ok)
}

// This method set the iterator to the given position in the image list specified with the index parameter.
// A zero index will set the first image as current, and so on. Negative indexes can be used to specify an
// image relative to the end of the images in the wand, with -1 being the last image in the wand.
// If the index is invalid (range too large for number of images in wand) the function will return false.
// In that case the current image will not change.
// After using any images added to the wand using AddImage() or ReadImage() will be added after the image indexed,
// regardless of if a zero (first image in list) or negative index (from end) is used.
// Jumping to index 0 is similar to ResetIterator() but differs in how NextImage() behaves afterward.
func (mw *MagickWand) SetImageIndex(index int) bool {
	return 1 == C.int(C.MagickSetImageIndex(mw.mw, C.long(index)))
}

// Sets the image interlace scheme.
func (mw *MagickWand) SetImageInterlaceScheme(interlace InterlaceType) error {
	ok := C.MagickSetImageInterlaceScheme(mw.mw, C.InterlaceType(interlace))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image iterations.
func (mw *MagickWand) SetImageIterations(iterations uint) error {
	ok := C.MagickSetImageIterations(mw.mw, C.ulong(iterations))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image matte channel.
func (mw *MagickWand) SetImageMatte(matte bool) error {
	ok := C.MagickSetImageMatte(mw.mw, b2i(matte))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image matte color.
func (mw *MagickWand) SetImageMatteColor(matte *PixelWand) error {
	ok := C.MagickSetImageMatteColor(mw.mw, matte.pw)
	return mw.getLastErrorIfFailed(ok)
}

// Associates one or options with a particular image format.
// (.e.g SetImageOption("jpeg","preserve-settings","true").
func (mw *MagickWand) SetImageOption(format, key, value string) error {
	csformat := C.CString(format)
	defer C.free(unsafe.Pointer(csformat))
	cskey := C.CString(key)
	defer C.free(unsafe.Pointer(cskey))
	csvalue := C.CString(value)
	defer C.free(unsafe.Pointer(csvalue))
	ok := C.MagickSetImageOption(mw.mw, csformat, cskey, csvalue)
	return mw.getLastErrorIfFailed(ok)
}

// Sets the page geometry of the image.
func (mw *MagickWand) SetImagePage(width, height uint, x, y int) error {
	ok := C.MagickSetImagePage(mw.mw, C.ulong(width), C.ulong(height), C.long(x), C.long(y))
	return mw.getLastErrorIfFailed(ok)
}

// Sets pixel data in the image at the location you specify
// (.e.g SetImagePixels(0,0,0,640,1,"RGB",CharPixel,pixels));
//
// xOffset, yOffset: offset (from top left) on base canvas image on which to composite image data
//
// columns, rows: dimensions of image
//
// pixelMap: ordering of the pixel array
//
// storageType: define the data type of the pixels. Float and double types are expected to be normalized [0..1] otherwise [0..MaxRGB]
//
// pixels: contain the pixel components as defined by pixelMap and storageType
func (mw *MagickWand) SetImagePixels(xOffset, yOffset int, columns, rows uint, pixelMap string, storageType StorageType, pixels []byte) error {
	cspixelMap := C.CString(pixelMap)
	defer C.free(unsafe.Pointer(cspixelMap))
	ok := C.MagickSetImagePixels(mw.mw, C.long(xOffset), C.long(yOffset), C.ulong(columns), C.ulong(rows), cspixelMap, C.StorageType(storageType), (*C.uchar)(unsafe.Pointer(&pixels[0])))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image chromaticity red primary point.
func (mw *MagickWand) SetImageRedPrimary(x, y float64) error {
	ok := C.MagickSetImageRedPrimary(mw.mw, C.double(x), C.double(y))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image rendering intent.
func (mw *MagickWand) SetImageRenderingIntent(ri RenderingIntent) error {
	ok := C.MagickSetImageRenderingIntent(mw.mw, C.RenderingIntent(ri))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image resolution.
func (mw *MagickWand) SetImageResolution(xRes, yRes float64) error {
	ok := C.MagickSetImageResolution(mw.mw, C.double(xRes), C.double(yRes))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image type that will be used when the image is saved.
func (mw *MagickWand) SetImageSavedType(it ImageType) error {
	ok := C.MagickSetImageSavedType(mw.mw, C.ImageType(it))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image scene.
func (mw *MagickWand) SetImageScene(scene uint) error {
	ok := C.MagickSetImageScene(mw.mw, C.ulong(scene))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image type.
func (mw *MagickWand) SetImageType(imgtype ImageType) error {
	ok := C.MagickSetImageType(mw.mw, C.ImageType(imgtype))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image units of resolution.
func (mw *MagickWand) SetImageUnits(units ResolutionType) error {
	ok := C.MagickSetImageUnits(mw.mw, C.ResolutionType(units))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image virtual pixel method.
func (mw *MagickWand) SetImageVirtualPixelMethod(method VirtualPixelMethod) VirtualPixelMethod {
	return VirtualPixelMethod(C.MagickSetImageVirtualPixelMethod(mw.mw, C.VirtualPixelMethod(method)))
}

// Sets the image chromaticity white point.
func (mw *MagickWand) SetImageWhitePoint(x, y float64) error {
	ok := C.MagickSetImageWhitePoint(mw.mw, C.double(x), C.double(y))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image interlacing scheme
func (mw *MagickWand) SetInterlaceScheme(scheme InterlaceType) error {
	ok := C.MagickSetInterlaceScheme(mw.mw, C.InterlaceType(scheme))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the passphrase.
func (mw *MagickWand) SetPassphrase(passphrase string) error {
	cspassphrase := C.CString(passphrase)
	defer C.free(unsafe.Pointer(cspassphrase))
	ok := C.MagickSetPassphrase(mw.mw, cspassphrase)
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image resolution.
func (mw *MagickWand) SetResolution(xRes, yRes float64) error {
	ok := C.MagickSetResolution(mw.mw, C.double(xRes), C.double(yRes))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the resolution units of the magick wand.
// It should be used in conjunction with SetResolution().
// This method works both before and after an image has been read.
func (mw *MagickWand) SetResolutionUnits(units ResolutionType) error {
	ok := C.MagickSetResolutionUnits(mw.mw, C.ResolutionType(units))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the limit for a particular resource in megabytes.
func (mw *MagickWand) SetResourceLimit(rtype ResourceType, limit uint) error {
	ok := C.MagickSetResourceLimit(C.ResourceType(rtype), C.ulong(limit))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image sampling factors.
//
// samplingFactors: An array of floats representing the sampling factor for
// each color component (in RGB order).
func (mw *MagickWand) SetSamplingFactors(samplingFactors []float64) error {
	ok := C.MagickSetSamplingFactors(mw.mw, C.ulong(len(samplingFactors)), (*C.double)(&samplingFactors[0]))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the size of the magick wand.  Set it before you
// read a raw image format such as RGB, GRAY, or CMYK.
func (mw *MagickWand) SetSize(width, heigh uint) error {
	ok := C.MagickSetSize(mw.mw, C.ulong(width), C.ulong(heigh))
	return mw.getLastErrorIfFailed(ok)
}

// Sharpens an image. We convolve the image with a Gaussian operator of the
// given radius and standard deviation (sigma). For reasonable results, the
// radius should be larger than sigma. Use a radius of 0 and SharpenImage()
// selects a suitable radius for you.
//
// radius: the radius of the Gaussian, in pixels, not counting the center pixel.
//
// sigma: the standard deviation of the Gaussian, in pixels.
//
func (mw *MagickWand) SharpenImage(radius, sigma float64) error {
	ok := C.MagickSharpenImage(mw.mw, C.double(radius), C.double(sigma))
	return mw.getLastErrorIfFailed(ok)
}

// Shaves pixels from the image edges. It allocates the memory necessary for
// the new Image structure and returns a pointer to the new image.
func (mw *MagickWand) ShaveImage(cols, rows uint) error {
	ok := C.MagickShaveImage(mw.mw, C.ulong(cols), C.ulong(rows))
	return mw.getLastErrorIfFailed(ok)
}

// Slides one edge of an image along the X or Y axis, creating a parallelogram.
// An X direction shear slides an edge along the X axis, while a Y direction
// shear slides an edge along the Y axis. The amount of the shear is controlled
// by a shear angle. For X direction shears, xShear is measured relative to the
// Y axis, and similarly, for Y direction shears yShear is measured relative to
// the X axis. Empty triangles left over from shearing the image are filled
// with the background color.
func (mw *MagickWand) ShearImage(background *PixelWand, xShear, yShear float64) error {
	ok := C.MagickShearImage(mw.mw, background.pw, C.double(xShear), C.double(yShear))
	return mw.getLastErrorIfFailed(ok)
}

// Applies a special effect to the image, similar to the effect achieved in a
// photo darkroom by selectively exposing areas of photo sensitive paper to
// light. Threshold ranges from 0 to QuantumRange and is a measure of the
// extent of the solarization.
//
// threshold: define the extent of the solarization.
//
func (mw *MagickWand) SolarizeImage(threshold float64) error {
	ok := C.MagickSolarizeImage(mw.mw, C.double(threshold))
	return mw.getLastErrorIfFailed(ok)
}

// Is a special effects method that randomly displaces each pixel in a block
// defined by the radius parameter.
//
// radius: Choose a random pixel in a neighborhood of this extent.
//
func (mw *MagickWand) SpreadImage(radius float64) error {
	ok := C.MagickSpreadImage(mw.mw, C.double(radius))
	return mw.getLastErrorIfFailed(ok)
}

// Hides a digital watermark within the image. Recover the hidden watermark
// later to prove that the authenticity of an image. Offset defines the start
// position within the image to hide the watermark.
//
// offset: start hiding at this offset into the image.
//
func (mw *MagickWand) SteganoImage(watermark *MagickWand, offset int) *MagickWand {
	return newMagickWand(C.MagickSteganoImage(mw.mw, watermark.mw, C.long(offset)))
}

// Composites two images and produces a single image that is the composite of
// a left and right image of a stereo pair.
func (mw *MagickWand) StereoImage(offset *MagickWand) *MagickWand {
	return newMagickWand(C.MagickStereoImage(mw.mw, offset.mw))
}

// Strips an image of all profiles and comments.
func (mw *MagickWand) StripImage() error {
	ok := C.MagickStripImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Swirls the pixels about the center of the image, where degrees indicates the
// sweep of the arc through which each pixel is moved. You get a more dramatic
// effect as the degrees move from 1 to 360.
//
// degrees: define the tightness of the swirling effect.
//
func (mw *MagickWand) SwirlImage(degrees float64) error {
	ok := C.MagickSwirlImage(mw.mw, C.double(degrees))
	return mw.getLastErrorIfFailed(ok)
}

// Repeatedly tiles the texture image across and down the image canvas.
func (mw *MagickWand) TextureImage(texture *MagickWand) *MagickWand {
	return newMagickWand(C.MagickTextureImage(mw.mw, texture.mw))
}

// Changes the value of individual pixels based on the intensity of each pixel
// compared to threshold. The result is a high-contrast, two color image.
//
// threshold: define the threshold value.
//
func (mw *MagickWand) ThresholdImage(threshold float64) error {
	ok := C.MagickThresholdImage(mw.mw, C.double(threshold))
	return mw.getLastErrorIfFailed(ok)
}

// Changes the value of individual pixels based on the intensity of each pixel
// compared to threshold. The result is a high-contrast, two color image.
//
// threshold: define the threshold value.
//
func (mw *MagickWand) ThresholdImageChannel(channel ChannelType, threshold float64) error {
	ok := C.MagickThresholdImageChannel(mw.mw, C.ChannelType(channel), C.double(threshold))
	return mw.getLastErrorIfFailed(ok)
}

// Applies a color vector to each pixel in the image. The length of the vector
// is 0 for black and white and at its maximum for the midtones. The vector
// weighting function is f(x)=(1-(4.0*((x-0.5)*(x-0.5)))).
//
// tint: the tint pixel wand.
//
// opacity: the opacity pixel wand.
//
func (mw *MagickWand) TintImage(tint, opacity *PixelWand) error {
	ok := C.MagickTintImage(mw.mw, tint.pw, opacity.pw)
	return mw.getLastErrorIfFailed(ok)
}

// Is a convenience method that behaves like ResizeImage() or CropImage() but
// accepts scaling and/or cropping information as a region geometry
// specification. If the operation fails, a NULL image handle is returned.
// crop: a crop geometry string. This geometry defines a subregion of the
// image to crop.
// geometry: an image geometry string. This geometry defines the final size
// of the image.
func (mw *MagickWand) TransformImage(crop string, geometry string) *MagickWand {
	cscrop, csgeo := C.CString(crop), C.CString(geometry)
	defer C.free(unsafe.Pointer(cscrop))
	defer C.free(unsafe.Pointer(csgeo))

	return newMagickWand(C.MagickTransformImage(mw.mw, cscrop, csgeo))
}

// Changes any pixel that matches color with the color
// defined by fill.
func (mw *MagickWand) TransparentImage(target *PixelWand, opacity C.Quantum, fuzz float64) error {
	ok := C.MagickTransparentImage(mw.mw, target.pw, opacity, C.double(fuzz))
	return mw.getLastErrorIfFailed(ok)
}

// Remove edges that are the background color from the image.
//
// fuzz: by default target must match a particular pixel color exactly.
// However, in many cases two colors may differ by a small amount. The fuzz
// member of image defines how much tolerance is acceptable to consider two
// colors as the same. For example, set fuzz to 10 and the color red at
// intensities of 100 and 102 respectively are now interpreted as the same
// color for the purposes of the floodfill.
func (mw *MagickWand) TrimImage(fuzz float64) error {
	ok := C.MagickTrimImage(mw.mw, C.double(fuzz))
	return mw.getLastErrorIfFailed(ok)
}

// Unsharpens an image. We convolve the image with a Gaussian operator of the
// given radius and standard deviation (sigma). For reasonable results, radius
// should be larger than sigma. Use a radius of 0 and UnsharpMaskImage()
// selects a suitable radius for you.
//
// radius: the radius of the Gaussian, in pixels, not counting the center pixel.
//
// sigma: the standard deviation of the Gaussian, in pixels.
//
// amount: the percentage of the difference between the original and the blur
// image that is added back into the original.
//
// threshold: the threshold in pixels needed to apply the diffence amount.
//
func (mw *MagickWand) UnsharpMaskImage(radius, sigma, amount, threshold float64) error {
	ok := C.MagickUnsharpMaskImage(mw.mw, C.double(radius), C.double(sigma), C.double(amount), C.double(threshold))
	return mw.getLastErrorIfFailed(ok)
}

// Creates a "ripple" effect in the image by shifting the pixels vertically
// along a sine wave whose amplitude and wavelength is specified by the given
// parameters.
//
// amplitude, wavelength: Define the amplitude and wave length of the sine wave.
//
func (mw *MagickWand) WaveImage(amplitude, wavelength float64) error {
	ok := C.MagickWaveImage(mw.mw, C.double(amplitude), C.double(wavelength))
	return mw.getLastErrorIfFailed(ok)
}

// Is like ThresholdImage() but force all pixels above the threshold into white
// while leaving all pixels below the threshold unchanged.
func (mw *MagickWand) WhiteThresholdImage(threshold *PixelWand) error {
	ok := C.MagickWhiteThresholdImage(mw.mw, threshold.pw)
	return mw.getLastErrorIfFailed(ok)
}

// Writes an image to the specified filename.
func (mw *MagickWand) WriteImage(filename string) error {
	csfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(csfilename))
	ok := C.MagickWriteImage(mw.mw, csfilename)
	return mw.getLastErrorIfFailed(ok)
}

// Writes an image sequence to an open file descriptor.
func (mw *MagickWand) WriteImageFile(out *os.File) error {
	file, err := cfdopen(out, "w")
	if err != nil {
		return err
	}
	defer C.fclose(file)
	ok := C.MagickWriteImageFile(mw.mw, file)
	return mw.getLastErrorIfFailed(ok)
}

// Implements direct to memory image formats.  It
// returns the image as a byte array (a formatted "file" in memory),
// starting from the current position in the image sequence.
// Use MagickSetImageFormat() to set the format to write to the blob
// (GIF, JPEG,  PNG, etc.).
func (mw *MagickWand) WriteImageBlob() []byte {
	length := C.size_t(0)
	d := C.MagickWriteImageBlob(mw.mw, &length)
	if d != nil {
		defer C.MagickRelinquishMemory(unsafe.Pointer(d))
		return C.GoBytes(unsafe.Pointer(d), C.int(length))
	}
	return make([]byte, 0)
}

// Writes an image or image sequence.
func (mw *MagickWand) WriteImages(filename string, adjoin bool) error {
	csfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(csfilename))
	ok := C.MagickWriteImages(mw.mw, csfilename, b2i(adjoin))
	return mw.getLastErrorIfFailed(ok)
}

// Writes an image sequence to an open file descriptor.
func (mw *MagickWand) WriteImagesFile(out *os.File, adjoin bool) error {
	file, err := cfdopen(out, "w")
	if err != nil {
		return err
	}
	defer C.fclose(file)
	ok := C.MagickWriteImagesFile(mw.mw, file, b2i(adjoin))
	return mw.getLastErrorIfFailed(ok)
}



// cfdopen returns a C-level FILE*. mode should be as described in fdopen(3).
// Caller is responsible for closing the file when successfully returned,
// via C.fclose()
func cfdopen(file *os.File, mode string) (*C.FILE, error) {
	cmode := C.CString(mode)
	defer C.free(unsafe.Pointer(cmode))

	cfile, err := C.fdopen(C.dup(C.int(file.Fd())), cmode)
	if err != nil {
		return nil, err
	}
	if cfile == nil {
		return nil, syscall.EINVAL
	}

	return cfile, nil
}

