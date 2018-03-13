package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
	"errors"
)

type DrawingWandException struct {
	kind        ExceptionType
	description string
}

func (dwe *DrawingWandException) Error() string {
	return fmt.Sprintf("%s: %s", dwe.kind.String(), dwe.description)
}

// Clears any exceptions associated with the wand
func (dw *DrawingWand) clearException() bool {
	return 1 == C.int(C.DrawClearException(dw.dw))
}

// Returns the kind, reason and description of any error that occurs when using other methods in this API
func (dw *DrawingWand) GetLastError() error {
	var et C.ExceptionType
	csdescription := C.DrawGetException(dw.dw, &et)
	defer C.MagickRelinquishMemory(unsafe.Pointer(csdescription))
	if ExceptionType(et) != EXCEPTION_UNDEFINED {
		dw.clearException()
		return &DrawingWandException{ExceptionType(C.int(et)), C.GoString(csdescription)}
	}
	return errors.New("undefined exception")
}

func (dw *DrawingWand) getLastErrorIfFailed(ok C.uint) error {
	if ok == 0 {
		return dw.GetLastError()
	} else {
		return nil
	}
}