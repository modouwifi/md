package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

const configFile = "config.json"

type Config struct {
	Filename string `json:"-"`
	Password string `json:"password,omitempty"`
	Cookie   *http.Cookie
}

func (c *Config) Read() error {
	f, err := os.Open(c.Filename)
	if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(c)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) Write() error {
	f, err := os.OpenFile(c.Filename, os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	err = encoder.Encode(c)
	if err != nil {
		return err
	}

	return nil
}

func RcfileStat(dir string) (string, error) {
	file := filepath.Join(dir, configFile)
	if _, err := os.Stat(file); err != nil {
		return file, err
	}
	return file, nil
}

func LocateRcfile() (string, error) {
	home, err := homePath()
	if err != nil {
		return "", err
	}
	file, err := RcfileStat(filepath.Join(home, configDir))
	if err == nil {
		return file, nil
	}
	return file, fmt.Errorf("error: Config file not found")
}
