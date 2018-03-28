package command

import (
	"fmt"
	"strconv"

	"github.com/ctripcorp/nephele/img4go/gm"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/log"
)

type Quality struct {
	Quality uint
}

const (
	qualityKeyV string = "v"
)

//Verify  quality verify
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

//Exec quality exec
func (q *Quality) Exec(ctx *context.Context, wand *gm.MagickWand) error {
	log.TraceBegin(ctx, "", "URL.Command", "quality", q.Quality)
	defer log.TraceEnd(ctx, nil)
	if q.Quality == 0 || q.Quality == 100 {
		return nil
	}
	return wand.SetCompressionQuality(q.Quality)
}
