package cmd

// import (
// 	"strings"

// 	"github.com/ctripcorp/nephele.image.service/svr/logic/util"
// 	"github.com/nephele/img4go/gm"

// 	"github.com/nephele/context"
// )

// const (
// 	CMDTITLE = "URL.Command"
// )

// type ICommand interface {
// 	Exec(ctx context.Context, img *gm.MagickWand) error
// }

// type IConvert interface {
// 	Exec(ctx context.Context, bts []byte) ([]byte, error)
// }

// type Cmds struct {
// 	BeforeCommands []IConvert
// 	Commands       []ICommand
// 	AfterCommands  []IConvert
// 	bts            []byte
// }

// func CreateCmds(ctx context.Context, imgURL *ImageURL, bts []byte) (*Cmds, error) {
// 	cs := &Cmds{}

// 	cs.bts = bts
// 	cs.BeforeCommands = []IConvert{&TTOIConvert{ImgURL: imgURL}}

// 	cs.Commands = []ICommand{&PanoramaCommand{ImgURL: imgURL}, &AutoOrientCommand{ImgURL: imgURL}, &CorpCommand{ImgURL: imgURL}, &WaterMarkCommand{ImgURL: imgURL, Type: Priority}}
// 	channelName, _ := imgURL.GetChannelName(evt, structs.Meta)
// 	sequences, err := structs.Meta.GetConfig(evt, channelName, "", util.CFG_SequenceOfOperation)
// 	if err != nil {
// 		return nil, err
// 	}
// 	arr := strings.Split(sequences, ",")
// 	for _, cmd := range arr {
// 		switch cmd {
// 		case CMD_STRIP:
// 			cs.Commands = append(cs.Commands, &StripCommand{ImgURL: imgURL})
// 			break
// 		case CMD_RESIZE:
// 			cs.Commands = append(cs.Commands, &ResizeCommand{ImgURL: imgURL})
// 			break
// 		case CMD_QUALITY:
// 			cs.Commands = append(cs.Commands, &QualityCommand{ImgURL: imgURL})
// 			break
// 		case CMD_ROTATE:
// 			cs.Commands = append(cs.Commands, &RotateCommand{ImgURL: imgURL})
// 			break
// 		case CMD_WATERMARK:
// 			cs.Commands = append(cs.Commands, &WaterMarkCommand{ImgURL: imgURL, Type: Name}, &WaterMarkCommand{ImgURL: imgURL, Type: Logo}) //
// 			break
// 		case CMD_FORMAT:
// 			cs.Commands = append(cs.Commands, &FormatCommand{ImgURL: imgURL})
// 			break
// 		case CMD_DIGITALWATERMARK:
// 			cs.Commands = append(cs.Commands, &DigitalWaterMarkCommand{ImgURL: imgURL})
// 			break
// 		case CMD_SHARPEN:
// 			cs.Commands = append(cs.Commands, &SharpenCommand{ImgURL: imgURL})
// 			break
// 		case CMD_PACKAGE:
// 			cs.Commands = append(cs.Commands, &PackageCommand{ImgURL: imgURL})
// 			break
// 		}
// 	}
// 	cs.AfterCommands = []IConvert{&PackageConvert{ImgURL: imgURL}}
// 	return cs, nil
// }

// func (c *Cmds) FileSize() int {
// 	return len(c.bts)
// }

// func (c *Cmds) Build(ctx context.Context) (bts1 []byte, err error) {
// 	bts := c.bts
// 	//exec convert
// 	if c.BeforeCommands != nil {
// 		for _, convert := range c.BeforeCommands {
// 			bts, err = convert.Exec(evt, bts)
// 			if err != nil {
// 				return
// 			}
// 			if err = util.IsCtxDead(evt, ctx); err != nil {
// 				return
// 			}
// 		}
// 	}
// 	//exec command
// 	var img *img4g.Image
// 	img, err = img4g.CreateImage(evt, bts)
// 	if err != nil {
// 		return
// 	}
// 	defer img.Destory(evt)
// 	for _, command := range c.Commands {
// 		if err = command.Exec(evt, img); err != nil {
// 			return
// 		}
// 		if err = util.IsCtxDead(evt, ctx); err != nil {
// 			return
// 		}
// 	}
// 	bts, err = img.WriteImageBlob(evt)
// 	if err != nil {
// 		return
// 	}

// 	//exec convert
// 	if c.AfterCommands != nil {
// 		for _, convert := range c.AfterCommands {
// 			bts, err = convert.Exec(evt, bts)
// 			if err != nil {
// 				return
// 			}
// 			if err = util.IsCtxDead(evt, ctx); err != nil {
// 				return
// 			}
// 		}
// 	}
// 	bts1 = bts
// 	return
// }
