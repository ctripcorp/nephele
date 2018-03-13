package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

import (
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"
)

type PixelWand struct {
	pw   *C.PixelWand
	init sync.Once
}

// Returns a new pixel wand
func newPixelWand(cpw *C.PixelWand) *PixelWand {
	pw := &PixelWand{pw: cpw}
	runtime.SetFinalizer(pw, Destroy)
	pw.IncreaseCount()

	return pw
}

// Returns a new pixel wand
func NewPixelWand() *PixelWand {
	return newPixelWand(C.NewPixelWand())
}

// Makes an exact copy of the wand
func (pw *PixelWand) Clone() *PixelWand {
	return newPixelWand(C.ClonePixelWand(pw.pw))
}

// Deallocates resources associated with a pixel wand
func (pw *PixelWand) Destroy() {
	if pw.pw == nil {
		return
	}

	pw.init.Do(func() {
		C.DestroyPixelWand(pw.pw)
		pw.pw = nil

		pw.DecreaseCount()
	})
}

// Increase PixelWand ref counter and set according "can`t be terminated status"
func (pw *PixelWand) IncreaseCount() {
	atomic.AddInt64(&pixelWandCounter, int64(1))
	unsetCanTerminate()
}

// Decrease PixelWand ref counter and set according "can be terminated status"
func (pw *PixelWand) DecreaseCount() {
	atomic.AddInt64(&pixelWandCounter, int64(-1))
	setCanTerminate()
}

// Returns the normalized black color of the pixel wand
func (pw *PixelWand) GetBlack() float64 {
	return float64(C.PixelGetBlack(pw.pw))
}

// Returns the black color of the pixel wand
func (pw *PixelWand) GetBlackQuantum() Quantum {
	return Quantum(C.PixelGetBlackQuantum(pw.pw))
}

// Returns the normalized blue color of the pixel wand
func (pw *PixelWand) GetBlue() float64 {
	return float64(C.PixelGetBlue(pw.pw))
}

// Returns the blue color of the pixel wand
func (pw *PixelWand) GetBlueQuantum() Quantum {
	return Quantum(C.PixelGetBlueQuantum(pw.pw))
}

// Returns the color of the pixel wand as a string
func (pw *PixelWand) GetColorAsString() string {
	p := C.PixelGetColorAsString(pw.pw)
	defer C.MagickRelinquishMemory(unsafe.Pointer(p))
	return C.GoString(p)
}

// Returns the color count associated with this color
func (pw *PixelWand) GetColorCount() uint {
	return uint(C.PixelGetColorCount(pw.pw))
}

// Returns the normalized cyan color of the pixel wand
func (pw *PixelWand) GetCyan() float64 {
	return float64(C.PixelGetCyan(pw.pw))
}

// Returns the cyan color of the pixel wand
func (pw *PixelWand) GetCyanQuantum() Quantum {
	return Quantum(C.PixelGetCyanQuantum(pw.pw))
}

// Returns the normalized green color of the pixel wand
func (pw *PixelWand) GetGreen() float64 {
	return float64(C.PixelGetGreen(pw.pw))
}

// Returns the green color of the pixel wand
func (pw *PixelWand) GetGreenQuantum() Quantum {
	return Quantum(C.PixelGetGreenQuantum(pw.pw))
}

// Returns the normalized magenta color of the pixel wand
func (pw *PixelWand) GetMagenta() float64 {
	return float64(C.PixelGetMagenta(pw.pw))
}

// Returns the magenta color of the pixel wand
func (pw *PixelWand) GetMagentaQuantum() Quantum {
	return Quantum(C.PixelGetMagentaQuantum(pw.pw))
}

// Returns the normalized opacity color of the pixel wand
func (pw *PixelWand) GetOpacity() float64 {
	return float64(C.PixelGetOpacity(pw.pw))
}

// Returns the opacity color of the pixel wand
func (pw *PixelWand) GetOpacityQuantum() Quantum {
	return Quantum(C.PixelGetOpacityQuantum(pw.pw))
}

// Gets the color of the pixel wand as a PixelPacket
func (pw *PixelWand) GetQuantumColor() *PixelPacket {
	var pp C.PixelPacket
	C.PixelGetQuantumColor(pw.pw, &pp)
	return newPixelPacketFromCAPI(&pp)
}

// Returns the normalized red color of the pixel wand
func (pw *PixelWand) GetRed() float64 {
	return float64(C.PixelGetRed(pw.pw))
}

