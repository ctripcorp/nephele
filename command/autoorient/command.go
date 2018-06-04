package autoorient

import (
	"context"

	"github.com/ctrip-nephele/gmagick"
)

type Command struct {
	AutoOrient string
}

const (
	commandKeyA string = "autoOrient"
)

func (c *Command) Support() string {
	return "wand"
}

//Verify autoorient verify
func (c *Command) Verify(ctx context.Context, params map[string]string) error {
	//log.Debugf(ctx, "autoOrient verification")
	return nil
}

func (c *Command) ExecuteOnBlob(ctx context.Context, blob []byte) ([]byte, error) {
	return nil, nil
}

//Exec autoorient exec
func (c *Command) ExecuteOnWand(ctx context.Context, wand *gmagick.MagickWand) error {
	//log.TraceBegin(ctx, "", "URL.Command", "autoOrient", "")
	//defer log.TraceEnd(ctx, nil)
	orientation := wand.GetImageOrientation()
	return wand.AutoOrientImage(orientation)
}
