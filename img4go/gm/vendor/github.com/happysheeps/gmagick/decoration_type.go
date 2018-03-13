package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

type DecorationType int

const (
	DECORATION_NONE         DecorationType = C.NoDecoration
	DECORATION_UNDERLINE    DecorationType = C.UnderlineDecoration
	DECORATION_OVERLINE     DecorationType = C.OverlineDecoration
	DECORATION_LINE_THROUGH DecorationType = C.LineThroughDecoration
)