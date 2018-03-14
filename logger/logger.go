package logger

import (
	"go.uber.org/zap"
	"io"
)

type Config struct {
}

type Logger struct {
	*zap.Logger
}

func New(conf Config, writer io.Writer) *Logger {
	return nil
}

func NewConfig() Config {
	return Config{}
}
