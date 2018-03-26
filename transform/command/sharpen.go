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
	log.Debugw(ctx, "begin sharpen verify")
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
	log.TraceBegin(ctx, fmt.Sprintf("sharpen,radius:%g,sigma:%g", s.Radius, s.Sigma), "URL.Command", "format")
	defer log.TraceEnd(ctx, nil)
	println(s.Radius, s.Sigma)
	return wand.Sharpen(s.Radius, s.Sigma)
}
