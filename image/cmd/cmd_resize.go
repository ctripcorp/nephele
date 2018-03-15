package cmd

import (
	"math"
	"strings"

	"github.com/nephele/context"
	"github.com/nephele/img4go/gm"
)

type ResizeCommand struct {
	Wand              *gm.MagickWand
	Width             uint
	Height            uint
	Method            string //Lfit/Fixed
	Limit, Percentage int
}

func (r *ResizeCommand) Exec(ctx context.Context) error {
	println("w", r.Wand.Width(), "h", r.Wand.Height())
	if r.Width > r.Wand.Width() && r.Height > r.Wand.Height() && r.Limit == 0 {
		return nil
	}
	var w, h uint
	if r.Percentage != 0 {
		w, h = r.percentage(ctx, r.Wand)
	}
	if strings.ToUpper(r.Method) == "FIXED" {
		w, h = r.fixed(ctx, r.Wand)
	}
	w, h = r.lfit(ctx, r.Wand)
	return r.Wand.LanczosResize(w, h)
}

//Lfit: 等比缩略
func (r *ResizeCommand) lfit(ctx context.Context, img *gm.MagickWand) (uint, uint) {
	var width, height uint
	width = img.Width()
	height = img.Height()
	w, h := r.Width, r.Height
	//auto compute weight or height
	if w == 0 {
		w = width * h / height
		return w, h
	}
	if h == 0 {
		h = height * w / width
		return w, h
	}

	p1 := float64(r.Width) / float64(r.Height)
	p2 := float64(width) / float64(height)

	if p2 > p1 {
		h = uint(math.Floor(float64(r.Width) / p2))
		if uint(math.Abs(float64(h-r.Height))) < 3 {
			h = r.Height
		}
	} else {
		w = uint(math.Floor(float64(r.Height) * p2))
		if uint(math.Abs(float64(w-r.Width))) < 3 {
			w = r.Width
		}
	}
	return w, h
}

//Fixed: 固定宽高，强制缩略
func (r *ResizeCommand) fixed(ctx context.Context, img *gm.MagickWand) (uint, uint) {
	if r.Width < 1 && r.Height < 1 {
		return r.Width, r.Height
	}
	if r.Width < 1 || r.Height < 1 {
		var width, height uint
		width = img.Width()
		height = img.Height()
		if r.Width < 1 {
			r.Width = width
		}
		if r.Height < 1 {
			r.Height = height
		}
	}
	return r.Width, r.Height
}

//倍数百分比。
func (r *ResizeCommand) percentage(ctx context.Context, img *gm.MagickWand) (uint, uint) {
	r.Width = uint(int(img.Width()) * r.Percentage / 100)
	r.Height = uint(int(img.Height()) * r.Percentage / 100)
	return r.Width, r.Height
}
