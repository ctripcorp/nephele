package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

type StorageType int

const (
	CharPixel			StorageType = C.CharPixel
	ShortPixel			StorageType = C.ShortPixel
	IntegerPixel		StorageType = C.IntegerPixel
	LongPixel			StorageType = C.LongPixel
	FloatPixel			StorageType = C.FloatPixel
	DoublePixel			StorageType = C.DoublePixel
)
