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
	log.Debugf(ctx, "rotate verification")
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
	var err error
	log.TraceBegin(ctx, "", "URL.Command", "rotate", "degree", r.Degree)
	defer log.TraceEnd(ctx, err)
	err = wand.Rotate(float64(r.Degree))
	return err
}
