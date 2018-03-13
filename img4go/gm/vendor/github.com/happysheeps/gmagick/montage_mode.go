package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

type MontageMode int

const (
	MONTAGE_MODE_UNDEFINED   MontageMode = C.UndefinedMode
	MONTAGE_MODE_FRAME       MontageMode = C.FrameMode
	MONTAGE_MODE_UNFRAME     MontageMode = C.UnframeMode
	MONTAGE_MODE_CONCATENATE MontageMode = C.ConcatenateMode
)
