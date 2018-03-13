package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

type CompositeOperator int

const (
	COMPOSITE_OP_UNDEFINED         CompositeOperator = C.UndefinedCompositeOp
	COMPOSITE_OP_NO                CompositeOperator = C.NoCompositeOp
	COMPOSITE_OP_ATOP              CompositeOperator = C.AtopCompositeOp
	COMPOSITE_OP_BUMPMAP           CompositeOperator = C.BumpmapCompositeOp
	COMPOSITE_OP_CLEAR             CompositeOperator = C.ClearCompositeOp
	COMPOSITE_OP_COLOR_BURN        CompositeOperator = C.ColorBurnCompositeOp
	COMPOSITE_OP_COLOR_DODGE       CompositeOperator = C.ColorDodgeCompositeOp
	COMPOSITE_OP_COLORIZE          CompositeOperator = C.ColorizeCompositeOp
	COMPOSITE_OP_COPY_BLACK        CompositeOperator = C.CopyBlackCompositeOp
	COMPOSITE_OP_COPY_BLUE         CompositeOperator = C.CopyBlueCompositeOp
	COMPOSITE_OP_COPY              CompositeOperator = C.CopyCompositeOp
	COMPOSITE_OP_COPY_CYAN         CompositeOperator = C.CopyCyanCompositeOp
	COMPOSITE_OP_COPY_GREEN        CompositeOperator = C.CopyGreenCompositeOp
	COMPOSITE_OP_COPY_MAGENTA      CompositeOperator = C.CopyMagentaCompositeOp
	COMPOSITE_OP_COPY_OPACITY      CompositeOperator = C.CopyOpacityCompositeOp
	COMPOSITE_OP_COPY_RED          CompositeOperator = C.CopyRedCompositeOp
	COMPOSITE_OP_COPY_YELLOW       CompositeOperator = C.CopyYellowCompositeOp
	COMPOSITE_OP_DARKEN            CompositeOperator = C.DarkenCompositeOp
	COMPOSITE_OP_DIFFERENCE        CompositeOperator = C.DifferenceCompositeOp
	COMPOSITE_OP_DISPLACE          CompositeOperator = C.DisplaceCompositeOp
	COMPOSITE_OP_DISSOLVE          CompositeOperator = C.DissolveCompositeOp
	COMPOSITE_OP_EXCLUSION         CompositeOperator = C.ExclusionCompositeOp
	COMPOSITE_OP_HARD_LIGHT        CompositeOperator = C.HardLightCompositeOp
	COMPOSITE_OP_HUE               CompositeOperator = C.HueCompositeOp
	COMPOSITE_OP_IN                CompositeOperator = C.InCompositeOp
	COMPOSITE_OP_LIGHTEN           CompositeOperator = C.LightenCompositeOp
	COMPOSITE_OP_LINEAR_LIGHT      CompositeOperator = C.LinearLightCompositeOp
	COMPOSITE_OP_LUMINIZE          CompositeOperator = C.LuminizeCompositeOp
	COMPOSITE_OP_MODULATE          CompositeOperator = C.ModulateCompositeOp
	COMPOSITE_OP_MULTIPLY          CompositeOperator = C.MultiplyCompositeOp
	COMPOSITE_OP_OUT               CompositeOperator = C.OutCompositeOp
	COMPOSITE_OP_OVER              CompositeOperator = C.OverCompositeOp
	COMPOSITE_OP_OVERLAY           CompositeOperator = C.OverlayCompositeOp
	COMPOSITE_OP_PLUS              CompositeOperator = C.PlusCompositeOp
	COMPOSITE_OP_REPLACE           CompositeOperator = C.ReplaceCompositeOp
	COMPOSITE_OP_SATURATE          CompositeOperator = C.SaturateCompositeOp
	COMPOSITE_OP_SCREEN            CompositeOperator = C.ScreenCompositeOp
	COMPOSITE_OP_SOFT_LIGHT        CompositeOperator = C.SoftLightCompositeOp
	COMPOSITE_OP_THRESHOLD         CompositeOperator = C.ThresholdCompositeOp
	COMPOSITE_OP_XOR               CompositeOperator = C.XorCompositeOp
	COMPOSITE_OP_VIVID_LIGHT       CompositeOperator = C.VividLightCompositeOp
	COMPOSITE_OP_PIN_LIGHT         CompositeOperator = C.PinLightCompositeOp
	COMPOSITE_OP_LINEAR_DODGE      CompositeOperator = C.LinearDodgeCompositeOp
	COMPOSITE_OP_LINEAR_BURN       CompositeOperator = C.LinearBurnCompositeOp
)
