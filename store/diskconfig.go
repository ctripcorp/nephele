package store

type DiskConfig struct {
	Dir string `toml:"dir"`
}

func (conf *DiskConfig) BuildStore() (Store, error) {
	return nil, nil
}
