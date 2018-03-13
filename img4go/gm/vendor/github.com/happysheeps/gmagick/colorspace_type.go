package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

type ColorspaceType int

const (
	COLORSPACE_UNDEFINED   ColorspaceType = C.UndefinedColorspace
	COLORSPACE_RGB         ColorspaceType = C.RGBColorspace
	COLORSPACE_GRAY        ColorspaceType = C.GRAYColorspace
	COLORSPACE_TRANSPARENT ColorspaceType = C.TransparentColorspace
	COLORSPACE_OHTA        ColorspaceType = C.OHTAColorspace
	COLORSPACE_XYZ         ColorspaceType = C.XYZColorspace
	COLORSPACE_YCBCR       ColorspaceType = C.YCbCrColorspace
	COLORSPACE_YCC         ColorspaceType = C.YCCColorspace
	COLORSPACE_YIQ         ColorspaceType = C.YIQColorspace
	COLORSPACE_YPBPR       ColorspaceType = C.YPbPrColorspace
	COLORSPACE_YUV         ColorspaceType = C.YUVColorspace
	COLORSPACE_CMYK        ColorspaceType = C.CMYKColorspace
	COLORSPACE_SRGB        ColorspaceType = C.sRGBColorspace
	COLORSPACE_HSL         ColorspaceType = C.HSLColorspace
	COLORSPACE_HWB         ColorspaceType = C.HWBColorspace
	COLORSPACE_REC601LUMA  ColorspaceType = C.Rec601LumaColorspace
	COLORSPACE_REC601YCBCR ColorspaceType = C.Rec601YCbCrColorspace
	COLORSPACE_REC709LUMA  ColorspaceType = C.Rec709LumaColorspace
	COLORSPACE_REC709YCBCR ColorspaceType = C.Rec709YCbCrColorspace
)
