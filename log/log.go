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

func Infof(format string, values ...interface{}) {
	instance.Infof(format, values...)
}
