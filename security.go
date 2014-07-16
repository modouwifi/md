package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/modouwifi/md/api"
)

var cmdSecurity = cli.Command{
	Name:  "security",
	Usage: "System security",
	Subcommands: []cli.Command{
		{
			Name:   "get",
			Usage:  "Get system security's status",
			Action: runSecurityGet,
		},
		{
			Name:   "set",
			Usage:  "Set system security's status",
			Action: runSecuritySet,
		},
	},
}

func runSecurityGet(c *cli.Context) {
	req, err := client.NewRequest("GET", "/security/get_config")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	var result api.SecurityEnabled
	err = req.ToJSON(&result)
	fmt.Println(result.Code)
	fmt.Println(result.Enabled)
}

func runSecuritySet(c *cli.Context) {
	req, err := client.NewRequest("POST", "/security/set_config")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	var result api.SecurityEnabled
	err = req.ToJSON(&result)
	fmt.Println(result.Code)
	fmt.Println(result.Enabled)
}
