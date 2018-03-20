package codec

import (
	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/index"
	"github.com/ctripcorp/nephele/transform"
)

// Codec represents how to encode and decode image file name or its commands.
type Codec interface {
	Encoder(ctx *context.Context) Encoder
	Decoder(ctx *context.Context) Decoder
}

// Encoder represents how to encode image file name or its commands.
type Encoder interface {
	// Encode image file name.
	Encode(seed string) string
}

// Decoder represnets how to build commands and index from request image url.
type Decoder interface {
	// Decode from request image url.
	// generally there will be multi versions for image file name encoding.
	// and also we will have different image handle commands.
	Decode(uri string) error

	// Create indexer from request image url.
	CreateIndex() index.Index

	// Create transformer from request image url.
	CreateTransformer() transform.Transformer
}

// Config represents how to build codec.
type Config interface {
	BuildCodec() (Codec, error)
}

var codec Codec

// Init codec with provided configuration.
func Init(conf Config) error {
	var err error
	codec, err = conf.BuildCodec()
	return err
}

// Return customized encoder.
func GetEncoder(ctx *context.Context) Encoder {
	return codec.Encoder(ctx)
}

// Return customized decoder.
func GetDecoder(ctx *context.Context) Decoder {
	return codec.Decoder(ctx)
}
