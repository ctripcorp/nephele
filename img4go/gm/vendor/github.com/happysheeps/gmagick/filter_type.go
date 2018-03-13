package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

type FilterType int

const (
	FILTER_UNDEFINED      FilterType = C.UndefinedFilter
	FILTER_POINT          FilterType = C.PointFilter
	FILTER_BOX            FilterType = C.BoxFilter
	FILTER_TRIANGLE       FilterType = C.TriangleFilter
	FILTER_HERMITE        FilterType = C.HermiteFilter
	FILTER_HANNING        FilterType = C.HanningFilter
	FILTER_HAMMING        FilterType = C.HammingFilter
	FILTER_BLACKMAN       FilterType = C.BlackmanFilter
	FILTER_GAUSSIAN       FilterType = C.GaussianFilter
	FILTER_QUADRATIC      FilterType = C.QuadraticFilter
	FILTER_CUBIC          FilterType = C.CubicFilter
	FILTER_CATROM         FilterType = C.CatromFilter
	FILTER_MITCHELL       FilterType = C.MitchellFilter
	FILTER_SINC           FilterType = C.SincFilter
	FILTER_LANCZOS        FilterType = C.LanczosFilter
)
