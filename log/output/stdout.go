package output

import (
	"os"
)

type StdoutConfig struct {
	Level string `toml:"level"`
}

func (sc *StdoutConfig) Build() (Output, error) {
	return &basicOutput{
		os.Stdout,
		nil,
		sc.Level,
	}, nil
}
