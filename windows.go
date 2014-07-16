// +build windows
package main

import (
	"os"
	"os/user"
	"syscall"
)

const (
	acceptPasswordFromStdin = false
	configDir               = "_modou"
)

func homePath() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return u.HomeDir, nil
}

// IsTerminal returns false on Windows.
func IsTerminal(f *os.File) bool {
	ft, _ := syscall.GetFileType(syscall.Handle(f.Fd()))
	return ft == syscall.FILE_TYPE_CHAR
}
