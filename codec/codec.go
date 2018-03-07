package codec

// Codec represents how to encode and decode image file name or its commands.
type Codec interface {
	Encoder() Encoder
	Decoder() Decoder
}

var codec Codec

// Init codec with provided configuration.
func Init(conf Config) error {
	var err error
	codec, err = conf.BuildCodec()
	return err
}

// Return customized encoder.
func GetEncoder() Encoder {
	return codec.Encoder()
}

// Return customized decoder.
func GetDecoder() Decoder {
	return codec.Decoder()
}
