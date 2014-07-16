// +build darwin freebsd linux netbsd openbsd
package main

import (
	"fmt"
	"os"
	"os/exec"
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

// IsTerminal returns true if f is a terminal.
func IsTerminal(f *os.File) bool {
	cmd := exec.Command("test", "-t", "0")
	cmd.Stdin = f
	return cmd.Run() == nil
}
