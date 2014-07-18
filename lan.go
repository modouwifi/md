package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/modouwifi/md/api"
)

var cmdLan = cli.Command{
	Name:  "lan",
	Usage: "LAN",
	Subcommands: []cli.Command{
		{
			Name:   "get",
			Usage:  "Get LAN config",
			Action: runLanGet,
		},
		{
			Name:   "check",
			Usage:  "Check LAN status",
			Action: runLanCheck,
		},
	},
}

func runLanGet(c *cli.Context) {
	req, err := client.NewRequest("GET", "/lan/get_lan_config")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	var result api.Lan
	err = req.ToJSON(&result)
	fmt.Println(result)
}

func runLanCheck(c *cli.Context) {
	req, err := client.NewRequest("GET", "/lan/check_set")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	var result api.ResponseMessage
	err = req.ToJSON(&result)
	fmt.Println(result)
}
