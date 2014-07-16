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
		{
			Name:   "ping",
			Usage:  "ping network",
			Action: runWanPing,
		},
		{
			Name:   "dhcp",
			Usage:  "Display DHCP Info",
			Action: runWanDHCPInfo,
		},
		{
			Name:   "pppoe",
			Usage:  "Display PPPOE Info",
			Action: runWanPPPOEInfo,
		},
		{
			Name:   "static",
			Usage:  "Display STATIC Info",
			Action: runWanSTATICInfo,
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

func runWanPing(c *cli.Context) {
	req, err := client.NewRequest("GET", "/wan/is_internet_available")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	var result api.ResponseMessage
	err = req.ToJSON(&result)
	fmt.Println(result)
}

func runWanDHCPInfo(c *cli.Context) {
	req, err := client.NewRequest("GET", "/wan/get_info/dhcp")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	var result api.WanDHCP
	err = req.ToJSON(&result)
	fmt.Println(result)
}

func runWanPPPOEInfo(c *cli.Context) {
	req, err := client.NewRequest("GET", "/wan/get_info/pppoe")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	var result api.WanPPPOE
	err = req.ToJSON(&result)
	fmt.Println(result)
}

func runWanSTATICInfo(c *cli.Context) {
	req, err := client.NewRequest("GET", "/wan/get_info/static")
	if err != nil {
		fmt.Println(err)
	}
	req.SetCookie(config.Cookie)
	var result api.WanSTATIC
	err = req.ToJSON(&result)
	fmt.Println(result)
}
