package strip

import (
	"context"

	"github.com/ctrip-nephele/gmagick"
)

//Strip strip command
type Command struct {
}

func (c *Command) Support() string {
	return "wand"
}

//Verify strip Verify params
func (c *Command) Verify(ctx context.Context, params map[string]string) error {
	//log.Debugf(ctx, "strip verification")
	return nil
}

func (c *Command) ExecuteOnBlob(ctx context.Context, blob []byte) ([]byte, error) {
	return nil, nil
}

// Exec strip exec
func (c *Command) ExecuteOnWand(ctx context.Context, wand *gmagick.MagickWand) error {
	//log.TraceBegin(ctx, "", "URL.Command", "strip")
	//defer log.TraceEnd(ctx, err)
	return wand.StripImage()
}
