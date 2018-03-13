package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

type ChannelType int

const (
	CHANNEL_UNDEFINED  ChannelType = C.UndefinedChannel
	CHANNEL_RED        ChannelType = C.RedChannel
	CHANNEL_GRAY       ChannelType = C.GrayChannel
	CHANNEL_CYAN       ChannelType = C.CyanChannel
	CHANNEL_GREEN      ChannelType = C.GreenChannel
	CHANNEL_MAGENTA    ChannelType = C.MagentaChannel
	CHANNEL_BLUE       ChannelType = C.BlueChannel
	CHANNEL_YELLOW     ChannelType = C.YellowChannel
	CHANNEL_OPACITY    ChannelType = C.OpacityChannel
	CHANNEL_BLACK      ChannelType = C.BlackChannel
	CHANNEL_INDEX      ChannelType = C.IndexChannel
	CHANNELS_ALL       ChannelType = C.AllChannels
)
