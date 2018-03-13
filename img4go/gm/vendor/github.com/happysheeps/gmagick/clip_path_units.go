package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

type ClipPathUnits int

const (
	CLIP_USER_SPACE          ClipPathUnits = C.UserSpace
	CLIP_USER_SPACE_ON_USE   ClipPathUnits = C.UserSpaceOnUse
	CLIP_OBJECT_BOUNDING_BOX ClipPathUnits = C.ObjectBoundingBox
)