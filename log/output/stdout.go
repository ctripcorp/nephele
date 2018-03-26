package output

import (
	"os"
)

type StdoutConfig struct {
	Level string `toml:"level"`
}

func (sc *StdoutConfig) Build() (Output, error) {
	return createBasicOutput(
		func() (WriteSyncer, error) {
			return os.Stdout, nil
		},
		sc.Level,
	)
}
