package command

import (
	"fmt"
	"strconv"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
	"github.com/ctripcorp/nephele/log"
)

type Sharpen struct {
	Radius float64
	Sigma  float64
}

const (
	sharpenR string = "r"
	sharpenS string = "s"
)

//Verify sharpen Verify
func (s *Sharpen) Verify(ctx *context.Context, params map[string]string) error {
	log.Debugf(ctx, "sharpen verification")
	for k, v := range params {
		if k == sharpenR {
			radius, e := strconv.ParseFloat(v, 64)
			if e != nil {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			s.Radius = radius
		}
		if k == sharpenS {
			sigma, e := strconv.ParseFloat(v, 64)
			if e != nil {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			s.Sigma = sigma
		}
	}
	return nil
}

//Exec sharpen exec
func (s *Sharpen) Exec(ctx *context.Context, wand *gm.MagickWand) error {
	log.TraceBegin(ctx, "sharpen executive", "radius", s.Radius, "sigma", s.Sigma)
	defer log.TraceEnd(ctx, nil)
	return wand.Sharpen(s.Radius, s.Sigma)
}
