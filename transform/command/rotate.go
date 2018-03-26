package command

import (
	"fmt"
	"strconv"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
	"github.com/ctripcorp/nephele/log"
)

//Rotate rotate command
type Rotate struct {
	Degree int
}

const (
	rotateV string = "v"
)

//Verify rotate verify params
func (r *Rotate) Verify(ctx *context.Context, params map[string]string) error {
	log.Debugw(ctx, "begin rotate verify")
	for k, v := range params {
		if k == rotateV {
			degree, e := strconv.Atoi(v)
			if e != nil || degree < 0 || degree > 360 {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			r.Degree = degree
		}
	}
	return nil
}

//Exec rotate exec
func (r *Rotate) Exec(ctx *context.Context, wand *gm.MagickWand) error {
	log.TraceBegin(ctx, fmt.Sprintf("rotate,degree:%d", r.Degree), "URL.Command", "rotate")
	defer log.TraceEnd(ctx, nil)
	return wand.Rotate(float64(r.Degree))
}
