package log

type PlainTextLogger interface {
	Debugf(format string, keysAndValues ...interface{})
	Infof(format string, keysAndValues ...interface{})
	Warnf(format string, keysAndValues ...interface{})
	Errorf(format string, keysAndValues ...interface{})
	Fatalf(format string, keysAndValues ...interface{})
}
