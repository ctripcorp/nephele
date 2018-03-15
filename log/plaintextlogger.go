package log

type PlainTextLogger interface {
    Debugf(format string, values ...interface{})
    Infof(format string, values ...interface{})
    Warnf(format string, values ...interface{})
    Errorf(format string, values ...interface{})
    Fatalf(format string, values ...interface{})
}
