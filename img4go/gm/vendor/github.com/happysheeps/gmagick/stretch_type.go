package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

type StretchType int

const (
	STRETCH_NORMAL          StretchType = C.NormalStretch
	STRETCH_ULTRA_CONDENSED StretchType = C.UltraCondensedStretch
	STRETCH_EXTRA_CONDENSED StretchType = C.ExtraCondensedStretch
	STRETCH_CONDENSED       StretchType = C.CondensedStretch
	STRETCH_SEMI_CONDENSED  StretchType = C.SemiCondensedStretch
	STRETCH_SEMI_EXPANDED   StretchType = C.SemiExpandedStretch
	STRETCH_EXPANDED        StretchType = C.ExpandedStretch
	STRETCH_EXTRA_EXPANDED  StretchType = C.ExtraExpandedStretch
	STRETCH_ULTRA_EXPANDED  StretchType = C.UltraExpandedStretch
	STRETCH_ANY             StretchType = C.AnyStretch
)