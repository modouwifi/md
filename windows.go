// +build windows
package main

import "os/user"

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
