package command

import (
	"fmt"
	"strconv"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/img4go/gm"
	"github.com/ctripcorp/nephele/log"
)

type Quality struct {
	Quality uint
}

const (
	qualityKeyV string = "v"
)

//Verify verify quality
func (q *Quality) Verify(ctx *context.Context, params map[string]string) error {
	log.Debugf(ctx, "quality verification")
	for k, v := range params {
		if k == qualityKeyV {
			quality, e := strconv.Atoi(v)
			if e != nil || quality < 0 || quality > 100 {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			q.Quality = uint(quality)
		}
	}
	return nil
}

//Exec exec
func (q *Quality) Exec(ctx *context.Context, wand *gm.MagickWand) error {
	var err error
	log.TraceBegin(ctx, "", "URL.Command", "quality", "quality", q.Quality)
	defer log.TraceEnd(ctx, err)
	if q.Quality == 0 {
		err = wand.PreserveJPEGSettings()
	} else {
		err = wand.SetCompressionQuality(uint(q.Quality))
	}
	return err
}
