package codec

// Config represents how to build codec.
type Config interface {
	BuildCodec() (Codec, error)
}
