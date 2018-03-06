package codec

type Encoder interface {
	Encode(seed string) string
}
