package main

import (
  "fmt"
	"bufio"
	"os"
	"path/filepath"
	"runtime"

	"github.com/codegangsta/cli"
)

const VERSION = "0.0.0"

var (
	client  *Client
	config  *Config
	app     = cli.NewApp()
	apiURL  = "http://modouwifi.net/api"
	mdAgent = "md/" + VERSION + " (" + runtime.GOOS + "; " + runtime.GOARCH + ")"
	stdin   = bufio.NewReader(os.Stdin)
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
	app.Usage = "[COMMANDS]"
	app.Version = VERSION
	app.Commands = append(app.Commands,
		cmdLogin,
		cmdLogout,

		cmdSystem,
	)
}

func main() {
	initClient()

	app.Run(os.Args)
}
