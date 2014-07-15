// +build darwin freebsd linux netbsd openbsd
package main

import (
	"fmt"
	"os"
)

const (
	acceptPasswordFromStdin = true
	configDir               = ".modou"
)

func homePath() (string, error) {
	home := os.Getenv("HOME")
	if home == "" {
		return "", fmt.Errorf("Environment variable HOME not set")
	}

	return home, nil
}
