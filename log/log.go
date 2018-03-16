package log

type Config interface {
	BuildLogger() (Logger, error)
}

var instance Logger

func Init(conf Config) (err error) {
	instance, err = conf.BuildLogger()
	return
}

func DefaultConfig() (Config, error) {
	return nil, nil
}

func Debugf(format string, values ...interface{}) {
	instance.Debugf(format, values...)
}

func Infof(format string, values ...interface{}) {
	instance.Infof(format, values...)
}

func Warnf(format string, values ...interface{}) {
	instance.Warnf(format, values...)
}

func Errorf(format string, values ...interface{}) {
	instance.Errorf(format, values...)
}

func Fatalf(format string, values ...interface{}) {
	instance.Fatalf(format, values...)
}

func Debugw(message string, keysAndValues ...interface{}) {
	instance.Debugw(message, keysAndValues...)
}

func Infow(message string, keysAndValues ...interface{}) {
	instance.Infow(message, keysAndValues...)
}

func Warnw(message string, keysAndValues ...interface{}) {
	instance.Warnw(message, keysAndValues...)
}

func Errorw(message string, keysAndValues ...interface{}) {
	instance.Errorw(message, keysAndValues...)
}

func Fatalw(message string, keysAndValues ...interface{}) {
	instance.Fatalw(message, keysAndValues...)
}

func TraceBegin(keysAndValues ...interface{}) TracerNeedATime {
	return instance.TraceBegin(keysAndValues...)
}

func TraceEnd(state interface{}, message ...string) TracerNeedATime {
	return instance.TraceEnd(state, message...)
}

func TraceEndRoot(state interface{}, message ...string) TracerNeedATime {
	return instance.TraceEndRoot(state, message...)
}
