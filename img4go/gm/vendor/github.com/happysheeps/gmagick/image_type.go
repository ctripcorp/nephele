package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

type ImageType int

const (
	IMAGE_TYPE_UNDEFINED              ImageType = C.UndefinedType
	IMAGE_TYPE_BILEVEL                ImageType = C.BilevelType
	IMAGE_TYPE_GRAYSCALE              ImageType = C.GrayscaleType
	IMAGE_TYPE_GRAYSCALE_MATTE        ImageType = C.GrayscaleMatteType
	IMAGE_TYPE_PALETTE                ImageType = C.PaletteType
	IMAGE_TYPE_PALETTE_MATTE          ImageType = C.PaletteMatteType
	IMAGE_TYPE_TRUE_COLOR             ImageType = C.TrueColorType
	IMAGE_TYPE_TRUE_COLOR_MATTE       ImageType = C.TrueColorMatteType
	IMAGE_TYPE_COLOR_SEPARATION       ImageType = C.ColorSeparationType
	IMAGE_TYPE_COLOR_SEPARATION_MATTE ImageType = C.ColorSeparationMatteType
	IMAGE_TYPE_OPTIMIZE               ImageType = C.OptimizeType
)