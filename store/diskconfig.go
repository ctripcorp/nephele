package store

type DiskConfig struct {
	Path string `toml:"path"`
}

func (conf *DiskConfig) BuildStore() (Store, error) {
	return &Disk{Path: conf.Path}, nil
}
