package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/codegangsta/cli"
)

const VERSION = "0.0.0"

var (
	client  *Client
	config  *Config
	apiURL  = "http://modouwifi.net/api"
	mdAgent = "md/" + VERSION + " (" + runtime.GOOS + "; " + runtime.GOARCH + ")"
	app     = cli.NewApp()
)

func initClient() {
	client = NewClient(apiURL, mdAgent)
	config = &Config{}
	file, err := LocateRcfile()
	if err != nil {
		err = os.MkdirAll(filepath.Dir(file), os.ModePerm)
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err := os.OpenFile(file, os.O_CREATE, os.ModePerm)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	config.Filename = file
	config.Read()
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app.Name = "md"
	app.Usage = "Command Line Utility for Modou"
	app.Version = VERSION
	app.Commands = append(app.Commands,
		cmdLogin,
		cmdLogout,

		cmdSystem,
		cmdSecurity,
	)
}

func main() {
	initClient()

	app.Run(os.Args)
}
