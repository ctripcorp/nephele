package command

import "github.com/ctripcorp/nephele/context"
import "github.com/ctripcorp/nephele/img4go/gm"

type GraphicsMagickCommand interface {
	Wand(ctx *context.Context, params map[string]string) error
	Blob(ctx *context.Context, wand *gm.MagickWand) error
}

const (
	RESIZE     string = "resize"
	CROP       string = "crop"
	ROTATE     string = "rotate"
	AUTOORIENT string = "autoorient"
	FORMAT     string = "format"
	QUALITY    string = "quality"
	WATERMARK  string = "watermark"
	SHARPEN    string = "sharpen"
	// STYLE      string = "style"
	// PANORAMA   string = "panorama"
)

var invalidInfoFormat = "The value: %s of parameter: %s is invalid."
