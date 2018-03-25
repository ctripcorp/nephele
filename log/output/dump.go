package output

import (
	"os"
	"path/filepath"
	"time"
)

type DumpConfig struct {
	Level     string `toml:"level"`
	Path      string `toml:"path"`
	TimeBlock int    `toml:"time-block"`
}

func (dc *DumpConfig) Build() (Output, error) {
	o, err := createBasicOutput(
		func() (WriteSyncer, error) {
			return os.OpenFile(filepath.Join(dc.Path, time.Now().Format("2006-01-02H15M04S05")),
				os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		},
		dc.Level,
	)
	if err != nil {
		return nil, err
	}
	go func(bo *basicOutput) {
		for {
			time.Sleep(time.Duration(dc.TimeBlock) * time.Minute)
			bo.Reset()
		}
	}(o)
	return o, err
}
