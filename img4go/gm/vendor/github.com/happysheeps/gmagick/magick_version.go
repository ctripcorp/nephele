package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

import "unsafe"

// Returns the ImageMagick API copyright as a string constant.
func GetCopyright() string {
	cstr := C.MagickGetCopyright()
	return C.GoString(cstr)
}

// Returns the ImageMagick home URL.
func GetHomeURL() string {
	cstr := C.MagickGetHomeURL()
	defer C.MagickRelinquishMemory(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// Returns the ImageMagick package name as a string constant.
func GetPackageName() string {
	cstr := C.MagickGetPackageName()
	return C.GoString(cstr)
}

// Returns the ImageMagick release date as a string constant.
func GetReleaseDate() string {
	cstr := C.MagickGetReleaseDate()
	return C.GoString(cstr)
}

// Returns the ImageMagick quantum depth as a string constant.
func GetQuantumDepth() (string, uint) {
	cst := C.ulong(0)
	csq := C.MagickGetQuantumDepth(&cst)
	return C.GoString(csq), uint(cst)
}

// Returns the specified resource limit in megabytes.
func GetResourceLimit(rtype ResourceType) int64 {
	return int64(C.MagickGetResourceLimit(C.ResourceType(rtype)))
}

// Returns the ImageMagick API version as a string constant and as a number.
func GetVersion() (string, uint) {
	cnver := C.ulong(0)
	csver := C.MagickGetVersion(&cnver)
	version := C.GoString(csver)
	nversion := uint(cnver)
	return version, nversion
}