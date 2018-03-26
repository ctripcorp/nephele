package command

import (
	"fmt"
	"strconv"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/log"
)

type Quality struct {
	Quality int
	Aio     int
}

const (
	qualityKeyV   string = "v"
	qualityKeyAIO string = "aio"
)

//Quality verify quality
func (q *Quality) Verify(ctx *context.Context, params map[string]string) error {
	log.Debugf(ctx, "quality verification")
	for k, v := range params {
		if k == qualityKeyV {
			quality, e := strconv.Atoi(v)
			if e != nil || quality < 0 || quality > 100 {
				return fmt.Errorf(invalidInfoFormat, v, k)
			}
			q.Quality = quality
		}
	}
	return nil
}
