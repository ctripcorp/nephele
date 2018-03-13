package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

type OrientationType int

const (
	UndefinedOrientation			OrientationType = C.UndefinedOrientation
	TopLeftOrientation				OrientationType = C.TopLeftOrientation
	TopRightOrientation				OrientationType = C.TopRightOrientation
	BottomRightOrientation			OrientationType = C.BottomRightOrientation
	BottomLeftOrientation			OrientationType = C.BottomLeftOrientation
	LeftTopOrientation				OrientationType = C.LeftTopOrientation
	RightTopOrientation				OrientationType = C.RightTopOrientation
	RightBottomOrientation		    OrientationType = C.RightBottomOrientation
	LeftBottomOrientation		    OrientationType = C.LeftBottomOrientation
)
