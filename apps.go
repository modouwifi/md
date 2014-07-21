package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/modouwifi/md/api"
)

var cmdApps = cli.Command{
	Name:  "apps",
	Usage: "List apps, show apps",
	Subcommands: []cli.Command{
		{
			Name:   "list",
			Usage:  "List installed apps",
			Action: runAppsList,
		},
	},
}

func runAppsList(c *cli.Context) {
	req, err := client.NewRequest("GET", "/plugin/installed_plugins")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	var result api.Apps
	err = req.ToJSON(&result)
	fmt.Println(result)
}
