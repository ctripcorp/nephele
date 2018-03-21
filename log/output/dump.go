package output

import (
    "os"
    "path/filepath"
)

type DumpConfig struct {
    Level string `toml:"level"`
    Path string `toml:"path"`
}

func (dc *DumpConfig) BuildConfig() (Output, error) {
    f, err := os.OpenFile(filepath.Join(dc.Path, "now.log"), 
            os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}

    return &basicOutput{
        f,
        dc.Level,
    }, nil
}
