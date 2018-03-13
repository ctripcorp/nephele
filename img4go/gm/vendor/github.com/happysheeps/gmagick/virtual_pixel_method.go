package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

type VirtualPixelMethod int

const (
	VIRTUAL_PIXEL_UNDEFINED            VirtualPixelMethod = C.UndefinedVirtualPixelMethod
	VIRTUAL_PIXEL_CONSTANT             VirtualPixelMethod = C.ConstantVirtualPixelMethod
	VIRTUAL_PIXEL_EDGE                 VirtualPixelMethod = C.EdgeVirtualPixelMethod
	VIRTUAL_PIXEL_MIRROR               VirtualPixelMethod = C.MirrorVirtualPixelMethod
	VIRTUAL_PIXEL_TILE                 VirtualPixelMethod = C.TileVirtualPixelMethod
)