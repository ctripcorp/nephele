package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

type FillRule int

const (
	FILL_UNDEFINED FillRule = C.UndefinedRule
	FILL_EVEN_ODD  FillRule = C.EvenOddRule
	FILL_NON_ZERO  FillRule = C.NonZeroRule
)