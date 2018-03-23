package output

import (
	"os"
	"path/filepath"
)

type DumpConfig struct {
	Level     string `toml:"level"`
	Path      string `toml:"path"`
	TimeBlock int    `toml:"time-block"`
}

func (dc *DumpConfig) Build() (Output, error) {
	f, err := os.OpenFile(filepath.Join(dc.Path, "now.log"),
		os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return &basicOutput{
		f,
		nil,
		dc.Level,
	}, nil
}
