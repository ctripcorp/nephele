package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

type LineCap int

const (
	LINE_CAP_UNDEFINED LineCap = C.UndefinedCap
	LINE_CAP_BUTT      LineCap = C.ButtCap
	LINE_CAP_ROUND     LineCap = C.RoundCap
	LINE_CAP_SQUARE    LineCap = C.SquareCap
)