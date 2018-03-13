# Go bindings for [GraphicsMagick](http://www.graphicsmagick.org)

## Install `GraphicsMagick` libraries and header files

### Windows
+ Install [msys2-x86_64](http://www.msys2.org/)
+ In msys2 shell: 
```
pacman -S mingw-w64-x86_64-gcc
pacman -S mingw-w64-x86_64-zlib
pacman -S mingw-w64-x86_64-pkg-config
pacman -S mingw-w64-x86_64-graphicsmagick
```
+ Add following environment variable:
```
set PATH=<msys64>\mingw64\bin;%PATH%
set PKG_CONFIG_PATH=<msys64>\mingw64\lib\pkgconfig
set MAGICK_CODER_MODULE_PATH=<msys64>\mingw64\lib\GraphicsMagick-1.3.25\modules-Q8\coders
```
(BTW: you should change `<msys64>` to your installation path of `msys2`)

### MacOS
Install `GraphicsMagick` using [Homebrew](https://brew.sh/) or [MacPorts](https://www.macports.org)    
An example of `MacPorts`:

+ `sudo port install graphicsmagick`
+ `export PKG_CONFIG_PATH=/opt/local/lib/pkgconfig`


### CentOS

+ ensure `epel-release` was installed. ([help](https://pkgs.org/download/epel-release))
+ `yum install GraphicsMagick-devel`

### Ubuntu

+ `sudo apt-get install libgraphicsmagick1-dev`

## Install golang bindings

+ `go get github.com/gographics/gmagick`


## Example

```
package main

import (
    "flag"
    "github.com/gographics/gmagick"
)

func resize(orig string, dest string) {
    mw := gmagick.NewMagickWand()
    defer mw.Destroy()
    mw.ReadImage(orig)
    filter := gmagick.FILTER_LANCZOS
    w := mw.GetImageWidth()
    h := mw.GetImageHeight()
    mw.ResizeImage(w/2, h/2, filter, 1)
    mw.WriteImage(dest)
}

func main() {
    f := flag.String("from", "", "original image file ...")
    t := flag.String("to", "", "target file ...")
    flag.Parse()

    gmagick.Initialize()
    defer gmagick.Terminate()

    resize(*f, *t)
}
```

====================================================
## Todo List


- [x] CloneMagickWand
- [x] ClonePixelWand
- [ ] ClonePixelWands
- [ ] CopyMagickString
- [x] DestroyMagickWand
- [x] DestroyPixelWand
- [ ] FormatMagickString
- [ ] FormatMagickStringList
- [x] MagickAdaptiveThresholdImage
- [x] MagickAddImage
- [x] MagickAddNoiseImage
- [x] MagickAffineTransformImage
- [x] MagickAnimateImages 
- [x] MagickAnnotateImage
- [x] MagickAppendImages
- [x] MagickAverageImages
- [x] MagickBlackThresholdImage
- [x] MagickBlurImage 
- [x] MagickBorderImage 
- [ ] MagickCdlImage 
- [x] MagickCharcoalImage 
- [x] MagickChopImage 
- [x] MagickClipImage
- [ ] MagickClipPathImage 
- [x] MagickCloneDrawingWand
- [x] MagickCoalesceImages 
- [ ] MagickColorFloodfillImage 
- [x] MagickColorizeImage 
- [x] MagickCommentImage 
- [x] MagickCompareImageChannels
- [x] MagickCompareImages
- [x] MagickCompositeImage 
- [x] MagickContrastImage 
- [x] MagickConvolveImage 
- [x] MagickCropImage 
- [x] MagickCycleColormapImage
- [x] MagickDeconstructImages
- [x] MagickDescribeImage 
- [x] MagickDespeckleImage 
- [x] MagickDestroyDrawingWand 
- [x] MagickDisplayImage 
- [x] MagickDisplayImages 
- [x] MagickDrawAffine 
- [ ] MagickDrawAllocateWand 
- [x] MagickDrawAnnotation 
- [x] MagickDrawArc 
- [x] MagickDrawBezier 
- [x] MagickDrawCircle 
- [ ] MagickDrawClearException 
- [x] MagickDrawColor 
- [x] MagickDrawComment 
- [x] MagickDrawComposite 
- [x] MagickDrawEllipse 
- [x] MagickDrawGetClipPath 
- [x] MagickDrawGetClipRule 
- [x] MagickDrawGetClipUnits 
- [ ] MagickDrawGetException 
- [x] MagickDrawGetFillColor 
- [x] MagickDrawGetFillOpacity 
- [x] MagickDrawGetFillRule 
- [x] MagickDrawGetFontFamily 
- [x] MagickDrawGetFont 
- [x] MagickDrawGetFontSize 
- [x] MagickDrawGetFontStretch 
- [x] MagickDrawGetFontStyle 
- [x] MagickDrawGetFontWeight 
- [x] MagickDrawGetGravity 
- [x] MagickDrawGetStrokeAntialias 
- [x] MagickDrawGetStrokeColor 
- [x] MagickDrawGetStrokeDashArray 
- [x] MagickDrawGetStrokeDashOffset 
- [x] MagickDrawGetStrokeLineCap 
- [x] MagickDrawGetStrokeLineJoin 
- [x] MagickDrawGetStrokeMiterLimit 
- [x] MagickDrawGetStrokeOpacity 
- [x] MagickDrawGetStrokeWidth 
- [x] MagickDrawGetTextAntialias 
- [x] MagickDrawGetTextDecoration 
- [x] MagickDrawGetTextEncoding 
- [x] MagickDrawGetTextUnderColor 
- [x] MagickDrawImage 
- [x] MagickDrawLine 
- [x] MagickDrawMatte 
- [x] MagickDrawPathClose 
- [x] MagickDrawPathCurveToAbsolute 
- [x] MagickDrawPathCurveToQuadraticBezierAbsolute 
- [x] MagickDrawPathCurveToQuadraticBezierRelative 
- [x] MagickDrawPathCurveToQuadraticBezierSmoothAbsolute
- [x] MagickDrawPathCurveToQuadraticBezierSmoothRelative
- [x] MagickDrawPathCurveToRelative 
- [x] MagickDrawPathCurveToSmoothAbsolute
- [x] MagickDrawPathCurveToSmoothRelative
- [x] MagickDrawPathEllipticArcAbsolute
- [x] MagickDrawPathEllipticArcRelative 
- [x] MagickDrawPathFinish
- [x] MagickDrawPathLineToAbsolute 
- [x] MagickDrawPathLineToHorizontalAbsolute 
- [x] MagickDrawPathLineToHorizontalRelative
- [x] MagickDrawPathLineToRelative
- [x] MagickDrawPathLineToVerticalAbsolute 
- [x] MagickDrawPathLineToVerticalRelative 
- [x] MagickDrawPathMoveToAbsolute 
- [x] MagickDrawPathMoveToRelative 
- [x] MagickDrawPathStart 
- [ ] MagickDrawPeekGraphicContext 
- [x] MagickDrawPoint 
- [x] MagickDrawPolygon 
- [x] MagickDrawPolyline 
- [x] MagickDrawPopClipPath 
- [x] MagickDrawPopDefs 
- [ ] MagickDrawPopGraphicContext 
- [x] MagickDrawPopPattern 
- [x] MagickDrawPushClipPath 
- [x] MagickDrawPushDefs 
- [ ] MagickDrawPushGraphicContext 
- [x] MagickDrawPushPattern
- [x] MagickDrawRectangle 
- [ ] MagickDrawRender(deprecated)
- [x] MagickDrawRotate 
- [x] MagickDrawRoundRectangle
- [x] MagickDrawScale
- [x] MagickDrawSetClipPath 
- [x] MagickDrawSetClipRule 
- [x] MagickDrawSetClipUnits 
- [x] MagickDrawSetFillColor 
- [x] MagickDrawSetFillOpacity 
- [x] MagickDrawSetFillPatternURL 
- [x] MagickDrawSetFillRule 
- [x] MagickDrawSetFontFamily 
- [x] MagickDrawSetFont 
- [x] MagickDrawSetFontSize 
- [x] MagickDrawSetFontStretch 
- [x] MagickDrawSetFontStyle 
- [x] MagickDrawSetFontWeight 
- [x] MagickDrawSetGravity 
- [x] MagickDrawSetStrokeAntialias 
- [x] MagickDrawSetStrokeColor 
- [x] MagickDrawSetStrokeDashArray 
- [x] MagickDrawSetStrokeDashOffset 
- [x] MagickDrawSetStrokeLineCap 
- [x] MagickDrawSetStrokeLineJoin 
- [x] MagickDrawSetStrokeMiterLimit 
- [x] MagickDrawSetStrokeOpacity 
- [x] MagickDrawSetStrokePatternURL 
- [x] MagickDrawSetStrokeWidth 
- [x] MagickDrawSetTextAntialias 
- [x] MagickDrawSetTextDecoration 
- [x] MagickDrawSetTextEncoding 
- [x] MagickDrawSetTextUnderColor 
- [x] MagickDrawSetViewbox 
- [x] MagickDrawSkewX 
- [x] MagickDrawSkewY 
- [x] MagickDrawTranslate 
- [X] MagickEdgeImage
- [x] MagickEmbossImage 
- [x] MagickEnhanceImage 
- [x] MagickEqualizeImage 
- [x] MagickExtentImage 
- [ ] MagickFlattenImages 
- [x] MagickFlipImage 
- [x] MagickFlopImage 
- [x] MagickFrameImage 
- [x] MagickFxImageChannel 
- [x] MagickFxImage 
- [x] MagickGammaImageChannel 
- [x] MagickGammaImage 
- [ ] MagickGetConfigureInfo
- [x] MagickGetCopyright
- [x] MagickGetException
- [x] MagickGetFilename
- [x] MagickGetHomeURL
- [x] MagickGetImageAttribute
- [x] MagickGetImageBackgroundColor
- [x] MagickGetImageBluePrimary
- [x] MagickGetImageBorderColor
- [ ] MagickGetImageBoundingBox
- [x] MagickGetImageChannelDepth
- [ ] MagickGetImageChannelExtrema
- [x] MagickGetImageChannelMean 
- [x] MagickGetImageColormapColor 
- [x] MagickGetImageColors 
- [x] MagickGetImageColorspace 
- [x] MagickGetImageCompose 
- [x] MagickGetImageCompression 
- [x] MagickGetImageDelay 
- [x] MagickGetImageDepth 
- [x] MagickGetImageDispose 
- [ ] MagickGetImageExtrema 
- [x] MagickGetImageFilename 
- [x] MagickGetImageFormat 
- [x] MagickGetImageFuzz 
- [x] MagickGetImageGamma 
- [x] MagickGetImageGeometry 
- [x] MagickGetImageGravity 
- [x] MagickGetImage 
- [x] MagickGetImageGreenPrimary 
- [x] MagickGetImageHeight 
- [x] MagickGetImageHistogram 
- [x] MagickGetImageIndex 
- [x] MagickGetImageInterlaceScheme 
- [x] MagickGetImageIterations 
- [ ] MagickGetImageMatte 
- [x] MagickGetImageMatteColor 
- [x] MagickGetImagePage 
- [ ] MagickGetImagePixels 
- [ ] MagickGetImageProfile 
- [x] MagickGetImageRedPrimary 
- [ ] MagickGetImageRenderingIntent 
- [x] MagickGetImageResolution 
- [x] MagickGetImageSavedType 
- [ ] MagickGetImageScene 
- [x] MagickGetImageSignature 
- [ ] MagickGetImageSize 
- [x] MagickGetImageType 
- [x] MagickGetImageUnits 
- [ ] MagickGetImageVirtualPixelMethod 
- [x] MagickGetImageWhitePoint 
- [x] MagickGetImageWidth 
- [x] MagickGetNumberImages 
- [x] MagickGetPackageName 
- [x] MagickGetQuantumDepth 
- [x] MagickGetReleaseDate 
- [x] MagickGetResourceLimit 
- [x] MagickGetSamplingFactors 
- [x] MagickGetSize 
- [x] MagickGetVersion 
- [x] MagickHaldClutImage 
- [x] MagickHasNextImage 
- [x] MagickHasPreviousImage 
- [x] MagickImplodeImage 
- [x] MagickLabelImage 
- [x] MagickLevelImageChannel 
- [x] MagickLevelImage 
- [x] MagickMagnifyImage 
- [x] MagickMapImage 
- [x] MagickMatteFloodfillImage 
- [x] MagickMedianFilterImage 
- [x] MagickMinifyImage 
- [x] MagickModulateImage 
- [x] MagickMontageImage 
- [x] MagickMorphImages 
- [x] MagickMosaicImages 
- [x] MagickMotionBlurImage 
- [x] MagickNegateImageChannel 
- [x] MagickNegateImage 
- [x] MagickNewDrawingWand 
- [x] MagickNextImage 
- [x] MagickNormalizeImage 
- [x] MagickOilPaintImage 
- [x] MagickOpaqueImage 
- [x] MagickPingImage 
- [x] MagickPreviewImages 
- [x] MagickPreviousImage 
- [ ] MagickProfileImage 
- [x] MagickQuantizeImage 
- [x] MagickQuantizeImages 
- [x] MagickQueryFontMetrics 
- [x] MagickQueryFonts 
- [x] MagickQueryFormats 
- [x] MagickRadialBlurImage 
- [x] MagickRaiseImage 
- [x] MagickReadImageBlob 
- [x] MagickReadImageFile 
- [x] MagickReadImage 
- [x] MagickReduceNoiseImage 
- [ ] MagickRelinquishMemory 
- [x] MagickRemoveImage
- [x] MagickRemoveImageProfile
- [x] MagickResampleImage
- [x] MagickResetIterator
- [x] MagickResizeImage
- [x] MagickRollImage 
- [x] MagickRotateImage 
- [x] MagickSampleImage 
- [x] MagickScaleImage 
- [x] MagickSeparateImageChannel 
- [x] MagickSetCompressionQuality 
- [ ] MagickSetDepth
- [ ] MagickSetFilename 
- [x] MagickSetFormat
- [x] MagickSetImageAttribute 
- [x] MagickSetImageBackgroundColor 
- [x] MagickSetImageBluePrimary 
- [x] MagickSetImageBorderColor 
- [x] MagickSetImageChannelDepth 
- [x] MagickSetImageColormapColor 
- [x] MagickSetImageColorspace 
- [x] MagickSetImageCompose 
- [x] MagickSetImageCompression
- [x] MagickSetImageDelay
- [x] MagickSetImageDepth
- [x] MagickSetImageDispose
- [x] MagickSetImageFilename
- [x] MagickSetImageFormat
- [x] MagickSetImageFuzz
- [x] MagickSetImageGamma
- [x] MagickSetImageGravity
- [x] MagickSetImage
- [x] MagickSetImageGreenPrimary
- [x] MagickSetImageIndex
- [x] MagickSetImageInterlaceScheme
- [x] MagickSetImageIterations
- [x] MagickSetImageMatte
- [x] MagickSetImageMatteColor
- [ ] MagickSetImageOption
- [x] MagickSetImagePage
- [ ] MagickSetImagePixels
- [ ] MagickSetImageProfile
- [x] MagickSetImageRedPrimary
- [x] MagickSetImageRenderingIntent
- [x] MagickSetImageResolution
- [x] MagickSetImageSavedType
- [x] MagickSetImageScene
- [x] MagickSetImageType
- [x] MagickSetImageUnits
- [x] MagickSetImageVirtualPixelMethod
- [x] MagickSetImageWhitePoint
- [x] MagickSetInterlaceScheme
- [x] MagickSetPassphrase
- [x] MagickSetResolution
- [x] MagickSetResolutionUnits
- [x] MagickSetResourceLimit
- [x] MagickSetSamplingFactors
- [x] MagickSetSize 
- [x] MagickSharpenImage 
- [x] MagickShaveImage 
- [x] MagickShearImage 
- [x] MagickSolarizeImage 
- [x] MagickSpreadImage 
- [x] MagickSteganoImage 
- [x] MagickStereoImage 
- [x] MagickStripImage 
- [x] MagickSwirlImage 
- [x] MagickTextureImage
- [x] MagickThresholdImageChannel
- [x] MagickThresholdImage
- [x] MagickTintImage
- [x] MagickTransformImage
- [x] MagickTransparentImage
- [x] MagickTrimImage
- [x] MagickUnsharpMaskImage
- [x] MagickWaveImage
- [x] MagickWhiteThresholdImage
- [x] MagickWriteImageBlob 
- [x] MagickWriteImageFile 
- [x] MagickWriteImage 
- [x] MagickWriteImagesFile 
- [x] MagickWriteImages 
- [x] NewMagickWand
- [x] NewPixelWand
- [ ] NewPixelWands
- [x] PixelGetBlack 
- [x] PixelGetBlackQuantum 
- [x] PixelGetBlue 
- [x] PixelGetBlueQuantum 
- [x] PixelGetColorAsString 
- [x] PixelGetColorCount 
- [x] PixelGetCyan 
- [x] PixelGetCyanQuantum 
- [ ] PixelGetException 
- [x] PixelGetGreen 
- [x] PixelGetGreenQuantum 
- [x] PixelGetMagenta 
- [x] PixelGetMagentaQuantum 
- [x] PixelGetOpacity 
- [x] PixelGetOpacityQuantum 
- [x] PixelGetQuantumColor 
- [x] PixelGetRed 
- [x] PixelGetRedQuantum 
- [x] PixelGetYellow 
- [x] PixelGetYellowQuantum 
- [x] PixelSetBlack 
- [x] PixelSetBlackQuantum 
- [x] PixelSetBlue 
- [x] PixelSetBlueQuantum 
- [x] PixelSetColorCount 
- [x] PixelSetColor 
- [x] PixelSetCyan 
- [x] PixelSetCyanQuantum 
- [x] PixelSetGreen 
- [x] PixelSetGreenQuantum 
- [x] PixelSetMagenta 
- [x] PixelSetMagentaQuantum 
- [x] PixelSetOpacity 
- [x] PixelSetOpacityQuantum 
- [x] PixelSetQuantumColor 
- [x] PixelSetRed 
- [x] PixelSetRedQuantum 
- [x] PixelSetYellow 
- [x] PixelSetYellowQuantum 
- [ ] QueryMagickColor 
