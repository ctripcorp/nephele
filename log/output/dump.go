package output

import (
	"os"
	"path/filepath"
)

type DumpConfig struct {
	Level string `toml:"level"`
	Path  string `toml:"path"`
}

func (dc *DumpConfig) Build() (Output, error) {
	f, err := os.OpenFile(filepath.Join(dc.Path, "now.log"),
		os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		println(err.Error())
		return nil, err
	}

	return &basicOutput{
		f,
		dc.Level,
	}, nil
}
