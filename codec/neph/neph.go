package neph

import (
	"github.com/ctripcorp/nephele/codec"
	"github.com/ctripcorp/nephele/context"
)

//Codec as neph codec
type Codec struct {
}

//Config config
type Config struct {
}

//DefaultConfig  new config
func DefaultConfig() (*Config, error) {
	return new(Config), nil
}

//BuildCodec build codec
func (conf *Config) BuildCodec() (codec.Codec, error) {
	return new(Codec), nil
}

//Encoder  new encoder
func (c *Codec) Encoder(ctx *context.Context) codec.Encoder {
	return &Encoder{ctx: ctx}
}

//Decoder new decoder
func (c *Codec) Decoder(ctx *context.Context) codec.Decoder {
	return &Decoder{ctx: ctx}
}
