package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

type InterlaceType int

const (
	INTERLACE_UNDEFINED InterlaceType = C.UndefinedInterlace
	INTERLACE_NO        InterlaceType = C.NoInterlace
	INTERLACE_LINE      InterlaceType = C.LineInterlace
	INTERLACE_PLANE     InterlaceType = C.PlaneInterlace
	INTERLACE_PARTITION InterlaceType = C.PartitionInterlace
)