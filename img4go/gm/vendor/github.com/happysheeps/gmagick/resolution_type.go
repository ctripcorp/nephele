package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

type ResolutionType int

const (
	RESOLUTION_UNDEFINED             ResolutionType = C.UndefinedResolution
	RESOLUTION_PIXELS_PER_INCH       ResolutionType = C.PixelsPerInchResolution
	RESOLUTION_PIXELS_PER_CENTIMETER ResolutionType = C.PixelsPerCentimeterResolution
)