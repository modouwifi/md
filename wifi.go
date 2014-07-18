package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/modouwifi/md/api"
)

var cmdWifi = cli.Command{
	Name:  "wifi",
	Usage: "WIFI",
	Subcommands: []cli.Command{
		{
			Name:   "get",
			Usage:  "Get WIFI config",
			Action: runWIFIGetConfig,
		},
		{
			Name:   "set",
			Usage:  "Set WIFI config",
			Action: runWIFISetConfig,
		},
		{
			Name:   "check",
			Usage:  "Check WIFI config",
			Action: runWIFICheck,
		},
		{
			Name:   "status",
			Usage:  "Show WIFI status",
			Action: runWIFICheck,
		},
	},
}

func runWIFIGetConfig(c *cli.Context) {
	req, err := client.NewRequest("GET", "/wifi/get_config")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	var result api.Wifi
	err = req.ToJSON(&result)
	fmt.Println(result)
}

func runWIFISetConfig(c *cli.Context) {
	req, err := client.NewRequest("POST", "/wifi/set_config")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	var result api.Wifi
	err = req.ToJSON(&result)
	fmt.Println(result)
}

func runWIFICheck(c *cli.Context) {
	req, err := client.NewRequest("GET", "/wifi/check_set")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	var result api.ResponseMessage
	err = req.ToJSON(&result)
	fmt.Println(result)
}

func runWIFIStatus(c *cli.Context) {
	req, err := client.NewRequest("GET", "/wifi/is_enabled")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	var result api.ResponseMessage
	err = req.ToJSON(&result)
	fmt.Println(result)
}
