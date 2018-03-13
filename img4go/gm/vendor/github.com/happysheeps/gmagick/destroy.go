package gmagick

import "github.com/happysheeps/gmagick/types"

// Destroy instance of Destroyer
// If GOGC=off you should call obj.Destroy() manually
func Destroy(d types.Destroyer) {
	d.Destroy()
}
