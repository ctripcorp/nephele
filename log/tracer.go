package log

type Tracer interface {
	Track(keysAndValues ...interface{})
	Sum(state interface{}, message ...string)
}
