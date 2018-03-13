package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

type PixelPacket struct {
	pp *C.PixelPacket
}

func newPixelPacketFromCAPI(mpp *C.PixelPacket) *PixelPacket {
	return &PixelPacket{mpp}
}