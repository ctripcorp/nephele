package gmagick

import "C"
import "unsafe"

// Convert a boolean to a integer
func b2i(boolean bool) C.uint {
	if boolean {
		return C.uint(1)
	}
	return C.uint(0)
}

func sizedDoubleArrayToFloat64Slice(p *C.double, num C.ulong) []float64 {
	var nums []float64
	q := uintptr(unsafe.Pointer(p))
	for i := 0; i < int(num); i++ {
		p = (*C.double)(unsafe.Pointer(q))
		nums = append(nums, float64(*p))
		q += unsafe.Sizeof(q)
	}
	return nums
}

func sizedCStringArrayToStringSlice(p **C.char, num C.ulong) []string {
	var strings []string
	q := uintptr(unsafe.Pointer(p))
	for i := 0; i < int(num); i++ {
		p = (**C.char)(unsafe.Pointer(q))
		if *p == nil {
			break
		}
		strings = append(strings, C.GoString(*p))
		q += unsafe.Sizeof(q)
	}
	return strings
}
