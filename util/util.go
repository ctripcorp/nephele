package util

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

func HomePath() (string, error) {
	var homeDir string
	// By default, store image and log files in current users home directory
	u, err := user.Current()
	if err == nil {
		homeDir = u.HomeDir
	} else if os.Getenv("HOME") != "" {
		homeDir = os.Getenv("HOME")
	} else {
		return homeDir, fmt.Errorf("failed to determine current user")
	}
	return filepath.Join(homeDir, "nephele"), nil
}

func FromToml(path string, v interface{}) error {
	var err error
	var blob []byte

	if blob, err = ioutil.ReadFile(path); err != nil {
		return err
	}

	// Handle any potential Byte-Order-Marks that may be in the config file.
	// This is for Windows compatibility only.
	bom := unicode.BOMOverride(transform.Nop)
	if blob, _, err = transform.Bytes(bom, blob); err != nil {
		return err
	}

	_, err = toml.Decode(string(blob), v)
	return err
}
