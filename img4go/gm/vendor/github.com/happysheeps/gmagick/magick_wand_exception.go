package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

type MagickWandException struct {
	kind        ExceptionType
	description string
}

func (mwe *MagickWandException) Error() string {
	return fmt.Sprintf("%s: %s", mwe.kind.String(), mwe.description)
}

// Returns the kind, reason and description of any error that occurs when using other methods in this API
func (mw *MagickWand) GetLastError() error {
	var et C.ExceptionType
	csdescription := C.MagickGetException(mw.mw, &et)
	defer C.MagickRelinquishMemory(unsafe.Pointer(csdescription))
	if ExceptionType(et) != EXCEPTION_UNDEFINED {
		return &MagickWandException{ExceptionType(C.int(et)), C.GoString(csdescription)}
	}
	return errors.New("undefined exception")
}

func (mw *MagickWand) getLastErrorIfFailed(ok C.uint) error {
	if ok == 0 {
		return mw.GetLastError()
	} else {
		return nil
	}
}
