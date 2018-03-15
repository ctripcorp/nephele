package util

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func HomeDir() (string, error) {
	var homeDir string
	// By default, store image and log files in current users home directory
	u, err := user.Current()
	if err == nil {
		homeDir = u.HomeDir
	} else if os.Getenv("HOME") != "" {
		homeDir = os.Getenv("HOME")
	} else {
		return homeDir, fmt.Errorf("failed to determine current user for storage")
	}
	return filepath.Join(homeDir, "nephele"), nil
}
