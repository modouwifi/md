package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/modouwifi/md/api"
)

var cmdWan = cli.Command{
	Name:  "wan",
	Usage: "WLAN",
	Subcommands: []cli.Command{
		{
			Name:   "info",
			Usage:  "Display WLAN Info",
			Action: runWanInfo,
		},
	},
}

func runWanInfo(c *cli.Context) {
	req, err := client.NewRequest("GET", "/wan/get_info")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	var result api.Wan
	err = req.ToJSON(&result)
	fmt.Println(result)
}
