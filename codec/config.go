package codec

type Config interface {
	BuildCodec() (Codec, error)
}
