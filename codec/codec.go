package codec

type Codec interface {
	Encoder() Encoder
	Decoder() Decoder
}

var codec Codec

func Init(conf Config) error {
	var err error
	codec, err = conf.BuildCodec()
	return err
}

func GetEncoder() Encoder {
	return codec.Encoder()
}

func GetDecoder() Decoder {
	return codec.Decoder()
}
