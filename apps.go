package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/modouwifi/md/api"
)

var cmdApps = cli.Command{
	Name:   "apps",
	Usage:  "List apps, show apps",
	Action: runApps,
	Subcommands: []cli.Command{
		{
			Name:   "list",
			Usage:  "List installed apps",
			Action: runAppsList,
		},
	},
}

func runApps(c *cli.Context) {
	args := c.Args()
	if len(args) == 0 {
		cli.ShowSubcommandHelp(c)
		return
	}
	appId := args.First()
	app := getAppInfoById(appId)
	fmt.Println(app)
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

func getAppInfoById(id string) api.App {
	req, err := client.NewRequest("GET", "/plugin/plugin_latest_info")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	req.QueryString("id=" + id)
	var result api.App
	err = req.ToJSON(&result)
	return result
}
