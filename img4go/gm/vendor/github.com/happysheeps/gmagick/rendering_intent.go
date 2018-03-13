package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

type RenderingIntent int

const (
	RENDERING_INTENT_UNDEFINED  RenderingIntent = C.UndefinedIntent
	RENDERING_INTENT_SATURATION RenderingIntent = C.SaturationIntent
	RENDERING_INTENT_PERCEPTUAL RenderingIntent = C.PerceptualIntent
	RENDERING_INTENT_ABSOLUTE   RenderingIntent = C.AbsoluteIntent
	RENDERING_INTENT_RELATIVE   RenderingIntent = C.RelativeIntent
)