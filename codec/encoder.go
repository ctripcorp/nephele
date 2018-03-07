package codec

// Encoder represents how to encode image file name or its commands.
type Encoder interface {
	// Encode image file name.
	Encode(seed string) string
}
