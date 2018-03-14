package v1

import (
	"github.com/nephele/codec"
	"github.com/nephele/context"
)

//Represents as v1 codec
type Codec struct {
}

type Config struct {
}

func NewCodecConfig() (*Config, error) {
	return new(Config), nil
}

func (conf *Config) BuildCodec() (codec.Codec, error) {
	return new(Codec), nil
}

func (c *Codec) Encoder(ctx *context.Context) codec.Encoder {
	return &Encoder{ctx: ctx}
}

func (c *Codec) Decoder(ctx *context.Context) codec.Decoder {
	return &Decoder{ctx: ctx}
}
