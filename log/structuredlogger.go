package log

type StructuredLogger interface {
	Debugw(message string, keysAndValues ...interface{})
	Infow(message string, keysAndValues ...interface{})
	Warnw(message string, keysAndValues ...interface{})
	Errorw(message string, keysAndValues ...interface{})
	Fatalw(message string, keysAndValues ...interface{})
}