// Returns the red color of the pixel wand
func (pw *PixelWand) GetRedQuantum() Quantum {
	return Quantum(C.PixelGetRedQuantum(pw.pw))
}

// Returns the normalized yellow color of the pixel wand
func (pw *PixelWand) GetYellow() float64 {
	return float64(C.PixelGetYellow(pw.pw))
}

// Returns the yellow color of the pixel wand
func (pw *PixelWand) GetYellowQuantum() Quantum {
	return Quantum(C.PixelGetYellowQuantum(pw.pw))
}

// Sets the normalized black color of the pixel wand
func (pw *PixelWand) SetBlack(black float64) {
	C.PixelSetBlack(pw.pw, C.double(black))
}

// Sets the black color of the pixel wand
func (pw *PixelWand) SetBlackQuantum(black Quantum) {
	C.PixelSetBlackQuantum(pw.pw, C.Quantum(black))
}

// Sets the normalized blue color of the pixel wand
func (pw *PixelWand) SetBlue(blue float64) {
	C.PixelSetBlue(pw.pw, C.double(blue))
}

// Sets the blue color of the pixel wand
func (pw *PixelWand) SetBlueQuantum(blue Quantum) {
	C.PixelSetBlueQuantum(pw.pw, C.Quantum(blue))
}

// Sets the color of the pixel wand with a string (e.g. "blue", "#0000ff", "rgb(0,0,255)", "cmyk(100,100,100,10)", etc.)
func (pw *PixelWand) SetColor(color string) bool {
	cscolor := C.CString(color)
	defer C.free(unsafe.Pointer(cscolor))
	return 1 == int(C.PixelSetColor(pw.pw, cscolor))
}

// Sets the color count of the pixel wand
func (pw *PixelWand) SetColorCount(count uint) {
	C.PixelSetColorCount(pw.pw, C.ulong(count))
}

// Sets the normalized cyan color of the pixel wand
func (pw *PixelWand) SetCyan(cyan float64) {
	C.PixelSetCyan(pw.pw, C.double(cyan))
}

// Sets the cyan color of the pixel wand
func (pw *PixelWand) SetCyanQuantum(cyan Quantum) {
	C.PixelSetCyanQuantum(pw.pw, C.Quantum(cyan))
}

// Sets the normalized green color of the pixel wand
func (pw *PixelWand) SetGreen(green float64) {
	C.PixelSetGreen(pw.pw, C.double(green))
}

// Sets the green color of the pixel wand
func (pw *PixelWand) SetGreenQuantum(green Quantum) {
	C.PixelSetGreenQuantum(pw.pw, C.Quantum(green))
}

// Sets the normalized magenta color of the pixel wand
func (pw *PixelWand) SetMagenta(magenta float64) {
	C.PixelSetMagenta(pw.pw, C.double(magenta))
}

// Sets the magenta color of the pixel wand
func (pw *PixelWand) SetMagentaQuantum(magenta Quantum) {
	C.PixelSetMagentaQuantum(pw.pw, C.Quantum(magenta))
}

// Sets the normalized opacity color of the pixel wand
func (pw *PixelWand) SetOpacity(opacity float64) {
	C.PixelSetOpacity(pw.pw, C.double(opacity))
}

// Sets the opacity color of the pixel wand
func (pw *PixelWand) SetOpacityQuantum(opacity Quantum) {
	C.PixelSetOpacityQuantum(pw.pw, C.Quantum(opacity))
}

// Sets the color of the pixel wand
func (pw *PixelWand) SetQuantumColor(color *PixelPacket) {
	C.PixelSetQuantumColor(pw.pw, color.pp)
}

// Sets the normalized red color of the pixel wand
func (pw *PixelWand) SetRed(red float64) {
	C.PixelSetRed(pw.pw, C.double(red))
}

// Sets the red color of the pixel wand
func (pw *PixelWand) SetRedQuantum(red Quantum) {
	C.PixelSetRedQuantum(pw.pw, C.Quantum(red))
}

// Sets the normalized yellow color of the pixel wand
func (pw *PixelWand) SetYellow(yellow float64) {
	C.PixelSetYellow(pw.pw, C.double(yellow))
}

// Sets the yellow color of the pixel wand
func (pw *PixelWand) SetYellowQuantum(yellow Quantum) {
	C.PixelSetYellowQuantum(pw.pw, C.Quantum(yellow))
}